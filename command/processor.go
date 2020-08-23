package command

import "fmt"

type Processor struct {
	handlers map[string]Handler
}

func NewCommandProcessor() *Processor {
	handlers := make(map[string]Handler)
	return &Processor{handlers: handlers}
}

func (processor Processor) AddCommandHandler(commandHandler Handler) error {

	name := commandHandler.GetName()
	if processor.handlers[name] != nil {
		return fmt.Errorf("HANDLE JA ADICIONADO [%s]", name)
	}
	processor.handlers[name] = commandHandler
	return nil
}

func (processor Processor) ProcessCommandHandler(name string , cmd  Command) error {

	commandHandler := processor.handlers[name]
	if commandHandler == nil {
		var errorReturn = fmt.Errorf("CommandHandler:[%s] NÃO DEFINIDO", name)
		return errorReturn
	}
	return commandHandler.Handle(cmd)
}

func (processor Processor) ExistsCommandHandler(name string) bool {

	CommandHandler := processor.handlers[name]
	if CommandHandler == nil {
		return false
	}
	return true
}
