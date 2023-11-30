package encoder

import (
	"context"
	"encoding/json"
	"net/http"
)

type Error struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) CodeErr() int {
	return e.Code
}

func (e *Error) DataErr() interface{} {
	return e.Data
}

func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	errorCode := 400
	parsedErr, ok := err.(*Error)
	if ok {
		if parsedErr.Code/100 == http.StatusUnauthorized {
			w.Header().Set("Connection", `close`)
			w.Header().Set("X-Content-Type-Options", `nosniff`)
			w.Header().Set("Www-Authenticate", `Basic realm="Authorization Required"`)
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Not Authorized"))
			return
		}
		errorCode = parsedErr.Code / 100
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(errorCode)
	json.NewEncoder(w).Encode(err)
}
