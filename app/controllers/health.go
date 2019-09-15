package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "200 Health OK!")
}