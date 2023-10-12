package errors

type AppError struct {
	Status  uint
	Message string
}

func (this AppError) Error() string {
	return this.Message
}

func NewAppError(status uint, message string) *AppError {
	return &AppError{Status: status, Message: message}
}

// TODO: Define errors internal to the application (Objective is to decouple from HTTP Status Codes)
