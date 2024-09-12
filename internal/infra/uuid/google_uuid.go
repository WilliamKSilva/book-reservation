package uuid

import (
	services_errors "github.com/WilliamKSilva/book-reservation/internal/services/errors"
	googleUuid "github.com/google/uuid"
)

type GoogleUuidService struct{}

func generate(uuid string, err error) (string, error) {
	defer func() {
		r := recover()
		if r != nil {
			err = &services_errors.InternalServerError{}
			uuid = ""
		}
	}()

	uuid = googleUuid.NewString()

	return uuid, nil
}

func (googleUuidService *GoogleUuidService) Generate() (string, error) {
	uuid, err := generate("", nil)
	return uuid, err
}
