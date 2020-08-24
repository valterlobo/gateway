package query

type Gateway struct {
	Processor *Processor
}

func (queryGateway *Gateway) Query(queryName string, queryRequest Resquest) Response {
	result := queryGateway.Processor.ProcessQueryHandler(queryName, queryRequest)
	return result
}
