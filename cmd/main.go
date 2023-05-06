package main

import (
	"context"
	"fmt"

	"github.com/bypepe77/rest-client-go/pkg/rest"
)

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
