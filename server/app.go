package server

import (
	"errors"
	"gateway/command"
	"gateway/myapp"
	"gateway/query"
	"github.com/gofiber/fiber"
)

type AppGateway struct {
	CommandGateway *command.Gateway
	QueryGateway   *query.Gateway
}

func NewAppGateway(api fiber.Router) {

	appGateway := AppGateway{}
	appGateway.ConfigQueryGateway()
	appGateway.AddQueryHandler()

	appGateway.ConfigCommandGateway()

	//ROUTER
	api.Post("/command/:namespace/:commandtype", appGateway.CommandHandler)
	api.Get("/query/:namespace/:querytype", appGateway.QueryHandler)

}
func (app *AppGateway) ConfigQueryGateway() error {

	processor := query.NewQueryProcessor()
	if processor == nil {
		return errors.New("QueryProcessor not config")
	}
	app.QueryGateway = &query.Gateway{Processor: processor}

	if app.QueryGateway == nil {
		return errors.New("QueryGateway not config")
	}

	return nil
}

func (app *AppGateway) ConfigCommandGateway() error {

	cmdProcessor := command.NewCommandProcessor()
	if cmdProcessor == nil {
		return errors.New("CMD Processor not config")
	}
	app.CommandGateway = &command.Gateway{Processor: cmdProcessor}

	if app.CommandGateway == nil {
		return errors.New("CommandGateway not config")
	}
	return nil

}

func (app *AppGateway) AddQueryHandler() error {

	return app.QueryGateway.Processor.AddQueryHandler(myapp.HelloQueryHandler{Database: "123 teste", Name: "hello.ola"})

}

func (app *AppGateway) CommandHandler(ctx *fiber.Ctx) {

	//namespace := ctx.Params("namespace")
	//commandtype := ctx.Params("commandtype")
	commandData := ctx.Body()

	ctx.JSON(commandData)

}

func (app *AppGateway) QueryHandler(ctx *fiber.Ctx) {

	namespace := ctx.Params("namespace")
	queryType := ctx.Params("querytype")

	queryName := namespace + "." + queryType
	exists := app.QueryGateway.Processor.ExistsQueryHandler(queryName)
	if !exists {
		ctx.Status(404).Send(errors.New("query not exist: " + queryName))
		return
	}
	queryRequest, errorParam := buildQueryRequest(ctx, namespace, queryType)
	if errorParam != nil {
		ctx.Status(500).Send(errorParam)
		return
	}

	queryResponse := app.QueryGateway.Query(queryName, queryRequest)
	ctx.JSON(queryResponse)

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
