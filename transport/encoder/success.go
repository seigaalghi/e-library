package encoder

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/seigaalghi/e-library/model"
)

func EncodeResponseWithData(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("content-type", "application/json")
	resp := model.BasicResponse{Data: response, Success: true, Message: ""}
	return json.NewEncoder(w).Encode(resp)
}
