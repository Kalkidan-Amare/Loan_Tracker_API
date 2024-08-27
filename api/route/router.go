package route

import (
	"loan-tracker/api/controllers"
	"loan-tracker/api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	userCtrl *controllers.UserController,
	registerCtrl *controllers.RegisterController,
	verifyEmailCtrl *controllers.VerifyEmailController,
	refreshTokenCtrl *controllers.RefreshTokenController,
	forgotPasswordCtrl *controllers.ForgotPasswordController,
	logoutCtrl *controllers.LogoutController,
	resetPasswordCtrl interface{},
	// resetPasswordCtrl *controllers.ResetPasswordController,
	jwtService middleware.JWTService,
) *gin.Engine {
	r := gin.Default()
	adminMiddleware := middleware.AdminMiddleware(jwtService)

	// User routes
	r.POST("/users/register", registerCtrl.Register)
	r.GET("/users/verify-email", verifyEmailCtrl.VerifyEmail)
	// r.POST("/users/password-reset", resetPasswordCtrl.RequestPasswordReset)
	// r.POST("/users/update-password", resetPasswordCtrl.UpdatePassword)
	r.POST("/login", userCtrl.Login)
	r.POST("/refreshtoken", refreshTokenCtrl.RefreshToken)
	r.POST("/forgotpassword", forgotPasswordCtrl.ForgotPassword)
	r.POST("/verfiyforgotpassword", forgotPasswordCtrl.VerifyForgotOTP)

	// Authenticated routes
	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware(jwtService))
	{
		auth.POST("/logout", logoutCtrl.Logout)

		// Admin-specific
		auth.GET("/getallusers", adminMiddleware, userCtrl.GetAllUsers)
		auth.DELETE("/deleteuser/:id", adminMiddleware, userCtrl.DeleteUser)
	}
	

	
	return r
}
