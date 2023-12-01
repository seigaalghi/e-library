package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func getRepository(t *testing.T) (Repository, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	return NewRepository(db), mock
}

func TestNewRepository(t *testing.T) {
	repo, mock := getRepository(t)

	assert.NotNil(t, repo)
	assert.NotNil(t, mock)
}
