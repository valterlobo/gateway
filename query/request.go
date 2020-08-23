package query



type Resquest struct {
	UUID      string
	Namespace string
	QueryType string
	Page      int64
	Size      int64
	Filter    []FilterParameter
	Sort      []SortParameter
	Domain    string
	Token     string
}
