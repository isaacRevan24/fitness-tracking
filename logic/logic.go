//go:generate go run github.com/golang/mock/mockgen -source logic.go -destination mock/logic_mock.go -package mock
package logic

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/isaacRevan24/fitness-tracking/model"
	"github.com/isaacRevan24/fitness-tracking/repository"
)

var (
	trackingRepository repository.TrackingRepository
)

type trackingLogic struct{}

type TrackingLogicInterface interface {
	AddWeightRegister(model.AddWeightRegisterReq) (uuid.UUID, error)
}

func NewTrackingLogic(repo repository.TrackingRepository) TrackingLogicInterface {
	trackingRepository = repo
	return &trackingLogic{}
}

func (*trackingLogic) AddWeightRegister(request model.AddWeightRegisterReq) (uuid.UUID, error) {
	response, error := trackingRepository.AddWeightRegister(request)
	if error != nil {
		fmt.Println(error)
		return uuid.Nil, error
	}
	return response, nil
}
