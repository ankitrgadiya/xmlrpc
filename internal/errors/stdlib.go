package errors

import (
	stderrors "errors"
)

// Import functions from Standard Library.
// nolint:gochecknoglobals
var (
	Is     = stderrors.Is
	As     = stderrors.As
	New    = stderrors.New
	Unwrap = stderrors.Unwrap
)
