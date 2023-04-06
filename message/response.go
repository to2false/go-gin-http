package message

import (
	"context"
)

type (
	ResponseInterface interface {
		Name() string

		StatusCode() int

		WithContext(ctx context.Context) ResponseInterface

		WithReasonPhrase(reason error) ResponseInterface

		GetBody() ([]byte, error)
	}
)

var registeredResponse = make(map[string]ResponseInterface)

func RegisterDefinedResponse(response ResponseInterface) {
	if response == nil {
		panic("cannot register a nil response")
	}
	if response.Name() == "" {
		panic("cannot register response with empty string result for Name()")
	}
	name := response.Name()
	registeredResponse[name] = response
}

func GetDefinedResponse(name string) ResponseInterface {
	return registeredResponse[name]
}
