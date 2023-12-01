package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/seigaalghi/e-library/model"
	"github.com/seigaalghi/e-library/service"
)

type endpoints struct {
	GetBooks endpoint.Endpoint
	Login    endpoint.Endpoint
	LendBook endpoint.Endpoint
}

func MakeEndpoints(svc service.Service) endpoints {
	return endpoints{
		GetBooks: MakeGetBooks(svc),
		Login:    MakeLogin(svc),
		LendBook: MakeLendBook(svc),
	}
}

func MakeGetBooks(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*model.GetBooksRequest)
		return svc.GetBooks(ctx, req)
	}
}

func MakeLogin(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*model.LoginRequest)
		return svc.Login(ctx, req)
	}
}

func MakeLendBook(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*model.LendBookRequest)
		return svc.LendBook(ctx, req)
	}
}
