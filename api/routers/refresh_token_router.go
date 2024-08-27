package routers

import (
	"github.com/gin-gonic/gin"
	"loan-tracker/api/controllers"
)

func RefreshTokenRouter(r *gin.Engine, tokenCtrl *controllers.RefreshTokenController) {
	// Token management route
	r.POST("/refreshtoken", tokenCtrl.RefreshToken)
}
