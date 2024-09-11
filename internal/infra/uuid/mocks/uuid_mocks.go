package uuid_mocks

type MockedUuidService struct{}

func (uuidService *MockedUuidService) Generate() string {
	return `
		904cf2f4-eb41-4512-bce1-a1082cc674f2
	`
}

func NewMockedUuidService() *MockedUuidService {
	return &MockedUuidService{}
}
