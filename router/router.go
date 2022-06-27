package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/fitness-tracking/handler"
	"github.com/isaacRevan24/fitness-tracking/model"
)

var (
	trackingHandler handler.TrackingHandlerInterface = handler.NewTrackingHandler()
	mapper          model.FitnessMapper              = model.NewFitnessMapper()
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
		var request model.FitnessRequest[model.AddWeightRegisterReq]
		parsingError := mapper.GenericRequestJsonMapper(&request, context)
		fmt.Println(request)
		if parsingError != nil {
			var response model.FitnessStatusResponse
			mapper.ToStatusResponse(&response, http.StatusBadRequest, "FIT-001", "Invalid request.")
			context.JSON(response.Status.HttpStatus, response)
			return
		}
		response := trackingHandler.AddWeightRegister()
		context.JSON(response.Status.HttpStatus, response)
	})
}
