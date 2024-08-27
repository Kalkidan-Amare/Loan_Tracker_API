package routers

import (
	"github.com/gin-gonic/gin"
	"loan-tracker/api/controllers"
	"loan-tracker/api/middleware"
)

func UserRouter(r *gin.Engine, userCtrl *controllers.UserController, jwtService middleware.JWTService) {
	// User routes
	r.POST("/login", userCtrl.Login)

	// Authenticated user routes
	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware(jwtService))
	{
		auth.GET("/getallusers", middleware.AdminMiddleware(jwtService), userCtrl.GetAllUsers)
		auth.DELETE("/deleteuser/:id", middleware.AdminMiddleware(jwtService), userCtrl.DeleteUser)
	}
}
