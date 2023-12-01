package service

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/seigaalghi/e-library/mocks"
	"github.com/seigaalghi/e-library/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLendBook(t *testing.T) {
	type payload struct {
		ctx context.Context
		req *model.LendBookRequest
	}

	tests := []struct {
		scenario string
		doMock   func(payload, *mocks.Repository, *mocks.HttpClient)
		assert   func(t *testing.T, res *model.LendBookResponse, err error)
		payload  payload
	}{
		{
			scenario: "failed to get book",
			doMock: func(p payload, mr *mocks.Repository, mhc *mocks.HttpClient) {
				mhc.On("CallWithLog", p.ctx, mock.AnythingOfType("*http.Request"), mock.AnythingOfType("*model.Book")).
					Return(0, "", errors.New("error"))
			},
			assert: func(t *testing.T, res *model.LendBookResponse, err error) {
				assert.Nil(t, res)
				assert.NotNil(t, err)
			},
			payload: payload{
				ctx: context.Background(),
				req: &model.LendBookRequest{
					EditionNumber: "1234",
					PickupDate:    time.Now(),
					DropOffDate:   time.Now(),
				},
			},
		},
		{
			scenario: "failed to create schedules",
			doMock: func(p payload, mr *mocks.Repository, mhc *mocks.HttpClient) {
				mhc.On("CallWithLog", p.ctx, mock.AnythingOfType("*http.Request"), mock.AnythingOfType("*model.Book")).
					Return(200, "", nil).Run(func(args mock.Arguments) {
					arg := args.Get(2).(*model.Book)
					json.Unmarshal([]byte(GetBookResponseTest), &arg)
				})
				mr.On("CreateSchedule", p.req.PickupDate, p.req.DropOffDate, mock.AnythingOfType("*model.Books")).Return(false, errors.New("error"))
			},
			assert: func(t *testing.T, res *model.LendBookResponse, err error) {
				assert.Nil(t, res)
				assert.NotNil(t, err)
			},
			payload: payload{
				ctx: context.Background(),
				req: &model.LendBookRequest{
					EditionNumber: "1234",
					PickupDate:    time.Now(),
					DropOffDate:   time.Now(),
				},
			},
		},
		{
			scenario: "invalid schedules",
			doMock: func(p payload, mr *mocks.Repository, mhc *mocks.HttpClient) {
				mhc.On("CallWithLog", p.ctx, mock.AnythingOfType("*http.Request"), mock.AnythingOfType("*model.Book")).
					Return(200, "", nil).Run(func(args mock.Arguments) {
					arg := args.Get(2).(*model.Book)
					json.Unmarshal([]byte(GetBookResponseTest), &arg)
				})
				mr.On("CreateSchedule", p.req.PickupDate, p.req.DropOffDate, mock.AnythingOfType("*model.Books")).Return(false, nil)
			},
			assert: func(t *testing.T, res *model.LendBookResponse, err error) {
				assert.NotNil(t, res)
				assert.Nil(t, err)
			},
			payload: payload{
				ctx: context.Background(),
				req: &model.LendBookRequest{
					EditionNumber: "1234",
					PickupDate:    time.Now(),
					DropOffDate:   time.Now(),
				},
			},
		},
		{
			scenario: "success",
			doMock: func(p payload, mr *mocks.Repository, mhc *mocks.HttpClient) {
				mhc.On("CallWithLog", p.ctx, mock.AnythingOfType("*http.Request"), mock.AnythingOfType("*model.Book")).
					Return(200, "", nil).Run(func(args mock.Arguments) {
					arg := args.Get(2).(*model.Book)
					json.Unmarshal([]byte(GetBookResponseTest), &arg)
				})
				mr.On("CreateSchedule", p.req.PickupDate, p.req.DropOffDate, mock.AnythingOfType("*model.Books")).Return(true, nil)
			},
			assert: func(t *testing.T, res *model.LendBookResponse, err error) {
				assert.NotNil(t, res)
				assert.Nil(t, err)
			},
			payload: payload{
				ctx: context.Background(),
				req: &model.LendBookRequest{
					EditionNumber: "1234",
					PickupDate:    time.Now(),
					DropOffDate:   time.Now(),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.scenario, func(t *testing.T) {
			svc, mr, mh := getService(t)
			tt.doMock(tt.payload, mr, mh)
			res, err := svc.LendBook(tt.payload.ctx, tt.payload.req)
			tt.assert(t, res, err)
		})
	}
}

const (
	GetBookResponseTest = `{
		"works": [
			{
				"key": "/works/OL21177W"
			}
		],
		"title": "Wuthering Heights",
		"publishers": [
			"Nelson Doubleday"
		],
		"publish_date": "1850?",
		"key": "/books/OL38586477M",
		"type": {
			"key": "/type/edition"
		},
		"identifiers": {
			"goodreads": [
				"61419159"
			]
		},
		"covers": [
			12818862
		],
		"ocaid": "wutheringheights0000unse_b2b2",
		"oclc_numbers": [
			"26850653"
		],
		"classifications": {},
		"pagination": "xix, 290p.",
		"languages": [
			{
				"key": "/languages/eng"
			}
		],
		"publish_places": [
			"Garden City, New York"
		],
		"number_of_pages": 290,
		"physical_format": "Hardcover",
		"lc_classifications": [
			"PR4172.A5 W8"
		],
		"source_records": [
			"ia:wutheringheights0000unse_b2b2"
		],
		"latest_revision": 4,
		"revision": 4,
		"created": {
			"type": "/type/datetime",
			"value": "2022-07-10T15:14:13.783257"
		},
		"last_modified": {
			"type": "/type/datetime",
			"value": "2022-07-22T09:39:42.814367"
		}
	}`
)
