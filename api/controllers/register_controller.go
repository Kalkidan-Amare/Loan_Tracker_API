package controllers

import (
	"loan-tracker/domain"
	"loan-tracker/internal"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterController struct {
	userUsecase domain.RegisterUsecaseInterface
	mailConfig      internal.SMTPConfig
}

func NewRegisterController(userUsecase domain.RegisterUsecaseInterface, mailConfig internal.SMTPConfig) *RegisterController {
	return &RegisterController{userUsecase, mailConfig}
}

func (ctrl *RegisterController) Register(c *gin.Context) {
	var input domain.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	verificationToken, err := ctrl.userUsecase.Register(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send verification email
	verificationLink := "http://localhost:8080/users/verify-email?token=" + verificationToken + "&email=" + input.Email
	body := "Please verify your email by clicking the following link: " + verificationLink
	err = internal.SendEmail(input.Email, "Email Verification", body, ctrl.mailConfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send verification email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully. Please check your email to verify your account."})
}
