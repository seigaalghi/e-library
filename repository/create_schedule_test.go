package repository

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/seigaalghi/e-library/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateSchedule(t *testing.T) {
	type payload struct {
		pickup  time.Time
		dropoff time.Time
		book    *model.Books
	}
	tests := []struct {
		scenario string
		payload  payload
		doMock   func(p payload, sm sqlmock.Sqlmock)
		assert   func(t *testing.T, res bool, err error)
	}{
		{
			scenario: "failed get count",
			payload: payload{
				pickup:  time.Now(),
				dropoff: time.Now(),
				book: &model.Books{
					Title: "ancient",
					Authors: []model.Authors{
						{
							Key:  "1234",
							Name: "Seiga",
						},
					},
					EditionNumber: "12345",
				},
			},
			doMock: func(p payload, sm sqlmock.Sqlmock) {
				sm.ExpectQuery(bookAvailability).WithArgs(p.book.EditionNumber).WillReturnError(errors.New("error"))
			},
			assert: func(t *testing.T, res bool, err error) {
				assert.NotNil(t, err)
				assert.False(t, res)
			},
		},
		{
			scenario: "already booked",
			payload: payload{
				pickup:  time.Now(),
				dropoff: time.Now(),
				book: &model.Books{
					Title: "ancient",
					Authors: []model.Authors{
						{
							Key:  "1234",
							Name: "Seiga",
						},
					},
					EditionNumber: "12345",
				},
			},
			doMock: func(p payload, sm sqlmock.Sqlmock) {
				sm.ExpectQuery(bookAvailability).WithArgs(p.book.EditionNumber).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
			},
			assert: func(t *testing.T, res bool, err error) {
				assert.Nil(t, err)
				assert.False(t, res)
			},
		},
		{
			scenario: "failed inserting schedule",
			payload: payload{
				pickup:  time.Now(),
				dropoff: time.Now(),
				book: &model.Books{
					Title: "ancient",
					Authors: []model.Authors{
						{
							Key:  "1234",
							Name: "Seiga",
						},
					},
					EditionNumber: "12345",
				},
			},
			doMock: func(p payload, sm sqlmock.Sqlmock) {
				sm.ExpectQuery(bookAvailability).WithArgs(p.book.EditionNumber).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				sm.ExpectExec(createScheduleQuery).WithArgs(p.book.Title, p.book.EditionNumber, sqlmock.AnyArg(), p.pickup, p.dropoff).WillReturnError(errors.New("error"))
			},
			assert: func(t *testing.T, res bool, err error) {
				assert.NotNil(t, err)
				assert.False(t, res)
			},
		},
		{
			scenario: "failed zero rows affected",
			payload: payload{
				pickup:  time.Now(),
				dropoff: time.Now(),
				book: &model.Books{
					Title: "ancient",
					Authors: []model.Authors{
						{
							Key:  "1234",
							Name: "Seiga",
						},
					},
					EditionNumber: "12345",
				},
			},
			doMock: func(p payload, sm sqlmock.Sqlmock) {
				sm.ExpectQuery(bookAvailability).WithArgs(p.book.EditionNumber).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				sm.ExpectExec(createScheduleQuery).WithArgs(p.book.Title, p.book.EditionNumber, sqlmock.AnyArg(), p.pickup, p.dropoff).
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			assert: func(t *testing.T, res bool, err error) {
				assert.NotNil(t, err)
				assert.False(t, res)
			},
		},
		{
			scenario: "success",
			payload: payload{
				pickup:  time.Now(),
				dropoff: time.Now(),
				book: &model.Books{
					Title: "ancient",
					Authors: []model.Authors{
						{
							Key:  "1234",
							Name: "Seiga",
						},
					},
					EditionNumber: "12345",
				},
			},
			doMock: func(p payload, sm sqlmock.Sqlmock) {
				sm.ExpectQuery(bookAvailability).WithArgs(p.book.EditionNumber).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				sm.ExpectExec(createScheduleQuery).WithArgs(p.book.Title, p.book.EditionNumber, sqlmock.AnyArg(), p.pickup, p.dropoff).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			assert: func(t *testing.T, res bool, err error) {
				assert.Nil(t, err)
				assert.True(t, res)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.scenario, func(t *testing.T) {
			svc, mock := getRepository(t)

			tt.doMock(tt.payload, mock)

			res, err := svc.CreateSchedule(tt.payload.pickup, tt.payload.dropoff, tt.payload.book)

			tt.assert(t, res, err)
		})
	}
}
