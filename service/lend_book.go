package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/seigaalghi/e-library/common"
	"github.com/seigaalghi/e-library/model"
	"github.com/seigaalghi/e-library/pkg/zaplog"
	"go.uber.org/zap"
)

func (s *service) LendBook(ctx context.Context, req *model.LendBookRequest) (*model.LendBookResponse, error) {
	logger := zaplog.WithContext(ctx)
	defer logger.Sync()

	res := &model.LendBookResponse{
		EditionNumber: req.EditionNumber,
		DropOffDate:   req.DropOffDate,
		PickupDate:    req.PickupDate,
		Valid:         false,
	}

	book, err := s.getBookByEditionNumberFromOpenLib(ctx, req.EditionNumber)
	if err != nil {
		logger.Info("failed getting book by edition number", zap.Error(err))
		return nil, common.NewErrBadRequestError(err.Error(), res)
	}

	valid, err := s.repo.CreateSchedule(req.PickupDate, req.DropOffDate, &model.Books{
		Title:         book.Title,
		Authors:       book.Authors,
		EditionNumber: req.EditionNumber,
	})

	if err != nil {
		logger.Info("failed creating schedule", zap.Error(err))
		return nil, common.NewErrBadRequestError(err.Error(), res)
	}

	res.Valid = valid
	res.Title = book.Title

	if !res.Valid {
		res.Message = "the book already booked for the requested date"
	}

	return res, nil
}

func (s *service) getBookByEditionNumberFromOpenLib(ctx context.Context, editionNumber string) (*model.Book, error) {
	logger := zaplog.WithContext(ctx)
	defer logger.Sync()

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://openlibrary.org/works/%s.json", editionNumber), nil)
	if err != nil {
		return nil, err
	}

	var res model.Book
	code, body, err := s.request.CallWithLog(ctx, req, &res)
	if err != nil || code != 200 {
		logger.Info("failed calling get book by subject", zap.Error(err), zap.Int("code", code), zap.String("body", body))
		return nil, errors.New(res.Error)
	}

	return &res, nil
}
