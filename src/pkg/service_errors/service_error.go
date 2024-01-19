package service_errors

type ServiceError struct {
	EndUserMessage  string `json:"endUserMessage"`
	Err             error
	TecnicalMessage string `json:"tecnicalMessage"`
}

func (e ServiceError) Error() string {
	return e.EndUserMessage
}
