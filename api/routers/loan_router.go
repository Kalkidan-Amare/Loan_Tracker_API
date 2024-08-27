package routers

import (
	"loan-tracker/api/controllers"
	"loan-tracker/api/middleware"
	"github.com/gin-gonic/gin"
)

func LoanRouter(r *gin.Engine, loanCtrl *controllers.LoanController, jwtService middleware.JWTService) {
	api := r.Group("/api")
	
	loanGroup := api.Group("/loans")
	loanGroup.Use(middleware.AuthMiddleware(jwtService))
	{
		loanGroup.POST("/", loanCtrl.ApplyForLoan)
		loanGroup.GET("/:id", loanCtrl.GetLoanStatus)
	}
}
