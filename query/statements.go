package query

import "errors"

type  SortParameter struct {

	Field string
	Direction Direction //asc desc
}

type FilterParameter struct {

	Field string
	Operator Operator
	Value interface{}
}

type Direction string

const (
	ASC Direction = "ASC"
	DESC Direction = "DESC"
)

func (direction Direction) IsValid() error {
	switch direction {
	case ASC, DESC:
		return nil
	}
	return errors.New("Invalid Direction type")
}


//const ASC Sort = Sort{Direction: "ASC"}


/*
   SORT
   direction -
   properties -
   ex:   properties => direction = name => asc
  Filter
       properties =>  operator  => value
*/

type Operator string

const (
	EQ Operator = "="
	GT Operator = ">"
	GTE Operator = ">="
	LT Operator = "<"
	LTE Operator = "<="
	NE Operator = "<>"
)

/*
https://postgres.rest/query-statements/
	$eq 	Matches values that are equal to a specified value.
	$gt 	Matches values that are greater than a specified value.
	$gte 	Matches values that are greater than or equal to a specified value.
	$lt 	Matches values that are less than a specified value.
	$lte 	Matches values that are less than or equal to a specified value.
	$ne 	Matches all values that are not equal to a specified value.
	$in 	Matches any of the values specified in an array.
	$nin 	Matches none of the values specified in an array.
	$null 	Matches if field is null.
	$notnull 	Matches if field is not null.
	$true 	Matches if field is true.
	$nottrue 	Matches if field is not true.
	$false 	Matches if field is false.
	$notfalse 	Matches if field is not false.
	$like 	Matches always cover the entire string.
	$ilike 	Matches case-insensitive always cover the entire string.
*/