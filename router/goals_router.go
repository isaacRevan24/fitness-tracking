package router

import (
	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/fitness-tracking/handler"
	"github.com/isaacRevan24/fitness-tracking/model"
)

func (*routerRegister) GoalsRouter(router *gin.RouterGroup) {
	router.POST("/", func(context *gin.Context) {
		var request model.FitnessRequest[model.AddWeightGoalsReq]
		parsingError := mapper.GenericRequestJsonMapper(&request, context)
		if parsingError != nil {
			parsingErrorResponse(context)
			return
		}
		goalsHandler := handler.NewGoalsHandler()
		response := goalsHandler.AddGoalsRegister(request.Body)
		context.JSON(response.Status.HttpStatus, response)
	})

	router.GET("/:clientId", func(context *gin.Context) {
		clientId := context.Param("clientId")
		goalsHandler := handler.NewGoalsHandler()
		response := goalsHandler.GetGoalsRegister(clientId)
		context.JSON(response.Status.HttpStatus, response)
	})
}
