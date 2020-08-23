package query


type Response struct {

	UUID        string
	Header      []string
	Page        int64
	Size        int64
	Records     []map[string]interface{}
	RequestUUID string
	Failed      bool
	Error       []error
}