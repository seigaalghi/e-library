package common

import "github.com/seigaalghi/e-library/transport/encoder"

var (
	ErrUnauthorized = &encoder.Error{
		Code:    40100,
		Message: "Unauthorized",
	}

	NewErrBadRequestError = func(msg string, data interface{}) *encoder.Error {
		return &encoder.Error{
			Code:    40000,
			Message: msg,
			Data:    data,
		}
	}

	ErrInternal = &encoder.Error{
		Code:    50000,
		Message: "Internal Server Error",
	}
)
