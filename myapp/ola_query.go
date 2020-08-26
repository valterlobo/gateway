package myapp

import (
	"fmt"
	"gateway/query"
)

type HelloQueryHandler struct {
	Database string
	Name     string
}

func (ola HelloQueryHandler) Handle(queryRequest query.Resquest) query.Response {

	ola.Database = "TESTE"
	fmt.Println("HelloQueryHandler")
	fmt.Println(queryRequest)
	fmt.Println(queryRequest.Sort)
	fmt.Println(queryRequest.Filter)

	var queryResponse = query.Response{UUID: "2332323" , RequestUUID: queryRequest.UUID , Success: true }
	return queryResponse
}

func (ola HelloQueryHandler) GetName()  string {

	return ola.Name
}




