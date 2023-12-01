package service

import (
	"testing"

	"github.com/seigaalghi/e-library/mocks"
	"github.com/stretchr/testify/assert"
)

func getService(t *testing.T) (Service, *mocks.Repository, *mocks.HttpClient) {
	mockRepo := mocks.NewRepository(t)
	mockHttp := mocks.NewHttpClient(t)

	return NewService(mockRepo, mockHttp), mockRepo, mockHttp
}

func TestNewService(t *testing.T) {
	ms, mr, mh := getService(t)

	assert.NotNil(t, ms)
	assert.NotNil(t, mr)
	assert.NotNil(t, mh)
}
