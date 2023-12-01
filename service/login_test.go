package service

import (
	"context"
	"testing"

	"github.com/seigaalghi/e-library/model"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	type payload struct {
		ctx context.Context
		req *model.LoginRequest
	}

	tests := []struct {
		scenario string
		payload  payload
		assert   func(t *testing.T, res *model.LoginResponse, err error)
	}{
		{
			scenario: "success",
			payload: payload{
				ctx: context.Background(),
				req: &model.LoginRequest{
					Email:    "abc123@gmail.com",
					Password: "abcd1234",
				},
			},
			assert: func(t *testing.T, res *model.LoginResponse, err error) {
				assert.Nil(t, err)
				assert.NotEmpty(t, res.Token)
				assert.NotEmpty(t, res.ExpiredAt)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.scenario, func(t *testing.T) {
			svc, _, _ := getService(t)

			res, err := svc.Login(tt.payload.ctx, tt.payload.req)

			tt.assert(t, res, err)
		})
	}
}
