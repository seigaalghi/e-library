package decoder

import (
	"context"
	"net/http"
	"strconv"

	"github.com/seigaalghi/e-library/model"
)

func GetBooks(ctx context.Context, req *http.Request) (interface{}, error) {
	subject := req.URL.Query().Get("subject")
	limit := req.URL.Query().Get("limit")
	page := req.URL.Query().Get("page")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}

	if pageInt == 0 {
		pageInt = 1
	}

	if limitInt > 100 {
		limitInt = 100
	}

	if limitInt == 0 {
		limitInt = 10
	}

	return &model.GetBooksRequest{
		Subject: subject,
		Page:    pageInt,
		Limit:   limitInt,
	}, nil
}
