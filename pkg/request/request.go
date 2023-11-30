package request

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/seigaalghi/e-library/pkg/zaplog"
	"go.uber.org/zap"
)

type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

type HttpClient interface {
	Call(req *http.Request, response interface{}) (int, string, error)
	CallWithLog(ctx context.Context, req *http.Request, response interface{}) (int, string, error)
}

type httpclient struct {
	client Client
}

func NewClient(client Client) *httpclient {
	if client == nil {
		client = http.DefaultClient
	}

	return &httpclient{
		client: client,
	}
}

func (r *httpclient) Call(req *http.Request, response interface{}) (int, string, error) {
	res, err := r.client.Do(req)
	if err != nil {
		return 0, "", err
	}
	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	strBody := string(body)
	if err != nil {
		return res.StatusCode, strBody, err
	}

	// Decode the response body into the provided response object
	err = json.Unmarshal(body, &response)
	if err != nil {
		return res.StatusCode, strBody, err
	}

	return res.StatusCode, strBody, err
}

func (r *httpclient) CallWithLog(ctx context.Context, req *http.Request, response interface{}) (int, string, error) {
	logger := zaplog.WithContext(ctx)
	defer logger.Sync()

	res, err := r.client.Do(req)
	if err != nil {
		return 0, "", err
	}
	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	strBody := string(body)
	if err != nil {
		return res.StatusCode, strBody, err
	}

	logger.Info("raw response body", zap.Any("strBody", body))

	// Decode the response body into the provided response object
	err = json.Unmarshal(body, &response)
	if err != nil {
		return res.StatusCode, strBody, err
	}

	logger.Info("parsed response body", zap.Any("parsedBody", response))

	return res.StatusCode, strBody, err
}
