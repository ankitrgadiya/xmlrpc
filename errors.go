package xmlrpc

import (
	"argc.in/xmlrpc/internal/codec"
	"argc.in/xmlrpc/internal/types"
)

var (
	ErrInvalidResponse           = types.ErrInvalidResponse
	ErrResponseMoreThanOneParam  = types.ErrResponseMoreThanOneParam
	ErrResponseInvalidStartToken = types.ErrResponseInvalidStartToken
	ErrInvalidParamCount         = codec.ErrInvalidParamCount
)
