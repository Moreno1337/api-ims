package request

import (
	"time"
)

type ClientRequest struct {
	Id                uint32    `json:"id" binding:"min=1"`
	CompanyName       string    `json:"company_name" binding:"min=2,max=200"`
	FantasyName       string    `json:"fantasy_name" binding:"max=200"`
	Cnpj              string    `json:"cnpj" binding:"min=18,max=18"`
	StateRegistration string    `json:"state_registration" binding:"max=50"`
	Sector            string    `json:"sector" binding:"max=100"`
	Cep               string    `json:"cep" binding:"required,min=9,max=9"`
	Number            string    `json:"number" binding:"required,min=1,max=10"`
	Address           string    `json:"address" binding:"required,min=2,max=250"`
	City              string    `json:"city" binding:"required,min=2,max=100"`
	Neighborhood      string    `json:"neighborhood" binding:"required,min=2,max=100"`
	Complement        string    `json:"complement" binding:"max=150"`
	State             string    `json:"state" binding:"required,min=2,max=2"`
	PersonName        string    `json:"person_name" binding:"min=2,max=200"`
	DateOfBirth       time.Time `json:"date_of_birth"`
	Cpf               string    `json:"cpf" binding:"min=14,max=14"`
	Gender            string    `json:"gender" binding:"min=2,max=50"`
	TelephoneNumber   string    `json:"telephone_number" binding:"min=14,max=14"`
	CellphoneNumber   string    `json:"cellphone_number" binding:"min=16,max=16"`
	Email             string    `json:"email" binding:"email,max=1000"`
}
