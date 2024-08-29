package uuid

import "github.com/google/uuid"

type GoogleUUIDGenerator struct{}

func (uuidGenerator GoogleUUIDGenerator) Generate() string {
	uuid := uuid.NewString()

	return uuid
}
