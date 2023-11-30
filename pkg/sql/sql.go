package sql

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/seigaalghi/e-library/pkg/zaplog"
	"go.uber.org/zap"
)

const (
	POSTGRES = "postgres"
	MYSQL    = "mysql"
	SQLITE   = "sqlite3"
)

type DBConfiguration struct {
	DBHost            string
	DBPort            string
	DBUser            string
	DBPassword        string
	DBName            string
	DBOptions         string
	MaxConnection     int
	MaxIdleConnection int
	Driver            string
	DBFile            string
}

func NewSqlConnection(cfg DBConfiguration) (*sql.DB, error) {
	logger := zaplog.WithContext(context.Background())
	defer logger.Sync()

	var connURL string
	switch cfg.Driver {
	case MYSQL:
		connURL = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			cfg.DBUser,
			cfg.DBPassword,
			cfg.DBHost,
			cfg.DBPort,
			cfg.DBName)
	case POSTGRES:
		connURL = fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			cfg.DBUser,
			cfg.DBPassword,
			cfg.DBHost,
			cfg.DBPort,
			cfg.DBName)
	case SQLITE:
		connURL = cfg.DBFile
	}

	db, err := sql.Open(cfg.Driver, connURL)
	if err != nil {
		logger.Info("failed to connect to db", zap.Error(err))
		return nil, err
	}

	if cfg.MaxConnection > 0 {
		db.SetMaxOpenConns(cfg.MaxConnection)
	}
	if cfg.MaxIdleConnection > 0 {
		db.SetMaxIdleConns(cfg.MaxIdleConnection)
	}

	if err := db.Ping(); err != nil {
		logger.Info("failed to ping db", zap.Error(err))
		return nil, err
	}

	logger.Info("DB Says Pong!, DB connected")

	return db, nil
}
