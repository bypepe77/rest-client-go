package rest

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type MockRestClient struct {
	mock.Mock
}

func (c *MockRestClient) Do(request *http.Request) (*http.Response, error) {
	args := c.Called(request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*http.Response), args.Error(1)
}

func NewMockedRestClient() *MockRestClient {
	return &MockRestClient{}
}
