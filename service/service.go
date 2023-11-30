package service

import (
	"github.com/seigaalghi/e-library/pkg/request"
	"github.com/seigaalghi/e-library/repository"
)

type service struct {
	repo    repository.Repository
	request request.HttpClient
}

func NewService(repo repository.Repository, request request.HttpClient) Service {
	return &service{
		repo:    repo,
		request: request,
	}
}
