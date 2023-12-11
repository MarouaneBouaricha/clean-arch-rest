package errors

// Used to return business error messages
type ServiceError struct {
	Message string `json:message`
}
