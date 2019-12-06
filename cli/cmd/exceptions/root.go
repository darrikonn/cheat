package exceptions

// AbortType : exception to abort cli execution
type AbortType struct {
	message string
}

func (e *AbortType) Error() string { return e.message }

// Abort : shorthand for AbortType for cleaner syntax
func Abort(message string) error {
	return &AbortType{message: message}
}

// CheatExceptionType : custom exception with shorthand
// message and original error
type CheatExceptionType struct {
	original error
	message  string
}

func (e *CheatExceptionType) Error() string { return e.message }

func (e *CheatExceptionType) Original() error { return e.original }

// CheatException : shorthand for CheatExceptionType for cleaner syntax
func CheatException(message string, original error) error {
	return &CheatExceptionType{message: message, original: original}
}
