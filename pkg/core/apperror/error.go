package apperror

type Code int

const (
	QueriedEntityNotFound Code = iota
	UpdatedEntityNotFound
	ExternalServiceError
	InternalServerError
)

type ApplicationError struct {
	baseError error
	message   string
}

type EntityNotFoundError struct {
	ApplicationError
}

func NewApplicationError(baseError error, message string) ApplicationError {
	return ApplicationError{
		baseError: baseError,
		message:   message,
	}
}

func NewEntityNotFoundError(baseError error, message string) EntityNotFoundError {
	return EntityNotFoundError{NewApplicationError(baseError, message)}
}

func (appError ApplicationError) Error() string {
	return appError.message
}
