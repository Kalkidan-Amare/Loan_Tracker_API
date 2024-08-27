package usecases

import (
	"errors"
	"fmt"
	"loan-tracker/domain"
	"loan-tracker/internal"

	"time"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type OTPUsecase struct {
	otpRepository domain.OTPRepositoryInterface
	userRepo      domain.UserRepositoryInterface
	// infrastruct   internal.EmailService
}

func NewOTPUsecase(or domain.OTPRepositoryInterface, ur domain.UserRepositoryInterface) *OTPUsecase {
	return &OTPUsecase{
		otpRepository: or,
		userRepo:      ur,
	}
}

func (ou *OTPUsecase) GenerateAndSendOTP(user *domain.User) error {
	if !internal.ValidateEmail(user.Email) {
		return errors.New("invalid email")
	}
	if !internal.ValidatePassword(user.Password) {
		return errors.New("password must be at least 8 characters long")
	}

	// Generate OTP
	otp := internal.GenerateOTP(6)
	storeOtp := domain.OTP{
		Otp:       otp,
		Email:     user.Email,
		Username:  user.Name,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7),

		Password: user.Password,
		Role:     user.Role,
	}

	// Store OTP in the database
	err := ou.otpRepository.StoreOTP(&storeOtp)
	if err != nil {
		return err
	}

	// Send OTP via email
	err = internal.SendOTPEmail(user.Email, otp)
	if err != nil {
		return err
	}

	return nil
}

// verification endpoint
func (ou *OTPUsecase) VerifyOTP(email, otp string) (*domain.OTP,error) {
        fmt.Println("email: ", email)
        storeOtp, err := ou.otpRepository.GetOTPByEmail(email)
        if err != nil {
                return nil, err
        }

        if time.Now().After(storeOtp.ExpiresAt) {
                return nil, errors.New("otp expired")
        }

        if storeOtp.Otp != otp {
                return nil, errors.New("invalid OTP")
        }

        // Clean up the used OTP
        if err =ou.otpRepository.DeleteOTPByEmail(email); err != nil {
                return nil, errors.New("couldn't delete OTP")
        }

        return storeOtp, nil
}

func (ou *OTPUsecase) ForgotPassword(email string) error {
  // panic(email)
  if !internal.ValidateEmail(email) {
          return  errors.New("invalid email")
  }

	user, err := ou.userRepo.GetUserByEmail(email)
	if err != nil {
		return errors.New("email not found")
	}

	// Generate OTP and store it in the database
	otp := internal.GenerateOTP(6)
	storeOtp := domain.OTP{
		Otp:       otp,
		Email:     user.Email,
		Username:  user.Name,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7),

		Password: user.Password,
		Role:     user.Role,
	}
	err = ou.otpRepository.StoreOTP(&storeOtp)
	if err != nil {
		return err
	}

	// Send OTP via email
	err = internal.SendOTPEmail(user.Email, otp)
	if err != nil {
		return err
	}

	return nil
}
