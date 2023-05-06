<div align="center">
 <h1>Rest Client package for Go!</h1>
</div>

The rest-client package provides a client implementation for making HTTP requests to RESTful APIs. It offers a convenient and straightforward way to interact with REST endpoints by encapsulating the underlying HTTP communication details.

The package includes a Client struct that allows you to configure the base URL, headers, and an HTTP client implementation. It provides methods for performing common HTTP methods such as GET, POST, PUT, and DELETE, along with the ability to pass request data and receive response data.

The Client struct uses an interface, HTTPClient, to abstract the HTTP client implementation, allowing you to easily swap different implementations, such as a real HTTP client or a mock client, for testing purposes.

By utilizing the rest-client package, you can simplify the process of sending HTTP requests to RESTful APIs, handle responses, and manage common functionality like setting headers or handling errors.


## Installation
```bash
go get github.com/bypepe77/rest-client-go
```

### Usage

```go
import disposable "github.com/bypepe77/disposable-mail-api/pkg"

type Response struct {
	Field1 string `json:"field1"`
	Field2 int    `json:"field2"`
}

func main() {
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %v", "ApiKey"),
	}
	restClient := rest.NewClient("http://example.com", headers, nil)

	response := &Response{}
	restClient.Get(context.Background(), "/endpoint", response)

	fmt.Println(response.Field1)
}
  
}
```

## Testing

The provided code defines a MockRestClient struct, which serves as a simulated HTTP client for testing purposes. It allows you to mimic the behavior of a real HTTP client without making actual network requests.

The MockRestClient implements the Do method from the HTTPClient interface. It utilizes the github.com/stretchr/testify/mock package to extend the mock.Mock struct and set expectations for its behavior.

The Do method takes an http.Request object as input and returns an http.Response and an error. By using the Called, args.Get, and args.Error methods of the mock.Mock struct, the method records the call and retrieves the mocked response and error values specified during test setup.

Additionally, the NewMockedRestClient function creates a new instance of MockRestClient for easy instantiation in tests.

In summary, the MockRestClient enables you to create a simulated HTTP client, facilitating controlled testing of code that interacts with HTTP clients. It helps isolate and verify specific behaviors without relying on actual network communication, making unit and integration testing more straightforward.

Example test: 

```go
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
```
