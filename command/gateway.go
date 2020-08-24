package command


type Gateway struct {
	Processor *Processor
}



func (cmdGateway *Gateway) Send(cmdName string, cmd Command) error {

	return cmdGateway.Processor.ProcessCommandHandler(cmdName , cmd)
}
