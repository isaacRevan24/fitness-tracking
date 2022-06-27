package model

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type mapper struct{}

type FitnessMapper interface {
	ToStatusResponse(response *FitnessStatusResponse, status int, code string, message string)
	ToFitnessResponse(response *FitnessResponse, status *BaseStatus, T any)
	ToBaseStatus(status int, code string, message string) BaseStatus
	GenericRequestJsonMapper(T any, context *gin.Context) error
}

func NewFitnessMapper() FitnessMapper {
	return &mapper{}
}

func (*mapper) ToBaseStatus(status int, code string, message string) BaseStatus {
	return BaseStatus{HttpStatus: status, Code: code, Message: message}
}

func (*mapper) ToStatusResponse(response *FitnessStatusResponse, status int, code string, message string) {
	statusValue := BaseStatus{HttpStatus: status, Code: code, Message: message}
	statusResponse := FitnessStatusResponse{Status: statusValue}
	*response = statusResponse
}

func (*mapper) ToFitnessResponse(response *FitnessResponse, status *BaseStatus, T any) {
	*response = FitnessResponse{Status: *status, T: T}
}

func (*mapper) GenericRequestJsonMapper(request interface{}, context *gin.Context) error {
	if err := context.ShouldBindJSON(request); err != nil {
		return errors.New("missing argument")
	}
	return nil
}
