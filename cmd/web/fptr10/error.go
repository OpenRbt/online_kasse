package fptr10

type Error struct {
	ErrorCode        int
	ErrorDescription string
}

func (e *Error) Error() string {
	return e.ErrorDescription
}
