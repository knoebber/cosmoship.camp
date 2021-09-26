package usererror

// Error returns a message to the user.
type Error struct {
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

// New returns a new usererror.
func New(message string) *Error {
	return &Error{Message: message}
}
