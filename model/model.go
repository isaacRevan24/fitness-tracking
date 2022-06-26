package model

type FitnessStatusResponse struct {
	HttpStatus int    `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}

type FitnessStatus struct {
	Status FitnessStatusResponse `json:"status"`
}
