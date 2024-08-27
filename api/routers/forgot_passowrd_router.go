package routers

import (
	"github.com/gin-gonic/gin"
	"loan-tracker/api/controllers"
)

func PasswordRouter(
	r *gin.Engine,
	forgotPasswordCtrl *controllers.ForgotPasswordController,
	// resetPasswordCtrl *controllers.ResetPasswordController,
) {

	r.POST("/forgotpassword", forgotPasswordCtrl.ForgotPassword)
	r.POST("/verfiyforgotpassword", forgotPasswordCtrl.VerifyForgotOTP)
	// r.POST("/users/password-reset", resetPasswordCtrl.RequestPasswordReset)
	// r.POST("/users/update-password", resetPasswordCtrl.UpdatePassword)
}
