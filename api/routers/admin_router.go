package routers

import (
	"loan-tracker/api/controllers"
	"loan-tracker/api/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRouter(r *gin.Engine, loanCtrl *controllers.LoanController, userCtrl *controllers.UserController, jwtService middleware.JWTService) {
	adminGroup := r.Group("/admin")
	adminGroup.Use(middleware.AdminMiddleware(jwtService))
	{
		adminGroup.GET("/loans", loanCtrl.ViewAllLoans)
		adminGroup.PATCH("/loans/:id/status", loanCtrl.UpdateLoanStatus)
		adminGroup.DELETE("/loans/:id", loanCtrl.DeleteLoan)
		adminGroup.GET("/logs", loanCtrl.ViewSystemLogs)

		adminGroup.GET("/getallusers", userCtrl.GetAllUsers)
		adminGroup.DELETE("/deleteuser/:id", userCtrl.DeleteUser)

	}
}
