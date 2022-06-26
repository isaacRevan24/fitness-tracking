package handler

import (
	"net/http"

	"github.com/isaacRevan24/fitness-tracking/model"
)

type trackingHandler struct{}

type TrackingHandlerInterface interface {
	AddWeightRegister() model.FitnessStatusResponse
}

func NewTrackingHandler() TrackingHandlerInterface {
	return &trackingHandler{}
}

func (*trackingHandler) AddWeightRegister() model.FitnessStatusResponse {
	status := model.FitnessStatusResponse{HttpStatus: http.StatusOK, Code: "fit-000", Message: "ok"}
	return status
}
