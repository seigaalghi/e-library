package service

import (
	"context"

	"github.com/seigaalghi/e-library/model"
)

type Service interface {
	Login(ctx context.Context, req *model.LoginRequest) (*model.LoginResponse, error)
	GetBooks(ctx context.Context, req *model.GetBooksRequest) (*model.GetBooksResponse, error)
	LendBook(ctx context.Context, req *model.LendBookRequest) (*model.LendBookResponse, error)
}
