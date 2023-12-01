package decoder

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/seigaalghi/e-library/common"
	"github.com/seigaalghi/e-library/model"
	"github.com/seigaalghi/e-library/pkg/zaplog"
	"go.uber.org/zap"
)

func LendBook(ctx context.Context, req *http.Request) (interface{}, error) {
	logger := zaplog.WithContext(ctx)
	defer logger.Sync()

	defer req.Body.Close()
	var mod model.LendBookRequest
	var err error
	body, _ := io.ReadAll(req.Body)
	if err := json.Unmarshal(body, &mod); err != nil {
		logger.Error("failed parsing body", zap.Error(err))
		return nil, common.NewErrBadRequestError(err.Error(), nil)
	}

	mod.DropOffDate, err = time.Parse("2006-01-02", mod.DropOffDateStr)
	if err != nil {
		return nil, common.NewErrBadRequestError(err.Error(), nil)
	}
	mod.PickupDate, err = time.Parse("2006-01-02", mod.PickupDateStr)
	if err != nil {
		return nil, common.NewErrBadRequestError(err.Error(), nil)
	}

	v := validator.New()
	err = v.Struct(&mod)
	if err != nil {
		logger.Error("failed validating body", zap.Error(err))
		return nil, common.NewErrBadRequestError(err.Error(), nil)
	}

	return &mod, nil
}
