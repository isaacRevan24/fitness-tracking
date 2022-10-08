package handler

import (
	"net/http"

	"github.com/isaacRevan24/fitness-tracking/model"
)

func (handler goalsHandler) AddGoalsRegisters(request model.AddWeightGoalsReq) model.FitnessResponse {
	var response model.FitnessResponse
	register, error := GoalsRepo.AddGoalsRegister(request)
	if error != nil {
		responseStatus := mapper.ToBaseStatus(http.StatusBadRequest, model.BAD_REQUEST_ERROR_STATUS, model.ADD_GOALS_REGISTER_ERROR)
		mapper.ToFitnessResponse(&response, responseStatus, nil)
		return response
	}
	status := mapper.ToBaseStatus(http.StatusOK, model.SUCCESS_CODE_STATUS, model.SUCCESS_MESSAGE)
	mapper.ToFitnessResponse(&response, status, register)
	return response
}
