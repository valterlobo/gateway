package query

type Response struct {
	UUID        string
	Data        interface{}
	RequestUUID string
	Success     bool
	Error       []error
}
