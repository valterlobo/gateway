package query

import (
	"fmt"
	"strings"
)

type Resquest struct {
	UUID      string
	Namespace string
	QueryType string
	Page      int32
	Size      int32
	Filter    map[string]map[string]string
	Sort      map[string]string
	Domain    string
	Token     string
}


func BuildFilter(strFilter string ) map[string]map[string]string {

	strParts :=  strings.Split(strFilter , "," )
	mapFilter := make(map[string]map[string]string)
	for i := range strParts {
		operator := strings.Split( strParts[i]  , "." )
		fmt.Println(operator)
		if len(operator) == 3 {
			mapFilter[operator[0]] = map[string]string { "field" : operator[0] , "operator" : operator[1], "value" : operator[2]}
		}
	}

	return mapFilter
}



func BuildSort(strSort string ) map[string]string {

	strParts :=  strings.Split(strSort , "," )
	mapSort := make(map[string]string)
	for i := range strParts {
		operator := strings.Split( strParts[i]  , "." )
		fmt.Println(operator)
		if len(operator) == 2 {
			mapSort[operator[0]] = operator[1]
		}
	}
	return mapSort
}