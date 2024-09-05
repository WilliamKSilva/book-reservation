package DTOs

type CreateUserRequestDTO struct {
	// @example {"name": "John Doe", "email": "john.doe@example.com", "password": "teste1234", "cpf": "23212332112", "birth_date": "2024-08-12"}
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CPF       string `json:"cpf"`
	BirthDate string `json:"birth_date"`
}

type CreateUserResponseDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CPF       string `json:"cpf"`
	BirthDate string `json:"birth_date"`
}

type FindUserByEmailResponseDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CPF       string `json:"cpf"`
	BirthDate string `json:"birth_date"`
}
