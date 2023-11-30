package service

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/seigaalghi/e-library/common"
	"github.com/seigaalghi/e-library/model"
	"github.com/seigaalghi/e-library/pkg/zaplog"
	"go.uber.org/zap"
)

func (s *service) Login(ctx context.Context, req *model.LoginRequest) (*model.LoginResponse, error) {
	logger := zaplog.WithContext(ctx)
	defer logger.Sync()

	exp := time.Now().Add(time.Minute * 15).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": req.Email,
		"exp":   exp,
		"iat":   time.Now().Unix(),
	})

	accessToken, err := token.SignedString([]byte(common.JwtKey))
	if err != nil {
		logger.Info("failed signing token", zap.Error(err))
		return nil, common.ErrInternal
	}

	return &model.LoginResponse{
		Token:     accessToken,
		ExpiredAt: exp,
	}, nil
}
