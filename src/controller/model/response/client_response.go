package response

import (
	"time"
)

type ClientResponse struct {
	Id                uint32
	CompanyName       string
	FantasyName       string
	Cnpj              string
	StateRegistration string
	Sector            string
	Cep               string
	Number            string
	Address           string
	City              string
	Neighborhood      string
	Complement        string
	State             string
	PersonName        string
	DateOfBirth       time.Time
	Cpf               string
	Gender            string
	TelephoneNumber   string
	CellphoneNumber   string
	Email             string
}