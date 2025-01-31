package entity

import (
	"errors"

	"github.com/Moreno1337/api-ims/internal/clients/domain/validators"
	valueobject "github.com/Moreno1337/api-ims/internal/clients/domain/value-object"
)

type Client interface {
}

type client struct {
	id                valueobject.ID
	companyName       valueobject.CompanyName
	fantasyName       valueobject.FantasyName
	cnpj              valueobject.Cnpj
	stateRegistration valueobject.StateRegistration
	sector            valueobject.Sector
	cep               valueobject.Cep
	number            valueobject.Number
	address           valueobject.Address
	city              valueobject.City
	neighborhood      valueobject.Neighborhood
	complement        valueobject.Complement
	state             valueobject.State
	personName        valueobject.PersonName
	dateOfBirth       valueobject.DateOfBirth
	cpf               valueobject.Cpf
	gender            valueobject.Gender
	telephoneNumber   valueobject.TelephoneNumber
	cellphoneNumber   valueobject.CellphoneNumber
	email             valueobject.Email
}

func NewClient(
	id valueobject.ID,
	companyName valueobject.CompanyName,
	fantasyName valueobject.FantasyName,
	cnpj valueobject.Cnpj,
	stateRegistration valueobject.StateRegistration,
	sector valueobject.Sector,
	cep valueobject.Cep,
	number valueobject.Number,
	address valueobject.Address,
	city valueobject.City,
	neighborhood valueobject.Neighborhood,
	complement valueobject.Complement,
	state valueobject.State,
	personName valueobject.PersonName,
	dateOfBirth valueobject.DateOfBirth,
	cpf valueobject.Cpf,
	gender valueobject.Gender,
	telephoneNumber valueobject.TelephoneNumber,
	cellphoneNumber valueobject.CellphoneNumber,
	email valueobject.Email,
) (Client, error) {

	err := validateData(
		companyName,
		cnpj,
		cep,
		number,
		address,
		city,
		neighborhood,
		state,
		personName,
		dateOfBirth,
		cpf,
		telephoneNumber,
		cellphoneNumber,
		email,
	)
	if err != nil {
		return nil, err
	}

	return &client{
		id:                id,
		companyName:       companyName,
		fantasyName:       fantasyName,
		cnpj:              cnpj,
		stateRegistration: stateRegistration,
		sector:            sector,
		cep:               cep,
		number:            number,
		address:           address,
		city:              city,
		neighborhood:      neighborhood,
		complement:        complement,
		state:             state,
		personName:        personName,
		dateOfBirth:       dateOfBirth,
		cpf:               cpf,
		gender:            gender,
		telephoneNumber:   telephoneNumber,
		cellphoneNumber:   cellphoneNumber,
		email:             email,
	}, nil
}

func validateData(
	companyName valueobject.CompanyName,
	cnpj valueobject.Cnpj,
	cep valueobject.Cep,
	number valueobject.Number,
	address valueobject.Address,
	city valueobject.City,
	neighborhood valueobject.Neighborhood,
	state valueobject.State,
	personName valueobject.PersonName,
	dateOfBirth valueobject.DateOfBirth,
	cpf valueobject.Cpf,
	telephoneNumber valueobject.TelephoneNumber,
	cellphoneNumber valueobject.CellphoneNumber,
	email valueobject.Email,
) error {
	if len(cpf) == 0 && len(cnpj) == 0 {
		return errors.New("client should have either a valid cnpj or cpf")
	}

	if len(cpf) > 0 {
		if err := validators.ValidateCpf(cpf); err != nil {
			return err
		}

		if len(personName) == 0 {
			return errors.New("person name can't be empty")
		}
	}

	if len(cnpj) > 0 {
		if err := validators.ValidateCnpj(cnpj); err != nil {
			return err
		}

		if len(companyName) == 0 {
			return errors.New("company name can't be empty")
		}
	}

	if err := validators.ValidateCEP(cep); err != nil {
		return err
	}

	if err := validators.ValidateAddress(address); err != nil {
		return err
	}

	if err := validators.ValidateNumber(number); err != nil {
		return err
	}

	if err := validators.ValidateCity(city); err != nil {
		return err
	}

	if err := validators.ValidateNeighborhood(neighborhood); err != nil {
		return err
	}

	if err := validators.ValidateState(state); err != nil {
		return err
	}

	if len(dateOfBirth) > 0 {
		if err := validators.ValidateDateOfBirth(dateOfBirth); err != nil {
			return err
		}
	}

	if len(telephoneNumber) > 0 {
		if err := validators.ValidateTelephoneNumber(telephoneNumber); err != nil {
			return err
		}
	}

	if len(cellphoneNumber) > 0 {
		if err := validators.ValidateCellphoneNumber(cellphoneNumber); err != nil {
			return err
		}
	}

	if len(email) > 0 {
		if err := validators.ValidateEmail(email); err != nil {
			return err
		}
	}

	return nil
}
