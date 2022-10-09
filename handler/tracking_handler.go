package handler

import (
	"net/http"

	"github.com/isaacRevan24/fitness-tracking/model"
)

func (handler trackingHandler) AddWeightRegister(request model.AddWeightRegisterReq) model.FitnessResponse {
	var response model.FitnessResponse
	id, error := Repo.AddWeightRegister(request)
	if error != nil {
		responseStatus := mapper.ToBaseStatus(http.StatusBadRequest, model.BAD_REQUEST_ERROR_STATUS, model.ADD_WEIGHT_REGISTER_ERROR)
		mapper.ToFitnessResponse(&response, responseStatus, nil)
		return response
	}
	responseBody := model.AddWeightRegisterRes{WeightTrackId: id}
	mapper.ToFitnessResponse(&response, buildSuccessMessageStatus(), responseBody)
	return response
}

func (handler trackingHandler) GetWeightRegister(clientId string, weightId string) model.FitnessResponse {
	var response model.FitnessResponse
	register, error := Repo.GetWeightRegister(clientId, weightId)
	if error != nil {
		responseStatus := mapper.ToBaseStatus(http.StatusBadRequest, model.BAD_REQUEST_ERROR_STATUS, model.GET_WEIGHT_REGISTER_ERROR)
		mapper.ToFitnessResponse(&response, responseStatus, nil)
		return response
	}
	mapper.ToFitnessResponse(&response, buildSuccessMessageStatus(), register)
	return response
}

func (handler trackingHandler) UpdateWeightRegister(request model.UpdateWeightRegisterReq) model.FitnessResponse {
	var response model.FitnessResponse
	error := Repo.UpdateWeightRegister(request)
	if error != nil {
		responseStatus := mapper.ToBaseStatus(http.StatusBadRequest, model.BAD_REQUEST_ERROR_STATUS, model.UPDATE_WEIGHT_REGISTER_ERROR)
		mapper.ToFitnessResponse(&response, responseStatus, nil)
		return response
	}
	mapper.ToFitnessResponse(&response, buildSuccessMessageStatus(), nil)
	return response
}

func (handler trackingHandler) DeleteWeightRegister(request model.DeleteWeightRegisterReq) model.FitnessResponse {
	var response model.FitnessResponse
	error := Repo.DeleteWeightRegister(request)
	if error != nil {
		responseStatus := mapper.ToBaseStatus(http.StatusBadRequest, model.BAD_REQUEST_ERROR_STATUS, model.DELETE_WEIGHT_REGISTER_ERROR)
		mapper.ToFitnessResponse(&response, responseStatus, nil)
		return response
	}
	mapper.ToFitnessResponse(&response, buildSuccessMessageStatus(), nil)
	return response
}
