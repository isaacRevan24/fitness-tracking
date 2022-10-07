//go:generate go run github.com/golang/mock/mockgen -source handler.go -destination mock/handler_mock.go -package mock
package handler

import (
	"net/http"

	"github.com/isaacRevan24/fitness-tracking/model"
	"github.com/isaacRevan24/fitness-tracking/repository"
)

var (
	Repo      repository.TrackingRepository = repository.NewTrackingRepository()
	GoalsRepo repository.GoalsRepository    = repository.NewGoalsRepository()
	mapper    model.FitnessMapper           = model.NewFitnessMapper()
)

type trackingHandler struct{}

type goalsHandler struct{}

type TrackingHandlerInterface interface {
	AddWeightRegister(request model.AddWeightRegisterReq) model.FitnessResponse
	GetWeightRegister(clientId string, weightId string) model.FitnessResponse
	UpdateWeightRegister(request model.UpdateWeightRegisterReq) model.FitnessResponse
	DeleteWeightRegister(request model.DeleteWeightRegisterReq) model.FitnessResponse
}

type GoalsHandlerInterface interface {
	AddGoalsRegisters(request model.AddWeightGoalsReq) model.FitnessResponse
}

func NewTrackingHandler() TrackingHandlerInterface {
	return trackingHandler{}
}

func NewGoalsHandler() GoalsHandlerInterface {
	return goalsHandler{}
}

func (handler trackingHandler) AddWeightRegister(request model.AddWeightRegisterReq) model.FitnessResponse {
	var response model.FitnessResponse
	id, error := Repo.AddWeightRegister(request)
	if error != nil {
		responseStatus := mapper.ToBaseStatus(http.StatusBadRequest, model.BAD_REQUEST_ERROR_STATUS, model.ADD_WEIGHT_REGISTER_ERROR)
		mapper.ToFitnessResponse(&response, responseStatus, nil)
		return response
	}
	responseStatus := mapper.ToBaseStatus(http.StatusOK, model.SUCCESS_CODE_STATUS, model.SUCCESS_MESSAGE)
	responseBody := model.AddWeightRegisterRes{WeightTrackId: id}
	mapper.ToFitnessResponse(&response, responseStatus, responseBody)
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
	responseStatus := mapper.ToBaseStatus(http.StatusOK, model.SUCCESS_CODE_STATUS, model.SUCCESS_MESSAGE)
	mapper.ToFitnessResponse(&response, responseStatus, register)
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
	responseStatus := mapper.ToBaseStatus(http.StatusOK, model.SUCCESS_CODE_STATUS, model.SUCCESS_MESSAGE)
	mapper.ToFitnessResponse(&response, responseStatus, nil)
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
	responseStatus := mapper.ToBaseStatus(http.StatusOK, model.SUCCESS_CODE_STATUS, model.SUCCESS_MESSAGE)
	mapper.ToFitnessResponse(&response, responseStatus, nil)
	return response
}

func (handler goalsHandler) AddGoalsRegisters(request model.AddWeightGoalsReq) model.FitnessResponse {
	var response model.FitnessResponse
	register, error := GoalsRepo.AddGoalsRegisters(request)
	if error != nil {
		responseStatus := mapper.ToBaseStatus(http.StatusBadRequest, model.BAD_REQUEST_ERROR_STATUS, model.ADD_GOALS_REGISTER_ERROR)
		mapper.ToFitnessResponse(&response, responseStatus, nil)
		return response
	}
	status := mapper.ToBaseStatus(http.StatusOK, model.SUCCESS_CODE_STATUS, model.SUCCESS_MESSAGE)
	mapper.ToFitnessResponse(&response, status, register)
	return response
}
