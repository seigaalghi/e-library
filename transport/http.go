package transport

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/seigaalghi/e-library/endpoint"
	"github.com/seigaalghi/e-library/middleware"
	"github.com/seigaalghi/e-library/pkg/request"
	"github.com/seigaalghi/e-library/pkg/sql"
	"github.com/seigaalghi/e-library/pkg/zaplog"
	"github.com/seigaalghi/e-library/repository"
	"github.com/seigaalghi/e-library/service"
	"github.com/seigaalghi/e-library/transport/decoder"
	"github.com/seigaalghi/e-library/transport/encoder"
	"go.uber.org/zap"

	kithttp "github.com/go-kit/kit/transport/http"
)

func NewHttpHandler(r *mux.Router) http.Handler {
	logger := zaplog.WithContext(context.Background())
	defer logger.Sync()

	db, err := sql.NewSqlConnection(sql.DBConfiguration{
		DBFile: "./data.db",
		Driver: sql.SQLITE,
	})

	if err != nil {
		logger.Info("failed connecting to DB", zap.Error(err))
	}

	repo := repository.NewRepository(db)
	request := request.NewClient(http.DefaultClient)
	svc := service.NewService(repo, request)

	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encoder.EncodeError),
		kithttp.ServerBefore(kithttp.PopulateRequestContext),
	}

	edp := endpoint.MakeEndpoints(svc)

	getBooks := middleware.BaseMiddleware("get_books")(edp.GetBooks)
	getBooksHandler := kithttp.NewServer(
		getBooks,
		decoder.GetBooks,
		encoder.EncodeResponseWithData,
		opts...,
	)

	login := middleware.TransportLogging("login")(edp.Login)
	loginHandler := kithttp.NewServer(
		login,
		decoder.Login,
		encoder.EncodeResponseWithData,
		opts...,
	)

	r.Handle("/login", loginHandler).Methods("POST")
	r.Handle("/books", getBooksHandler).Methods("GET")

	return r
}
