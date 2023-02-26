package errors

import e "errors"

var (
	ErrValidationFailed   = e.New("validation failed")
	ErrGetUUIDFromContext = e.New("invalid uuid from context")
	ErrBodyParse          = e.New("invalid body")
)
