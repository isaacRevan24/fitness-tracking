package handler

import (
	"net/http"

	"github.com/isaacRevan24/fitness-tracking/logic"
	"github.com/isaacRevan24/fitness-tracking/model"
	"github.com/isaacRevan24/fitness-tracking/repository"
)

var (
	trackingLogic logic.TrackingLogicInterface
	mapper        model.FitnessMapper = model.NewFitnessMapper()
)

type trackingHandler struct{}

type TrackingHandlerInterface interface {
	AddWeightRegister(request model.AddWeightRegisterReq) model.FitnessResponse
}

func NewTrackingHandler() TrackingHandlerInterface {
	trackingRepository := repository.NewTrackingRepository()
	trackingLogic = logic.NewTrackingLogic(trackingRepository)
	return &trackingHandler{}
}

func (*trackingHandler) AddWeightRegister(request model.AddWeightRegisterReq) model.FitnessResponse {
	var response model.FitnessResponse
	id, error := trackingLogic.AddWeightRegister(request)
	if error != nil {
		responseStatus := mapper.ToBaseStatus(http.StatusBadRequest, "fit-001", "Error guardando registro.")
		mapper.ToFitnessResponse(&response, &responseStatus, nil)
		return response
	}
	responseStatus := mapper.ToBaseStatus(http.StatusOK, "fit-000", "all ok")
	responseBody := model.AddWeightRegisterRes{WeightTrackId: id}
	mapper.ToFitnessResponse(&response, &responseStatus, responseBody)
	return response
}
