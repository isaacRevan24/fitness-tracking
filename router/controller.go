package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type routerRegister struct{}

type RouterRegisterInterface interface {
	TrackingRouter(router *gin.RouterGroup)
}

func NewRouterRegister() RouterRegisterInterface {
	return &routerRegister{}
}

func (*routerRegister) TrackingRouter(router *gin.RouterGroup) {
	router.POST("/weight", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "fit-000",
		})
	})
}
