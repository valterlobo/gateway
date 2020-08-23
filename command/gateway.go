package command

import "fmt"

/*
type Gateway interface {
	Send(commandName string, command interface{} ) error
}
*/

type Gateway struct {
	Config         string
	Processor *Processor
}



func (cmdGateway *Gateway) Send(cmdName string, cmd Command) error {

	fmt.Println("\n /COMMAND ENGINE" + cmdGateway.Config)
	return cmdGateway.Processor.ProcessCommandHandler(cmdName , cmd)
}
