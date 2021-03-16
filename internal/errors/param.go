package errors

import "fmt"

var (
	ErrNotAPointer = NewInvalidParamError("not a pointer")
)

type ErrInvalidParam struct {
	msg string
}

func NewInvalidParamError(msg string) error { return &ErrInvalidParam{msg: msg} }

func (e *ErrInvalidParam) Error() string {
	return fmt.Sprintf("invalid parameter: %s", e.msg)
}

func (e *ErrInvalidParam) Is(target error) bool {
	_, ok := target.(*ErrInvalidParam)
	return ok
}
