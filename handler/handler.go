//go:generate go run github.com/golang/mock/mockgen -source handler.go -destination ../test/mock/handler_mock.go -package mock
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
		responseStatus := mapper.ToBaseStatus(http.StatusBadRequest, model.BAD_REQUEST_ERROR_STATUS, model.ADD_WEIGHT_REGISTER_ERROR)
		mapper.ToFitnessResponse(&response, &responseStatus, nil)
		return response
	}
	responseStatus := mapper.ToBaseStatus(http.StatusOK, model.SUCCESS_CODE_STATUS, model.SUCCESS_MESSAGE)
	responseBody := model.AddWeightRegisterRes{WeightTrackId: id}
	mapper.ToFitnessResponse(&response, &responseStatus, responseBody)
	return response
}
