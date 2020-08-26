package query

type Response struct {
	UUID        string
	Data        interface{}
	RequestUUID string
	Success     bool
	ErrorMessage   string
}

func BuildQueryReponseError(errorMessage string , queryRequest Resquest ) Response {

	return Response{UUID: GenerateUUID(), RequestUUID: queryRequest.UUID, ErrorMessage: errorMessage, Success: false}
}

func BuildQueryReponseSucess(content interface{} , queryRequest Resquest ) Response {

	return Response{UUID: GenerateUUID(), RequestUUID: queryRequest.UUID, Data : content, Success: true}
}





