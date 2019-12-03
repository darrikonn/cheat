package exceptions

// AbortType : exception to abort cli execution
type AbortType struct {
	msg string
}

// Abort : shorthand for AbortType for cleaner syntax
var Abort = &AbortType{}

func (e *AbortType) Error() string { return e.msg }
