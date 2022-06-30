//go:generate go run github.com/golang/mock/mockgen -source router.go -destination mock/router_mock.go -package mock
package router

import (
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
		if parsingError != nil {
			var response model.FitnessStatusResponse
			mapper.ToStatusResponse(&response, http.StatusBadRequest, model.BAD_REQUEST_ERROR_STATUS, model.INVALID_REQUEST)
			context.JSON(response.Status.HttpStatus, response)
			return
		}
		response := trackingHandler.AddWeightRegister(request.Body)
		context.JSON(response.Status.HttpStatus, response)
	})
}
