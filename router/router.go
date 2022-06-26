package router

import (
	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/fitness-tracking/handler"
)

var (
	trackingHandler handler.TrackingHandlerInterface = handler.NewTrackingHandler()
)

type routerRegister struct{}

type RouterRegisterInterface interface {
	TrackingRouter(router *gin.RouterGroup)
}

func NewRouterRegister() RouterRegisterInterface {
	return &routerRegister{}
}

func (*routerRegister) TrackingRouter(router *gin.RouterGroup) {
	router.POST("/weight", func(context *gin.Context) {
		response := trackingHandler.AddWeightRegister()
		context.JSON(response.Status.HttpStatus, response)
	})
}
