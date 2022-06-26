package model

type mapper struct{}

type FitnessMapper interface {
	ToSTatus(response *FitnessStatus, status int, code string, message string)
}

func NewFitnessMapper() FitnessMapper {
	return &mapper{}
}

func (*mapper) ToSTatus(response *FitnessStatus, status int, code string, message string) {
	statusValue := FitnessStatusResponse{HttpStatus: status, Code: code, Message: message}
	statusResponse := FitnessStatus{Status: statusValue}
	*response = statusResponse
}
