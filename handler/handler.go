package handler

import (
	"net/http"

	"github.com/isaacRevan24/fitness-tracking/model"
)

var (
	mapper model.FitnessMapper = model.NewFitnessMapper()
)

type trackingHandler struct{}

type TrackingHandlerInterface interface {
	AddWeightRegister() model.FitnessStatus
}

func NewTrackingHandler() TrackingHandlerInterface {
	return &trackingHandler{}
}

func (*trackingHandler) AddWeightRegister() model.FitnessStatus {
	var response model.FitnessStatus
	mapper.ToSTatusResponse(&response, http.StatusOK, "fit-00", "ok")
	return response
}
