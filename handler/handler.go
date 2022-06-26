package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type trackingHandler struct{}

type TrackingHandlerInterface interface {
	AddWeightRegister(context *gin.Context)
}

func NewTrackingHandler() TrackingHandlerInterface {
	return &trackingHandler{}
}

func (*trackingHandler) AddWeightRegister(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{
		"status": "fit-000",
	})

}
