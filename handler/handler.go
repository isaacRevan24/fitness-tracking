package handler

import (
	"net/http"

	"github.com/isaacRevan24/fitness-tracking/model"
)

type trackingHandler struct{}

type TrackingHandlerInterface interface {
	AddWeightRegister() model.FitnessStatus
}

func NewTrackingHandler() TrackingHandlerInterface {
	return &trackingHandler{}
}

func (*trackingHandler) AddWeightRegister() model.FitnessStatus {
	status := model.FitnessStatusResponse{HttpStatus: http.StatusOK, Code: "fit-000", Message: "ok"}
	statusResponse := model.FitnessStatus{Status: status}
	return statusResponse
}
