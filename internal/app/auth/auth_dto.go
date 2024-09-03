package auth

type LoginUser struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CPF       string `json:"cpf"`
	BirthDate string `json:"birth_date"`
}

type LoginResponseDTO struct {
	User        LoginUser
	AccessToken JwtToken
}

type RegisterRequestDTO struct {
	Name      string `json:"name" example:"John Doe"`
	Email     string `json:"email" example:"john.doe@example.com"`
	Password  string `json:"password" example:"password123"`
	CPF       string `json:"cpf" example:"32212276723"`
	BirthDate string `json:"birth_date" example:"2024-08-15"`
}

type RegisterResponseDTO struct {
	User        RegisterUser
	AccessToken JwtToken
}

type RegisterUser struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CPF       string `json:"cpf"`
	BirthDate string `json:"birth_date"`
}
