package myapp

import (
	"fmt"
	"gateway/query"
	"lobo.tech/contatogo/queryutils"
	"lobo.tech/contatogo/repository"
)

type ContatoQueryHandler struct {
	ContatoRepository *repository.ContatoXRepository
	Name              string
}

func NewContatoQueryHandler() ContatoQueryHandler {

	contatoXRepository := &repository.ContatoXRepository{DBX: repository.NewDBX()}
	var contatoQueryHandler = ContatoQueryHandler{Name: "contato.search", ContatoRepository: contatoXRepository}
	return contatoQueryHandler
}

func (contatoQuery ContatoQueryHandler) Handle(queryRequest query.Resquest) query.Response {

	fmt.Println("HelloQueryHandlerHelloQueryHandler")
	fmt.Println(queryRequest)
	fmt.Println(queryRequest.Sort)
	fmt.Println(queryRequest.Filter)
	//queryRequest.Filter
	var queryResponse query.Response

	pageable := queryutils.PageableData{Page: queryRequest.Page , Size: queryRequest.Size}
	contatosPage, err := contatoQuery.ContatoRepository.GetByNamePage("lobo" , pageable)
	if err != nil {
		queryResponse = query.Response{UUID: query.GenerateUUID(), RequestUUID: queryRequest.UUID, Error: []error{err}, Success: false}
	} else {
		queryResponse = query.Response{UUID: query.GenerateUUID(), RequestUUID: queryRequest.UUID, Data: contatosPage, Success: true}
	}
	return queryResponse
}

func (contatoQuery ContatoQueryHandler) GetName() string {

	return contatoQuery.Name
}
