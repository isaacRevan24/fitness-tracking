//go:generate go run github.com/golang/mock/mockgen -source router.go -destination mock/router_mock.go -package mock
package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/fitness-tracking/handler"
	"github.com/isaacRevan24/fitness-tracking/model"
)

var (
	mapper model.FitnessMapper = model.NewFitnessMapper()
)

type routerRegister struct{}

type RouterRegisterInterface interface {
	TrackingRouter(router *gin.RouterGroup)
	GoalsRouter(router *gin.RouterGroup)
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
		trackingHandler := handler.NewTrackingHandler()
		response := trackingHandler.AddWeightRegister(request.Body)
		context.JSON(response.Status.HttpStatus, response)
	})

	router.GET("/weight", func(context *gin.Context) {
		clientId := context.Query("clientId")
		weightId := context.Query("weightId")
		trackingHandler := handler.NewTrackingHandler()
		response := trackingHandler.GetWeightRegister(clientId, weightId)
		context.JSON(response.Status.HttpStatus, response)
	})

	router.PATCH("/weight", func(context *gin.Context) {
		var request model.FitnessRequest[model.UpdateWeightRegisterReq]
		parsingError := mapper.GenericRequestJsonMapper(&request, context)
		if parsingError != nil {
			var response model.FitnessStatusResponse
			mapper.ToStatusResponse(&response, http.StatusBadRequest, model.BAD_REQUEST_ERROR_STATUS, model.INVALID_REQUEST)
			context.JSON(response.Status.HttpStatus, response)
			return
		}
		trackingHandler := handler.NewTrackingHandler()
		response := trackingHandler.UpdateWeightRegister(request.Body)
		context.JSON(response.Status.HttpStatus, response)
	})

	router.DELETE("/weight", func(context *gin.Context) {
		var request model.FitnessRequest[model.DeleteWeightRegisterReq]
		parsingError := mapper.GenericRequestJsonMapper(&request, context)
		if parsingError != nil {
			var response model.FitnessStatusResponse
			mapper.ToStatusResponse(&response, http.StatusBadRequest, model.BAD_REQUEST_ERROR_STATUS, model.INVALID_REQUEST)
			context.JSON(response.Status.HttpStatus, response)
			return
		}
		trackingHandler := handler.NewTrackingHandler()
		response := trackingHandler.DeleteWeightRegister(request.Body)
		context.JSON(response.Status.HttpStatus, response)
	})
}

func (*routerRegister) GoalsRouter(router *gin.RouterGroup) {
	router.POST("/", func(context *gin.Context) {
		var request model.FitnessRequest[model.AddWeightGoalsReq]
		parsingError := mapper.GenericRequestJsonMapper(&request, context)
		if parsingError != nil {
			var response model.FitnessStatusResponse
			mapper.ToStatusResponse(&response, http.StatusBadRequest, model.BAD_REQUEST_ERROR_STATUS, model.INVALID_REQUEST)
			context.JSON(response.Status.HttpStatus, response)
			return
		}
		goalsHandler := handler.NewGoalsHandler()
		response := goalsHandler.AddGoalsRegisters(request.Body)
		context.JSON(response.Status.HttpStatus, response)
	})
}
