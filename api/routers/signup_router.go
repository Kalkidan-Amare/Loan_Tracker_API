package routers

import (
	"github.com/gin-gonic/gin"
	"loan-tracker/api/controllers"
)

func AuthRouter(r *gin.Engine, registerCtrl *controllers.RegisterController, verifyEmailCtrl *controllers.VerifyEmailController) {
	// Authentication routes
	r.POST("/users/register", registerCtrl.Register)
	r.GET("/users/verify-email", verifyEmailCtrl.VerifyEmail)
}
