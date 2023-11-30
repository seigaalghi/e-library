package service

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net/http"
	"net/url"

	"github.com/seigaalghi/e-library/common"
	"github.com/seigaalghi/e-library/model"
	"github.com/seigaalghi/e-library/pkg/zaplog"
	"go.uber.org/zap"
)

func (s *service) GetBooks(ctx context.Context, req *model.GetBooksRequest) (*model.GetBooksResponse, error) {
	logger := zaplog.WithContext(ctx)
	defer logger.Sync()

	offset := req.Limit * (req.Page - 1)
	books, err := s.getBookFromOpenLib(ctx, req.Subject, offset, req.Limit)
	if err != nil {
		logger.Info("failed getting book from open lib", zap.Error(err))
		return nil, common.NewErrBadRequestError(err.Error(), []map[string]interface{}{})
	}

	mappedBook := make([]*model.Books, 0)
	for _, book := range books.Works {
		logger.Info("asdasd", zap.Any("asd", mappedBook))
		mappedBook = append(mappedBook, &model.Books{
			Title:         book.Title,
			Authors:       book.Authors,
			EditionNumber: book.CoverEditionKey,
		})
	}

	return &model.GetBooksResponse{
		Books:       mappedBook,
		TotalPages:  int(math.Ceil(float64(books.WorkCount) / float64(req.Limit))),
		CurrentPage: req.Page,
		Limit:       req.Limit,
	}, nil
}

func (s *service) getBookFromOpenLib(ctx context.Context, subject string, offset, limit int) (*model.GetBookBySubjectResponse, error) {
	logger := zaplog.WithContext(ctx)
	defer logger.Sync()

	requestURL, err := url.Parse(fmt.Sprintf("%s/subjects/%s.json", common.OpenLibraryBaseUrl, subject))
	if err != nil {
		logger.Info("failed creating url", zap.Error(err))
		return nil, err
	}

	queryParams := requestURL.Query()
	queryParams.Set("limit", fmt.Sprint(limit))
	queryParams.Set("offset", fmt.Sprint(offset))
	requestURL.RawQuery = queryParams.Encode()

	req, err := http.NewRequest(http.MethodGet, requestURL.String(), nil)
	if err != nil {
		logger.Info("failed initializing request", zap.Error(err))
	}

	var res model.GetBookBySubjectResponse
	code, body, err := s.request.CallWithLog(ctx, req, &res)
	if err != nil || code != 200 {
		logger.Info("failed calling get book by subject", zap.Error(err), zap.Int("code", code), zap.String("body", body))
		return nil, errors.New(res.Error)
	}

	return &res, nil
}
