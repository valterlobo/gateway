package command


type Command struct {
	UUID        string
	Namespace   string
	CommandType string
	Data        interface{}
}
