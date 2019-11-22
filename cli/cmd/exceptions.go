package cmd

type AbortType struct {
	msg string // description of error
}

var Abort = &AbortType{}

func (e *AbortType) Error() string { return e.msg }
