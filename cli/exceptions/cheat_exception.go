package exceptions

// CheatExceptionType : custom exception with shorthand
// message and original error
type CheatExceptionType struct {
	original error
	message  string
}

// Error : Returns the error message from CheatExceptionType
func (e *CheatExceptionType) Error() string { return e.message }

// Original : Returns the original error from CheatExceptionType
func (e *CheatExceptionType) Original() error { return e.original }

// CheatException : shorthand for CheatExceptionType for cleaner syntax
func CheatException(message string, original error) error {
	return &CheatExceptionType{message: message, original: original}
}
