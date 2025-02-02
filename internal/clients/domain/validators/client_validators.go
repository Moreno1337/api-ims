package validators

import (
	"errors"
	"os"
	"regexp"
	"time"

	valueobject "github.com/Moreno1337/api-ims/internal/clients/domain/value-object"
)

var EMAIL_REGEX = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

func ValidateCpf(cpf valueobject.Cpf) error {
	if len(cpf) != 14 {
		return errors.New("invalid cpf")
	}

	return nil
}

func ValidateCnpj(cnpj valueobject.Cnpj) error {
	if len(cnpj) != 18 {
		return errors.New("invalid cnpj")
	}

	return nil
}

func ValidateCEP(cep valueobject.Cep) error {
	if len(cep) != 9 {
		return errors.New("invalid cep")
	}

	return nil
}

func ValidateAddress(address valueobject.Address) error {
	if len(address) == 0 {
		return errors.New("address can't be empty")
	}

	return nil
}

func ValidateNumber(number valueobject.Number) error {
	if len(number) == 0 {
		return errors.New("number can't be empty")
	}

	return nil
}

func ValidateCity(city valueobject.City) error {
	if len(city) == 0 {
		return errors.New("city can't be empty")
	}

	return nil
}

func ValidateNeighborhood(neighborhood valueobject.Neighborhood) error {
	if len(neighborhood) == 0 {
		return errors.New("neighborhood can't be empty")
	}

	return nil
}

func ValidateState(state valueobject.State) error {
	if len(state) == 0 {
		return errors.New("state can't be empty")
	}

	return nil
}

func ValidateDateOfBirth(dob valueobject.DateOfBirth) error {
	date, err := time.Parse("2006-01-02", string(dob))
	if err != nil {
		return errors.New("invalid date format (expected YYYY-MM-DD)")
	}

	referenceDate := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)

	if date.After(time.Now().UTC()) || date.Before(referenceDate) {
		return errors.New("invalid date of birth (impossible date of birth)")
	}

	return nil
}

func ValidateTelephoneNumber(p valueobject.TelephoneNumber) error {
	if len(p) != 14 {
		return errors.New("invalid telephone number (valid format: (00) 0000-0000)")
	}

	return nil
}

func ValidateCellphoneNumber(c valueobject.CellphoneNumber) error {
	if len(c) != 16 {
		return errors.New("invalid cellphone number (valid format: (00) 0 0000-0000")
	}

	return nil
}

func ValidateEmail(email valueobject.Email) error {
	emailRegex := regexp.MustCompile(os.Getenv(EMAIL_REGEX))

	for _, value := range email {
		if !emailRegex.MatchString(value) {
			return errors.New("the email list contains one or more invalid emails")
		}
	}

	return nil
}
