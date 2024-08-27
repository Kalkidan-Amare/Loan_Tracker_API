package usecases

// import (
// 	"errors"
// 	"loan-tracker/domain"
// 	"loan-tracker/repositories"
// 	"loan-tracker/utils"
// )

// type ResetPasswordUsecase interface {
// 	RequestPasswordReset(email string) (string, error)
// 	UpdatePassword(token, newPassword string) error
// }

// type resetPasswordUsecase struct {
// 	userRepo repositories.UserRepository
// }

// func NewResetPasswordUsecase(userRepo repositories.UserRepository) ResetPasswordUsecase {
// 	return &resetPasswordUsecase{userRepo}
// }

// type PasswordResetRequestInput struct {
// 	Email string `json:"email"`
// }

// type PasswordUpdateInput struct {
// 	Token       string `json:"token"`
// 	NewPassword string `json:"new_password"`
// }

// func (u *resetPasswordUsecase) RequestPasswordReset(email string) (string, error) {
// 	user, err := u.userRepo.GetUserByEmail(email)
// 	if err != nil || user == nil {
// 		return "", errors.New("user not found")
// 	}

// 	resetToken, err := utils.GenerateRandomToken(32)
// 	if err != nil {
// 		return "", err
// 	}

// 	user.ResetToken = resetToken
// 	err = u.userRepo.UpdateUser(user.Name, user)
// 	if err != nil {
// 		return "", err
// 	}

// 	return resetToken, nil
// }

// func (u *resetPasswordUsecase) UpdatePassword(token, newPassword string) error {
// 	user, err := u.userRepo.GetByResetToken(token)
// 	if err != nil || user == nil {
// 		return errors.New("invalid or expired token")
// 	}

// 	hashedPassword, err := utils.HashPassword(newPassword)
// 	if err != nil {
// 		return err
// 	}

// 	user.Password = hashedPassword
// 	user.ResetToken = ""
// 	return u.userRepo.Update(user)
// }
