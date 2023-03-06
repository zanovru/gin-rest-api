package resources

type ErrorResource struct {
	Message string `json:"message"`
	Errors  any    `json:"errors"`
}

func NewErrorResource(message string, errors ...any) *ErrorResource {
	return &ErrorResource{Message: message, Errors: errors}
}
