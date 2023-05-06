package rest

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ResponseStruct struct {
	Field1 string `json:"field1"`
	Field2 int    `json:"field2"`
}

func TestClient_GET_OK(t *testing.T) {
	mockClient := NewMockedRestClient()
	client := NewClient("http://example.com", nil, mockClient)

	expectedURL := "http://example.com/endpoint"
	expectedResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewBufferString(`{"field1":"value1","field2":2}`)),
	}
	mockClient.On("Do", mock.Anything).Return(expectedResponse, nil)

	response := &ResponseStruct{}
	err := client.Get(context.Background(), "/endpoint", response)

	actualRequest := mockClient.Calls[0].Arguments.Get(0).(*http.Request)
	assert.NoError(t, err)
	assert.Equal(t, "GET", actualRequest.Method)
	assert.Equal(t, expectedURL, actualRequest.URL.String())
	assert.Equal(t, response.Field1, "value1")
	assert.Equal(t, response.Field2, 2)
}

func TestClient_GET_Error(t *testing.T) {
	mockClient := NewMockedRestClient()
	client := NewClient("http://example.com", nil, mockClient)

	expectedURL := "http://example.com/endpoint"
	mockClient.On("Do", mock.Anything).Return(nil, errors.New("some error"))

	response := &ResponseStruct{}
	err := client.Get(context.Background(), "/endpoint", response)

	actualRequest := mockClient.Calls[0].Arguments.Get(0).(*http.Request)
	assert.Error(t, err)
	assert.Equal(t, "GET", actualRequest.Method)
	assert.Equal(t, expectedURL, actualRequest.URL.String())
}

func TestClient_POST_Ok(t *testing.T) {
	mockClient := NewMockedRestClient()
	client := NewClient("http://example.com", nil, mockClient)

	expectedURL := "http://example.com/endpoint"
	expectedResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewBufferString(`{"field1":"value1","field2":2}`)),
	}
	mockClient.On("Do", mock.Anything).Return(expectedResponse, nil)

	response := &ResponseStruct{}
	err := client.Post(context.Background(), "/endpoint", nil, response)

	actualRequest := mockClient.Calls[0].Arguments.Get(0).(*http.Request)
	assert.NoError(t, err)
	assert.Equal(t, "POST", actualRequest.Method)
	assert.Equal(t, expectedURL, actualRequest.URL.String())
	assert.Equal(t, response.Field1, "value1")
	assert.Equal(t, response.Field2, 2)
}

func TestClient_POST_Error(t *testing.T) {
	mockClient := NewMockedRestClient()
	client := NewClient("http://example.com", nil, mockClient)

	expectedURL := "http://example.com/endpoint"
	mockClient.On("Do", mock.Anything).Return(nil, errors.New("some error"))

	response := &ResponseStruct{}
	err := client.Post(context.Background(), "/endpoint", nil, response)

	actualRequest := mockClient.Calls[0].Arguments.Get(0).(*http.Request)
	assert.Error(t, err)
	assert.Equal(t, "POST", actualRequest.Method)
	assert.Equal(t, expectedURL, actualRequest.URL.String())
}
