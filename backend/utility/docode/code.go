package docode

type ErrorWithCode struct {
	code    int
	message string
}

func NewError(code int, message string) *ErrorWithCode {
	return &ErrorWithCode{
		code:    code,
		message: message,
	}
}

func (e *ErrorWithCode) Error() string {
	return e.message
}

func (e *ErrorWithCode) Code() int {
	return e.code
}
