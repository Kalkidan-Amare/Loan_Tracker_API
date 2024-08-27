package usecases

import (
	"errors"
	"loan-tracker/api/middleware"
	"loan-tracker/domain"
	"loan-tracker/internal"
	// "loan-tracker/repositories"
	// "loan-tracker/utils"
)

type RegisterUsecase struct {
	userRepo domain.UserRepositoryInterface
}

func NewRegisterUsecase(userRepo domain.UserRepositoryInterface) *RegisterUsecase {
	return &RegisterUsecase{userRepo}
}

func (u *RegisterUsecase) Register(input domain.RegisterInput) (string, error) {
	// Check if user already exists
	existingUser, _ := u.userRepo.GetUserByEmail(input.Email)
	if existingUser != nil {
		return "", errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := middleware.HashPassword(input.Password)
	if err != nil {
		return "", err
	}

	// Generate verification token
	verificationToken, err := internal.GenerateRandomToken(32)
	if err != nil {
		return "", err
	}

	// Create user
	user := &domain.User{
		Name: 			input.Name,
		Email:             input.Email,
		Password:          hashedPassword,
		Role:              "user",
		VerificationToken: verificationToken,
		IsVerified:        false,
	}
	err = u.userRepo.CreateUser(user)
	if err != nil {
		return "", err
	}

	return verificationToken, nil
}
