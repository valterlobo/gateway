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

	fmt.Println("ContatoQueryHandler")
	fmt.Println(queryRequest)
	fmt.Println(queryRequest.Sort)
	fmt.Println(queryRequest.Filter)
	//queryRequest.Filter
	pageable := queryutils.PageableData{Page: queryRequest.Page, Size: queryRequest.Size , Sort: queryRequest.Sort}
	var queryResponse query.Response
	if queryRequest.Filter == nil || queryRequest.Filter["nome"] == nil {
		return query.BuildQueryReponseError("Filter/nome não enviado", queryRequest)
	}
	searchName := queryRequest.Filter["nome"]["value"]

	if searchName == "" {
		return query.BuildQueryReponseError("Nome não enviado", queryRequest)
	}
	contatosPage, err := contatoQuery.ContatoRepository.GetByNamePage(searchName, pageable)
	if err != nil {
		queryResponse = query.BuildQueryReponseError(err.Error(), queryRequest)
	} else {
		queryResponse = query.BuildQueryReponseSucess(contatosPage, queryRequest)
	}
	return queryResponse
}

func (contatoQuery ContatoQueryHandler) GetName() string {

	return contatoQuery.Name
}
