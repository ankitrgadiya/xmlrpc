package codec

import "errors"

var (
	ErrInvalidParam      = errors.New("parameter type is invalid")
	ErrInvalidParamCount = errors.New("number of parameters do not match")
	ErrInvalidStartToken = errors.New("start token is not supported")
)
