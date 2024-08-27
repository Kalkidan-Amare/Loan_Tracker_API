package domain

type RegisterInput struct {
	Name     string             `json:"name" validate:"required,min=2,max=100"`
	Email    string             `json:"email" validate:"required,email"`
	Password string             `json:"password"`
	Role 	 string 			`json:"role"`
}

type RegisterUsecaseInterface interface {
	Register(input RegisterInput) (string, error)
}