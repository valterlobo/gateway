package server

import (
	"errors"
	"fmt"
	"gateway/command"
	"gateway/query"
	"github.com/gofiber/fiber"
	"github.com/gofiber/utils"
	"log"
	"strconv"
)

type AppGateway struct {
	commandGateway *command.Gateway
	queryGateway   *query.Gateway
	serverFiber    *fiber.App
}

func NewAppGateway() *AppGateway {

	appGateway := &AppGateway{}
	errorQuery := appGateway.configQueryGateway()
	fail(errorQuery)
	errorCmd := appGateway.configCommandGateway()
	fail(errorCmd)
	appGateway.serverFiber = fiber.New()

	//ROUTER
	appGateway.serverFiber.Post("/command/:namespace/:commandtype", appGateway.commandHandler)
	appGateway.serverFiber.Get("/query/:namespace/:querytype", appGateway.queryHandler)
	appGateway.serverFiber.Get("/status", appGateway.status)

	return appGateway

}

func (app *AppGateway) Start() error {

	return app.serverFiber.Listen(3000)

}
func (app *AppGateway) configQueryGateway() error {

	processor := query.NewQueryProcessor()
	if processor == nil {
		return errors.New("QueryProcessor not config")
	}
	app.queryGateway = &query.Gateway{Processor: processor}

	if app.queryGateway == nil {
		return errors.New("QueryGateway not config")
	}

	return nil
}

func (app *AppGateway) configCommandGateway() error {

	cmdProcessor := command.NewCommandProcessor()
	if cmdProcessor == nil {
		return errors.New("CMD Processor not config")
	}
	app.commandGateway = &command.Gateway{Processor: cmdProcessor}

	if app.commandGateway == nil {
		return errors.New("CommandGateway not config")
	}
	return nil

}

func (app *AppGateway) AddQueryHandler(queryHandler query.Handler) error {

	return app.queryGateway.Processor.AddQueryHandler(queryHandler)

}

func (app *AppGateway) AddCommandHandler(cmdHandler command.Handler) error {

	return app.commandGateway.Processor.AddCommandHandler(cmdHandler)

}

func (app *AppGateway) commandHandler(ctx *fiber.Ctx) {

	namespace := ctx.Params("namespace")
	commandtype := ctx.Params("commandtype")

	cmdName := namespace + "." + commandtype
	exists := app.commandGateway.Processor.ExistsCommandHandler(cmdName)
	if !exists {
		ctx.Status(fiber.StatusNotFound).Send(errors.New("command not exist: " + cmdName))
		return
	}

	cmd, errorParam := buildCommand(ctx, namespace, commandtype)

	if errorParam != nil {
		ctx.Status(fiber.StatusInternalServerError).Send(errorParam)
		return
	}
	errorCommand := app.commandGateway.Send(cmdName, cmd)

	if errorCommand != nil {
		ctx.Status(fiber.StatusBadRequest).Send(errorCommand)
		return
	}

	errorJson := ctx.JSON(command.Response{UUID: utils.UUID(), RequestUUID: cmd.UUID})
	fail(errorJson)

}

func (app *AppGateway) queryHandler(ctx *fiber.Ctx) {

	namespace := ctx.Params("namespace")
	queryType := ctx.Params("querytype")

	queryName := namespace + "." + queryType
	exists := app.queryGateway.Processor.ExistsQueryHandler(queryName)

	if !exists {
		ctx.Status(fiber.StatusNotFound).Send(errors.New("query not exist: " + queryName))
		return
	}

	queryRequest, errorParam := buildQueryRequest(ctx, namespace, queryType)

	if errorParam != nil {
		ctx.Status(fiber.StatusBadRequest).Send(errorParam)
		return
	}

	queryResponse := app.queryGateway.Query(queryName, queryRequest)
	if queryResponse.Success {
		ctx.Status(fiber.StatusOK)
		ctx.JSON(queryResponse)
		return
	} else {
		ctx.Status(fiber.StatusBadRequest)
		ctx.JSON(queryResponse)
		return
	}
}

func (app *AppGateway) status(ctx *fiber.Ctx) {

	ctx.Send("ON")

}

func buildQueryRequest(ctx *fiber.Ctx, namespace string, queryType string) (query.Resquest, error) {

	strFilterValue := ctx.Query("filter")
	mapFilter := query.BuildFilter(strFilterValue)
	fmt.Println(strFilterValue)

	strSortValue := ctx.Query("sort")
	mapSort := query.BuildSort(strSortValue)
	fmt.Println(strSortValue)

	strPageValue := ctx.Query("page")
	strSizeValue := ctx.Query("size")
	page, errPage := strconv.ParseInt(strPageValue, 10, 32)
	if errPage != nil {
		page = 0
	}
	size, errSize := strconv.ParseInt(strSizeValue, 10, 32)
	if errSize != nil {
		size = 10
	}

	queryRequest := query.Resquest{UUID: query.GenerateUUID(),
		Namespace: namespace,
		QueryType: queryType,
		Filter:    mapFilter,
		Sort:      mapSort,
		Page:      int32(page),
		Size:      int32(size)}

	return queryRequest, nil
}

func buildCommand(ctx *fiber.Ctx, namespace string, commandtype string) (command.Command, error) {

	var cmd = &command.Command{}
	property := &command.Property{}

	errorParser := ctx.BodyParser(property)
	if errorParser != nil {
		return *cmd, errorParser
	}
	cmd = &command.Command{Namespace: namespace, CommandType: commandtype, Property: *property, UUID: utils.UUID()}
	return *cmd, nil
}

func fail(err error) {
	if err != nil {
		log.Fatal("failed:", err.Error())
	}
}
