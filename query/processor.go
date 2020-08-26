package query

import (
	"fmt"
	"github.com/gofiber/utils"
)

type Processor struct {
	handlers map[string]Handler
}

func NewQueryProcessor() *Processor {
	handlers := make(map[string]Handler)
	return &Processor{handlers: handlers}
}

func (processor Processor) AddQueryHandler(queryHandler Handler) error {

	name := queryHandler.GetName()
	if processor.handlers[name] != nil {
		return fmt.Errorf("HANDLE JA ADICIONADO [%s]", name)
	}
	processor.handlers[name] = queryHandler
	return nil

}

func (processor Processor) ProcessQueryHandler(name string, query Resquest) Response {

	queryHandler := processor.handlers[name]
	if queryHandler == nil {
		var errorMesg = fmt.Sprintf("QueryHandler:[%s] NÃO DEFINIDO", name)
		return Response{
			UUID:    GenerateUUID(),
			Success: false,
			ErrorMessage:   errorMesg,
		}
	}
	return queryHandler.Handle(query)
}

func (processor Processor) ExistsQueryHandler(name string) bool {

	queryHandler := processor.handlers[name]
	if queryHandler == nil {
		return false
	}
	return true
}

func GenerateUUID() string {

	return utils.UUID()
}
