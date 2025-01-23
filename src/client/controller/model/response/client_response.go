package response

import (
	"time"
)

type ClientResponse struct {
	Id                uint32    `json:"id"`
	CompanyName       string    `json:"company_name"`
	FantasyName       string    `json:"fantasy_name"`
	Cnpj              string    `json:"cnpj"`
	StateRegistration string    `json:"state_registration"`
	Sector            string    `json:"sector"`
	Cep               string    `json:"cep"`
	Number            string    `json:"number"`
	Address           string    `json:"address"`
	City              string    `json:"city"`
	Neighborhood      string    `json:"neighborhood"`
	Complement        string    `json:"complement"`
	State             string    `json:"state"`
	PersonName        string    `json:"person_name"`
	DateOfBirth       time.Time `json:"date_of_birth"`
	Cpf               string    `json:"cpf"`
	Gender            string    `json:"gender"`
	TelephoneNumber   string    `json:"telephone_number"`
	CellphoneNumber   string    `json:"cellphone_number"`
	Email             string    `json:"email"`
}
