package routers

import (
	"github.com/gin-gonic/gin"
	"loan-tracker/api/controllers"
	"loan-tracker/api/middleware"
)

func LogoutRouter(r *gin.Engine, logoutCtrl *controllers.LogoutController, jwtService middleware.JWTService) {
	// Logout route
	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware(jwtService))
	{
		auth.POST("/logout", logoutCtrl.Logout)
	}
}
