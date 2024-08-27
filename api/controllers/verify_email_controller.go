package controllers

import (
	"net/http"
	"loan-tracker/domain"
	"github.com/gin-gonic/gin"
)

type VerifyEmailController struct {
	verifyEmailUsecase domain.VerifyEmailUsecaseInterface
}

func NewVerifyEmailController(verifyEmailUsecase domain.VerifyEmailUsecaseInterface) *VerifyEmailController {
	return &VerifyEmailController{verifyEmailUsecase}
}

func (ctrl *VerifyEmailController) VerifyEmail(c *gin.Context) {
	token := c.Query("token")
	email := c.Query("email")

	err := ctrl.verifyEmailUsecase.VerifyEmail(token, email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
}
