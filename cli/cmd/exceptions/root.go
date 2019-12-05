package exceptions

// AbortType : exception to abort cli execution
type AbortType struct {
	msg string
}

// Abort : shorthand for AbortType for cleaner syntax
func Abort(msg string) error {
	return &AbortType{msg: msg}
}

func (e *AbortType) Error() string { return e.msg }
