package errors

import "fmt"

type ErrInvalidResponse struct {
	msg string
}

func NewInvalidResponseError(msg string) error { return &ErrInvalidResponse{msg: msg} }

func (e *ErrInvalidResponse) Error() string {
	return fmt.Sprintf("invalid response: %s", e.msg)
}

func (e *ErrInvalidResponse) Is(target error) bool {
	_, ok := target.(*ErrInvalidResponse)
	return ok
}
