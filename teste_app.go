package main

import (
	"gateway/myapp"
	"gateway/server"
)

func main() {

	appGateway := &server.AppGateway{}
	appGateway = server.NewAppGateway()
	appGateway.AddQueryHandler(myapp.HelloQueryHandler{Database: "123 teste", Name: "hello.ola"})
	appGateway.AddQueryHandler(myapp.NewContatoQueryHandler())
	appGateway.AddCommandHandler(myapp.OlaCommandHandler{Name: "hello.add"})
	appGateway.AddCommandHandler(myapp.NewAddContatoCommandHandler())
	appGateway.Start()

}
