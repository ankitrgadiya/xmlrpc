package errors

import (
	"encoding/xml"
	"fmt"
)

type ErrInvalidStartToken struct {
	start *xml.StartElement
}

func NewInvalidStartTokenError(start *xml.StartElement) error {
	return &ErrInvalidStartToken{start: start}
}

func (e *ErrInvalidStartToken) Error() string {
	return fmt.Sprintf("invalid start token: %s", e.start.Name.Local)
}

func (e *ErrInvalidStartToken) Is(target error) bool {
	_, ok := target.(*ErrInvalidStartToken)
	return ok
}
