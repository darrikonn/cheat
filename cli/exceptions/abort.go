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
