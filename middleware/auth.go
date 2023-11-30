package middleware

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport/http"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/seigaalghi/e-library/common"
	"github.com/seigaalghi/e-library/pkg/zaplog"
	"go.uber.org/zap"
)

func BaseMiddleware(requestType string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		next = JwtAuth(requestType)(next)
		next = TransportLogging(requestType)(next)

		return next
	}
}

func JwtAuth(requestType string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
			logger := zaplog.WithContext(ctx)
			defer logger.Sync()

			auth, ok := ctx.Value(http.ContextKeyRequestAuthorization).(string)
			if !ok {
				return nil, errors.New("error")
			}
			authToken := strings.ReplaceAll(auth, "Bearer ", "")

			if authToken == "" {
				return nil, common.ErrUnauthorized
			}

			token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				return []byte(common.JwtKey), nil
			})

			if err != nil {
				logger.Info("failed authorizing token", zap.Error(err))
				return nil, common.ErrUnauthorized
			}

			if !token.Valid {
				return nil, common.ErrUnauthorized
			}

			return next(ctx, request)
		}
	}
}

func TransportLogging(requestType string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
			requestID, ok := ctx.Value(http.ContextKeyRequestXRequestID).(string)
			if !ok || requestID == "" {
				requestID = uuid.NewString()
			}

			ctx = zaplog.NewContext(ctx, zap.String("request_type", requestType), zap.String("request_id", requestID))
			logger := zaplog.WithContext(ctx)

			defer func(requestTime time.Time, logger *zap.Logger) {
				logger.Info("HTTP Request - "+requestType,
					zap.Any("request", request),
					zap.Any("response", resp),
					zap.Error(err),
					zap.String("execution_time", time.Since(requestTime).String()),
				)
				defer logger.Sync()
			}(time.Now(), logger)

			return next(ctx, request)
		}
	}
}
