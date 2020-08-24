package command


type Handler interface {
	Handle(cmd Command)  error
	GetName() string
	Validate(cmd Command) error
}