package query



type Resquest struct {
	UUID      string
	Namespace string
	QueryType string
	Page      int32
	Size      int32
	Filter    map[string]FilterParameter
	Sort      []SortParameter
	Domain    string
	Token     string
}
