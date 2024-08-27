package domain

type EmailRequest struct {
    Email string `json:"email" binding:"required"`
}

type VerifyEmailUsecaseInterface interface {
	VerifyEmail(token, email string) error
}