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
	var response model.FitnessResponse
	responseStatus := mapper.ToBaseStatus(http.StatusOK, "fit-000", "all ok")
	responseBody := model.AddWeightRegisterRes{WeightTrackId: uuid.UUID{}}
	mapper.ToFitnessResponse(&response, &responseStatus, responseBody)
	return response
}
