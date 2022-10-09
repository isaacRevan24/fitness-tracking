package handler

import (
	"net/http"

	"github.com/isaacRevan24/fitness-tracking/model"
)

func (handler goalsHandler) AddGoalsRegister(request model.AddWeightGoalsReq) model.FitnessResponse {
	var response model.FitnessResponse
	register, error := GoalsRepo.AddGoalsRegister(request)
	if error != nil {
		responseStatus := mapper.ToBaseStatus(http.StatusBadRequest, model.BAD_REQUEST_ERROR_STATUS, model.ADD_GOALS_REGISTER_ERROR)
		mapper.ToFitnessResponse(&response, responseStatus, nil)
		return response
	}
	mapper.ToFitnessResponse(&response, buildSuccessMessageStatus(), register)
	return response
}

func (handler goalsHandler) GetGoalsRegister(request model.GetGoalsReq) model.FitnessResponse {
	var response model.FitnessResponse
	register, error := GoalsRepo.GetGoalsRegister(request.ClientId)
	if error != nil {
		responseStatus := mapper.ToBaseStatus(http.StatusBadRequest, model.BAD_REQUEST_ERROR_STATUS, model.GET_GOALS_REGISTER_ERROR)
		mapper.ToFitnessResponse(&response, responseStatus, nil)
		return response
	}
	mapper.ToFitnessResponse(&response, buildSuccessMessageStatus(), register)
	return response
}
