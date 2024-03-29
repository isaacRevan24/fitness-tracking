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
	AddGoalsRegister(request model.AddWeightGoalsReq) model.FitnessResponse
	GetGoalsRegister(clientId string) model.FitnessResponse
	UpdateWeightGoal(request model.UpdateWeightGoalReq) model.FitnessResponse
}

func NewTrackingHandler() TrackingHandlerInterface {
	return trackingHandler{}
}

func NewGoalsHandler() GoalsHandlerInterface {
	return goalsHandler{}
}

func buildSuccessMessageStatus() model.BaseStatus {
	return mapper.ToBaseStatus(http.StatusOK, model.SUCCESS_CODE_STATUS, model.SUCCESS_MESSAGE)
}
