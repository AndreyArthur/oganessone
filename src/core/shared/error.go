package shared

type Error struct {
	Type    string
	Name    string
	Message string
}

func NewError(errorType string, name string, message string) *Error {
	return &Error{
		Type:    errorType,
		Name:    name,
		Message: message,
	}
}
