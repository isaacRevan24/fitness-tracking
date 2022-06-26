package model

type FitnessStatusResponse struct {
	HttpStatus int    `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}

type FitnessStatus struct {
	Status FitnessStatusResponse `json:"status"`
}

type FitnessRequest struct {
	T any `json:"body"`
}

type FitnessResponse struct {
	Status FitnessStatusResponse `json:"status"`
	T      any                   `json:"body"`
}

type AddWeightRegister struct {
	Weight float32 `json:"weight" binding:"required"`
	//CreatedAt time.Time `json:"createdAt" binding:"required"`
	//ClientId  uuid.UUID `json:"clientId" binding:"required"`
}
