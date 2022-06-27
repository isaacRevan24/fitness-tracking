package model

import (
	"github.com/google/uuid"
)

type BaseStatus struct {
	HttpStatus int    `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}

type FitnessStatusResponse struct {
	Status BaseStatus `json:"status"`
}

type FitnessRequest[T any] struct {
	Body T `json:"body" binding:"required"`
}

type FitnessResponse struct {
	Status BaseStatus `json:"status"`
	T      any        `json:"body,omitempty"`
}

type AddWeightRegisterReq struct {
	Weight    float32   `json:"weight" binding:"required"`
	CreatedAt string    `json:"createdAt" binding:"required"`
	ClientId  uuid.UUID `json:"clientId" binding:"required"`
}

type AddWeightRegisterRes struct {
	WeightTrackId uuid.UUID `json:"weightTrackId"`
}
