package routers

import (
	"github.com/bksides/rumblr-go/app/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.GET("/health", controllers.HealthHandler)
}