package controllers

// import (
// 	"net/http"
// 	"loan-tracker-api/usecase"
// 	"loan-tracker-api/infrastructure"
// 	"github.com/gin-gonic/gin"
// )

// type ResetPasswordController struct {
// 	resetPasswordUsecase usecase.ResetPasswordUsecase
// 	mailConfig           infrastructure.SMTPConfig
// }

// func NewResetPasswordController(resetPasswordUsecase usecase.ResetPasswordUsecase, mailConfig infrastructure.SMTPConfig) *ResetPasswordController {
// 	return &ResetPasswordController{resetPasswordUsecase, mailConfig}
// }

// func (ctrl *ResetPasswordController) RequestPasswordReset(c *gin.Context) {
// 	var input usecase.PasswordResetRequestInput
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	resetToken, err := ctrl.resetPasswordUsecase.RequestPasswordReset(input.Email)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Send password reset email
// 	resetLink := "http://localhost:8080/users/password-reset?token=" + resetToken + "&email=" + input.Email
// 	body := "Please reset your password by clicking the following link: " + resetLink
// 	err = infrastructure.SendEmail(input.Email, "Password Reset", body, ctrl.mailConfig)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send password reset email"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Password reset link sent. Please check your email."})
// }

// func (ctrl *ResetPasswordController) UpdatePassword(c *gin.Context) {
// 	var input usecase.PasswordUpdateInput
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	err := ctrl.resetPasswordUsecase.UpdatePassword(input.Token, input.NewPassword)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
// }
