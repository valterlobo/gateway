package server

import (
	"errors"
	"gateway/command"
	"gateway/query"
	"github.com/gofiber/fiber"
	"github.com/gofiber/utils"
	"log"
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

	errorJson := ctx.JSON(Response{UUID: utils.UUID(), RequestUUID: cmd.UUID})
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
	errorJson := ctx.JSON(queryResponse)
	fail(errorJson)
}

func (app *AppGateway) status(ctx *fiber.Ctx) {

	ctx.Send("ON")

}

func buildQueryRequest(ctx *fiber.Ctx, namespace string, queryType string) (query.Resquest, error) {

	//TESTE
	sort1 := query.SortParameter{Field: "nome", Direction: query.DESC}
	sort2 := query.SortParameter{Field: "data_cadastro", Direction: query.ASC}
	filter1 := query.FilterParameter{Field: "nome", Operator: query.EQ, Value: "valter"}
	filter2 := query.FilterParameter{Field: "data_cadastro", Operator: query.NE, Value: "2020-01-20"}
	filter3 := query.FilterParameter{Field: "valor", Operator: query.GT, Value: "20800.67"}
	queryRequest := query.Resquest{UUID: "dsdsdsdsds", Namespace: namespace, QueryType: queryType,
		Sort: []query.SortParameter{sort1, sort2}, Filter: []query.FilterParameter{filter1, filter2, filter3}, Page: 1, Size: 5}

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
