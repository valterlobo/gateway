package query

import "fmt"

type Gateway struct {
	Config         string
	Processor *Processor
}




func (queryGateway *Gateway) Query(queryName string, queryRequest Resquest) Response {

	fmt.Println("\n /COMMAND ENGINE" + queryGateway.Config)
	result := queryGateway.Processor.ProcessQueryHandler(queryName, queryRequest)
	fmt.Println("\n /QUERY ENGINE" + queryGateway.Config)
	fmt.Println("\n /QUERY ENGINE RESULT ", result)
	return result
}


