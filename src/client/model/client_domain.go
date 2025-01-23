package model

import (
	"time"

	"github.com/Moreno1337/api-ims/src/client/controller/model/request"
)

type ClientDomainInterface interface {
	GetClientDomain() clientDomain
}

func NewClientDomain(cr *request.ClientRequest) ClientDomainInterface {
	return &clientDomain{
		cr.Id,
		cr.CompanyName,
		cr.FantasyName,
		cr.Cnpj,
		cr.StateRegistration,
		cr.Sector,
		cr.Cep,
		cr.Number,
		cr.Address,
		cr.City,
		cr.Neighborhood,
		cr.Complement,
		cr.State,
		cr.PersonName,
		cr.DateOfBirth,
		cr.Cpf,
		cr.Gender,
		cr.TelephoneNumber,
		cr.CellphoneNumber,
		cr.Email,
	}
}

type clientDomain struct {
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

func (c clientDomain) GetClientDomain() clientDomain {
	return c
}
