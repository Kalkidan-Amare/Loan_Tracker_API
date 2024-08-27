package route

import (
	"loan-tracker/api/controllers"
	// "loan-tracker/api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	registerCtrl *controllers.RegisterController,
	verifyEmailCtrl *controllers.VerifyEmailController,
	resetPasswordCtrl interface{},
	// resetPasswordCtrl *controllers.ResetPasswordController,
) *gin.Engine {
	r := gin.Default()

	// User routes
	r.POST("/users/register", registerCtrl.Register)
	r.GET("/users/verify-email", verifyEmailCtrl.VerifyEmail)
	// r.POST("/users/password-reset", resetPasswordCtrl.RequestPasswordReset)
	// r.POST("/users/update-password", resetPasswordCtrl.UpdatePassword)
	


	// adminRoutes := r.Group("/admin", middleware.AdminAuthMiddleware())
	// adminRoutes.GET("/users", adminCtrl.GetAllUsers)
	// adminRoutes.DELETE("/users/:id", adminCtrl.DeleteUser)

	return r
}
