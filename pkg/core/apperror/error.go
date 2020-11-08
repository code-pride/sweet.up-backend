package apperror

var (
	ErrEntityNotFound = NewBaseError("Entity not found")
)

type ExternalServiceError struct {
	serviceError error
	serviceName  string
}

func NewExternalServiceError(serviceError error, serviceName string) *ExternalServiceError {
	return &ExternalServiceError{
		serviceError: serviceError,
		serviceName:  serviceName,
	}
}

func (err *ExternalServiceError) Error() string {
	return "External service " + err.serviceName + " request request failed with error: " + err.serviceError.Error()
}

type DatabaseError struct {
	driverError error
}

func NewDatabaseError(driverError error) *DatabaseError {
	return &DatabaseError{
		driverError: driverError,
	}
}

func (err *DatabaseError) Error() string {
	return "Database request failed with error: " + err.driverError.Error()
}

type BaseError struct {
	message string
}

func NewBaseError(message string) *BaseError {
	return &BaseError{
		message: message,
	}
}

func (err *BaseError) Error() string {
	return err.message
}

type InternalServerError struct {
	BaseError
}

func NewInternalServerError(message string) *InternalServerError {
	return &InternalServerError{
		*NewBaseError(message),
	}
}

type UserReqError struct {
	BaseError
}

func NewUserReqError(message string) *UserReqError {
	return &UserReqError{
		*NewBaseError(message),
	}
}
