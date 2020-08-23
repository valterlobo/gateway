package main

import (
	"gateway/server"
	"github.com/gofiber/fiber"
)

func start(app *fiber.App) {

	server.NewAppGateway(app)

}

func stop(app *fiber.App) {

}

func main() {
	app := fiber.New()
	start(app)
	setupRoutes(app)
	app.Listen(3000)
	stop(app)

}

func setupRoutes(app *fiber.App) {

	// Routes
	app.Get("/", status)

}

func status(c *fiber.Ctx) {
	c.Send("ON")
}

/*
func commandHandler(ctx *fiber.Ctx) {

	//namespace := ctx.Params("namespace")
	//commandtype := ctx.Params("commandtype")
	commandData := ctx.Body()

	ctx.JSON(commandData)

}

func queryHandler(ctx *fiber.Ctx) {

	// namespace := ctx.Params("namespace")
	//queryType := ctx.Params("commandtype")

	ctx.JSON("Query")

}*/
