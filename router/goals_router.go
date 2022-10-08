package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaacRevan24/fitness-tracking/handler"
	"github.com/isaacRevan24/fitness-tracking/model"
)

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
		response := goalsHandler.AddGoalsRegister(request.Body)
		context.JSON(response.Status.HttpStatus, response)
	})
}
