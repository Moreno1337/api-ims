package request

import (
	"time"
)

type ClientRequest struct {
	Id                uint32
	CompanyName       string
	FantasyName       string
	Cnpj              string
	StateRegistration string
	Cep               string
	Sector            string
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
	CellPhoneNumber   string
	Email             string
}
