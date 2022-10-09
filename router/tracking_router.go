package router

import (
	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/fitness-tracking/handler"
	"github.com/isaacRevan24/fitness-tracking/model"
)

func (*routerRegister) TrackingRouter(router *gin.RouterGroup) {
	router.POST("/weight", func(context *gin.Context) {
		var request model.FitnessRequest[model.AddWeightRegisterReq]
		parsingError := mapper.GenericRequestJsonMapper(&request, context)
		if parsingError != nil {
			parsingErrorResponse(context)
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
			parsingErrorResponse(context)
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
			parsingErrorResponse(context)
			return
		}
		trackingHandler := handler.NewTrackingHandler()
		response := trackingHandler.DeleteWeightRegister(request.Body)
		context.JSON(response.Status.HttpStatus, response)
	})
}
