package query

type Handler interface {
	Handle(queryRequest Resquest) Response
	GetName() string
}
