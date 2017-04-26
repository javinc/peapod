package peapod

// General errors.
const (
	ErrUnauthorized = Error("unauthorized")
)

// Error represents a peapod error.
type Error string

// Error returns the error as a string.
func (e Error) Error() string { return string(e) }
