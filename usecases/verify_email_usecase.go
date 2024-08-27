package usecases

import (
	"errors"
	"loan-tracker/domain"
)

type VerifyEmailUsecase struct {
	userRepo domain.UserRepositoryInterface
}

func NewVerifyEmailUsecase(userRepo domain.UserRepositoryInterface) *VerifyEmailUsecase {
	return &VerifyEmailUsecase{userRepo}
}

func (u *VerifyEmailUsecase) VerifyEmail(token, email string) error {
	user, err := u.userRepo.GetUserByEmail(email)
	if err != nil || user == nil || user.VerificationToken != token {
		return errors.New("invalid token or email")
	}

	user.IsVerified = true
	user.VerificationToken = ""
	return u.userRepo.UpdateUser(user.Name,user)
}
