package myapp

import (
	"errors"
	"fmt"
	"gateway/command"
)

type OlaCommandHandler struct {
	Name string
}

func (ola OlaCommandHandler) Handle(cmd command.Command) error {

	fmt.Println("INICIO OlaCommandHandler")
	fmt.Println(cmd.UUID)
	fmt.Println(cmd.CommandType)
	fmt.Println(cmd.Namespace)
	fmt.Println(cmd.Property)
	fmt.Println(cmd.Property.Get("lastName"))
	fmt.Println("FIM OlaCommandHandler")

	return nil
}

func (ola OlaCommandHandler) GetName() string {

	return ola.Name
}

func (ola OlaCommandHandler) Validate(cmd command.Command) error {

	//fmt.Println(cmd.Property)

	firtname := cmd.Property.Get("firstName")

	lastname := cmd.Property.Get("lastName")
	if firtname == nil {

		return errors.New("firtname e requerido")
	}

	if lastname == nil {

		return errors.New("lastname e requerido")
	}

	//fmt.Println("VALIDADO ? ")

	return nil
}

/*
func (ola OlaCommandHandler) GetSchema()  []byte {

	schemaRaw := []byte(`{
    "$id": "https://qri.io/schema/",
    "$comment" : "sample comment",
    "title": "Person",
    "type": "object",
    "properties": {
        "firstName": {
            "type": "string"
        },
        "lastName": {
            "type": "string"
        },
        "age": {
            "description": "Age in years",
            "type": "integer",
            "minimum": 0
        },
        "friends": {
          "type" : "array",
          "items" : { "title" : "REFERENCE", "$ref" : "#" }
        }
    },
    "required": ["firstName", "lastName"]
  }`)
	return schemaRaw
}*/
