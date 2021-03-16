package types

import "errors"

var (
	ErrInvalidResponse           = errors.New("invalid method response")
	ErrResponseMoreThanOneParam  = errors.New("more than one parameter in methodResponse")
	ErrResponseInvalidStartToken = errors.New("invalid start token in response")
)
