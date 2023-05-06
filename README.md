<div align="center">
 <h1>Disposable Mail Api for go</h1>
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
