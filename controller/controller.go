package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type controllerRegister struct{}

type ControllerRegisterInterface interface {
	TrackingController(router *gin.RouterGroup)
}

func NewController() ControllerRegisterInterface {
	return &controllerRegister{}
}

func (*controllerRegister) TrackingController(router *gin.RouterGroup) {
	router.POST("/weight", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "fit-000",
		})
	})
}
