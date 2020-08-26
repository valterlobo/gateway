package myapp

import (
	"errors"
	"gateway/command"
	"lobo.tech/contatogo/model"
	"lobo.tech/contatogo/repository"
)

type AddContatoCommandHandler struct {
	Name              string
	ContatoRepository *repository.ContatoXRepository
}

//NewAddContatoCommandHandler
func NewAddContatoCommandHandler() AddContatoCommandHandler {

	contatoXRepository := &repository.ContatoXRepository{DBX: repository.NewDBX()}
	var addContatoCommandHandler = AddContatoCommandHandler{Name: "contato.add", ContatoRepository: contatoXRepository}
	return addContatoCommandHandler
}

func (addContatoHandler AddContatoCommandHandler) Handle(cmd command.Command) error {

	contato := commandToStruct(cmd)
	_, error := addContatoHandler.ContatoRepository.Save(contato)
	return error
}

func (addContatoHandler AddContatoCommandHandler) GetName() string {

	return addContatoHandler.Name
}

func (addContatoHandler AddContatoCommandHandler) Validate(cmd command.Command) error {

	//fmt.Println(cmd.Property)
	/*
		ID      int64  `json:"id"`
		Nome    string `json:"nome"  validate:"required"`
		Celular string `json:"celular" validate:"required"`
		Email   string `json:"email"`
		DateCreate time.Time `json:"dt_create"`*/

	nome := cmd.Property.Get("Nome")
	celular := cmd.Property.Get("Celular")
	email := cmd.Property.Get("Email")
	if nome == nil {

		return errors.New("nome e requerido")
	}

	if celular == nil {

		return errors.New("celular e requerido")
	}

	if email == nil {

		return errors.New("email e requerido")
	}

	return nil
}

func commandToStruct(cmd command.Command) model.Contato {

	return model.Contato{
		Nome:    cmd.Property.Get("Nome").(string),
		Celular: cmd.Property.Get("Celular").(string),
		Email:   cmd.Property.Get("Email").(string),
	}

}
