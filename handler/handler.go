package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/isaacRevan24/fitness-tracking/model"
)

var (
	mapper model.FitnessMapper = model.NewFitnessMapper()
)

type trackingHandler struct{}

type TrackingHandlerInterface interface {
	AddWeightRegister() model.FitnessResponse
}

func NewTrackingHandler() TrackingHandlerInterface {
	return &trackingHandler{}
}

func (*trackingHandler) AddWeightRegister() model.FitnessResponse {
	var status model.FitnessStatusResponse
	responseBody := model.AddWeightRegisterRes{WeightTrackId: uuid.UUID{}}
	var response model.FitnessResponse
	mapper.ToStatusResponse(&status, http.StatusOK, "fit-00", "ok")
	mapper.ToFitnessResponse(&response, &status.Status, responseBody)
	return response
}
