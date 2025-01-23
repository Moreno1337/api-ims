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
	id                uint32
	companyName       string
	fantasyName       string
	cnpj              string
	stateRegistration string
	sector            string
	cep               string
	number            string
	address           string
	city              string
	neighborhood      string
	complement        string
	state             string
	personName        string
	dateOfBirth       time.Time
	cpf               string
	gender            string
	telephoneNumber   string
	cellphoneNumber   string
	email             string
}

func (c clientDomain) GetClientDomain() clientDomain {
	return c
}
