package errors

// ErrorType enumerates the known error types for the application
type ErrorType int

const (
	DatabaseError ErrorType = iota + 1
	ValidationError
	InternalError
)

// AppError represents an error with an associated type and message
type AppError struct {
	Type    ErrorType
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

// New crea un nuevo AppError
func New(typ ErrorType, msg string) error {
	return &AppError{
		Type:    typ,
		Message: msg,
	}
}
