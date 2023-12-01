package service

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/seigaalghi/e-library/mocks"
	"github.com/seigaalghi/e-library/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetBooks(t *testing.T) {
	type payload struct {
		ctx context.Context
		req *model.GetBooksRequest
	}

	tests := []struct {
		scenario string
		doMock   func(payload, *mocks.Repository, *mocks.HttpClient)
		assert   func(t *testing.T, res *model.GetBooksResponse, err error)
		payload  payload
	}{
		{
			scenario: "failed get books",
			doMock: func(p payload, mr *mocks.Repository, mhc *mocks.HttpClient) {
				mhc.On("CallWithLog", p.ctx, mock.AnythingOfType("*http.Request"), mock.AnythingOfType("*model.GetBookBySubjectResponse")).
					Return(0, "", errors.New("error"))
			},
			assert: func(t *testing.T, res *model.GetBooksResponse, err error) {
				assert.Nil(t, res)
				assert.NotNil(t, err)
			},
			payload: payload{
				ctx: context.Background(),
				req: &model.GetBooksRequest{
					Subject: "love",
					Page:    1,
					Limit:   15,
				},
			},
		},
		{
			scenario: "books not found",
			doMock: func(p payload, mr *mocks.Repository, mhc *mocks.HttpClient) {
				mhc.On("CallWithLog", p.ctx, mock.AnythingOfType("*http.Request"), mock.AnythingOfType("*model.GetBookBySubjectResponse")).
					Return(200, GetBooksNotFoundTest, nil).Run(func(args mock.Arguments) {
					arg := args.Get(2).(*model.GetBookBySubjectResponse)
					json.Unmarshal([]byte(GetBooksNotFoundTest), &arg)
				})
			},
			assert: func(t *testing.T, res *model.GetBooksResponse, err error) {
				assert.NotNil(t, res)
				assert.Nil(t, err)
			},
			payload: payload{
				ctx: context.Background(),
				req: &model.GetBooksRequest{
					Subject: "love",
					Page:    1,
					Limit:   15,
				},
			},
		},
		{
			scenario: "success",
			doMock: func(p payload, mr *mocks.Repository, mhc *mocks.HttpClient) {
				mhc.On("CallWithLog", p.ctx, mock.AnythingOfType("*http.Request"), mock.AnythingOfType("*model.GetBookBySubjectResponse")).
					Return(200, GetBooksSuccessTest, nil).Run(func(args mock.Arguments) {
					arg := args.Get(2).(*model.GetBookBySubjectResponse)
					json.Unmarshal([]byte(GetBooksSuccessTest), &arg)
				})
			},
			assert: func(t *testing.T, res *model.GetBooksResponse, err error) {
				assert.NotNil(t, res)
				assert.Nil(t, err)
			},
			payload: payload{
				ctx: context.Background(),
				req: &model.GetBooksRequest{
					Subject: "love",
					Page:    1,
					Limit:   15,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.scenario, func(t *testing.T) {
			svc, mr, mh := getService(t)
			tt.doMock(tt.payload, mr, mh)
			res, err := svc.GetBooks(tt.payload.ctx, tt.payload.req)
			tt.assert(t, res, err)
		})
	}
}

const (
	GetBooksNotFoundTest = `{
		"key": "/subjects/loveasd",
		"name": "loveasd",
		"subject_type": "subject",
		"work_count": 0,
		"works": []
	}`
	GetBooksSuccessTest = `{
		"key": "/subjects/love",
		"name": "love",
		"subject_type": "subject",
		"work_count": 16425,
		"works": [
			{
				"key": "/works/OL21177W",
				"title": "Wuthering Heights",
				"edition_count": 2113,
				"cover_id": 12818862,
				"cover_edition_key": "OL38586477M",
				"subject": [
					"British and irish fiction (fictional works by one author)",
					"Children's fiction"
				],
				"ia_collection": [
					"365-Books-by-Women-Authors",
					"additional_collections"					
				],
				"lendinglibrary": false,
				"printdisabled": true,
				"lending_edition": "OL39222415M",
				"lending_identifier": "wutheringheights0000emil_z8q3",
				"authors": [
					{
						"key": "/authors/OL24529A",
						"name": "Emily Brontë"
					}
				],
				"first_publish_year": 1847,
				"ia": "wutheringheights0000emil_z8q3",
				"public_scan": true,
				"has_fulltext": true,
				"availability": {
					"status": "open",
					"available_to_browse": false,
					"available_to_borrow": false,
					"available_to_waitlist": false,
					"is_printdisabled": false,
					"is_readable": true,
					"is_lendable": false,
					"is_previewable": true,
					"identifier": "wutheringheights0000emil_z8q3",
					"isbn": null,
					"oclc": null,
					"openlibrary_work": "OL21177W",
					"openlibrary_edition": "OL39222415M",
					"last_loan_date": null,
					"num_waitlist": null,
					"last_waitlist_date": null,
					"is_restricted": false,
					"is_browseable": false,
					"__src__": "core.models.lending.get_availability"
				}
			},
			{
				"key": "/works/OL468431W",
				"title": "The Great Gatsby",
				"edition_count": 1175,
				"cover_id": 10590366,
				"cover_edition_key": "OL22570129M",
				"subject": [
					"Married people, fiction"
				],
				"ia_collection": [
					"JaiGyan"
				],
				"lendinglibrary": false,
				"printdisabled": true,
				"lending_edition": "OL26491056M",
				"lending_identifier": "the-great-gatsby",
				"authors": [
					{
						"key": "/authors/OL27349A",
						"name": "F. Scott Fitzgerald"
					}
				],
				"first_publish_year": 1920,
				"ia": "the-great-gatsby",
				"public_scan": true,
				"has_fulltext": true,
				"availability": {
					"status": "open",
					"available_to_browse": false,
					"available_to_borrow": false,
					"available_to_waitlist": false,
					"is_printdisabled": false,
					"is_readable": true,
					"is_lendable": false,
					"is_previewable": true,
					"identifier": "the-great-gatsby",
					"isbn": null,
					"oclc": null,
					"openlibrary_work": "OL3871697W",
					"openlibrary_edition": "OL26491056M",
					"last_loan_date": null,
					"num_waitlist": null,
					"last_waitlist_date": null,
					"is_restricted": false,
					"is_browseable": false,
					"__src__": "core.models.lending.get_availability"
				}
			}
		]
	}`
)
