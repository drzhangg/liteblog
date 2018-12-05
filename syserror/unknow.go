package syserror

type UnKnowError struct {
	msg string
	reason error
}

func New(msg string, reason error) Error {
	return UnKnowError{msg,reason}
}
