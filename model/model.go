package model

const (
	SUCCESS_CODE_STATUS          = "FIT-00"
	INTERNAL_SERVER_ERROR_STATUS = "FIT-01"
	BAD_REQUEST_ERROR_STATUS     = "FIT-02"
)

const (
	SUCCESS_MESSAGE              = "Success."
	DATA_BASE_ERROR              = "Unable to connect to database."
	ADD_WEIGHT_REGISTER_ERROR    = "Unable to save weight register."
	GET_WEIGHT_REGISTER_ERROR    = "Unable to get weight register."
	INVALID_REQUEST              = "Invalid request parameters."
	UPDATE_WEIGHT_REGISTER_ERROR = "Unable to update weight register."
	DELETE_WEIGHT_REGISTER_ERROR = "Unable to delete weight register."
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
	Weight    float32 `json:"weight" binding:"required"`
	CreatedAt string  `json:"createdAt" binding:"required"`
	ClientId  string  `json:"clientId" binding:"required"`
}

type AddWeightRegisterRes struct {
	WeightTrackId string `json:"weightTrackId"`
}

type GetWeightRegisterRes struct {
	Weight    float32 `json:"weight"`
	CreatedAt string  `json:"createdAt"`
}

type UpdateWeightRegisterReq struct {
	ClientId      string  `json:"clientId" binding:"required"`
	WeightTrackId string  `json:"weightTrackId" binding:"required"`
	Weight        float32 `json:"weight" binding:"required"`
}

type DeleteWeightRegisterReq struct {
	WeightTrackId string `json:"weightTrackId" binding:"required"`
	ClientId      string `json:"clientId" binding:"required"`
}
