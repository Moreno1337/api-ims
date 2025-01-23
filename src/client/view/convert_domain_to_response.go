package view

import (
	"github.com/Moreno1337/api-ims/src/client/controller/model/response"
	"github.com/Moreno1337/api-ims/src/client/model"
)

func ConvertDomainToResponse(clientDomain model.ClientDomainInterface) response.ClientResponse {
	cd := clientDomain.GetClientDomain()

	return response.ClientResponse{
		Id:                cd.Id,
		CompanyName:       cd.CompanyName,
		FantasyName:       cd.FantasyName,
		Cnpj:              cd.Cnpj,
		StateRegistration: cd.StateRegistration,
		Sector:            cd.Sector,
		Cep:               cd.Cep,
		Number:            cd.Number,
		Address:           cd.Address,
		City:              cd.City,
		Neighborhood:      cd.Neighborhood,
		Complement:        cd.Complement,
		State:             cd.State,
		PersonName:        cd.PersonName,
		DateOfBirth:       cd.DateOfBirth,
		Cpf:               cd.Cpf,
		Gender:            cd.Gender,
		TelephoneNumber:   cd.TelephoneNumber,
		CellphoneNumber:   cd.CellphoneNumber,
		Email:             cd.Email,
	}
}
