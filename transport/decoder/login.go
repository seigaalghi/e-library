package decoder

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/seigaalghi/e-library/common"
	"github.com/seigaalghi/e-library/model"
	"github.com/seigaalghi/e-library/pkg/zaplog"
	"go.uber.org/zap"
)

func Login(ctx context.Context, r *http.Request) (interface{}, error) {
	logger := zaplog.WithContext(ctx)
	defer logger.Sync()

	defer r.Body.Close()
	var mod model.LoginRequest
	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &mod); err != nil {
		logger.Error("failed parsing body", zap.Error(err))
		return nil, common.NewErrBadRequestError(err.Error(), nil)
	}

	v := validator.New()
	err := v.Struct(&mod)
	if err != nil {
		logger.Error("failed validating body", zap.Error(err))
		return nil, common.NewErrBadRequestError(err.Error(), nil)
	}

	return &mod, nil
}
