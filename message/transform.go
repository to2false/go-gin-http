package message

import (
	"context"
	"google.golang.org/protobuf/proto"
	"net/http"
)

type (
	Transformer interface {
		Name() string
		ContentType() string
		PreProcessRequest(r *http.Request)
		Transform(ctx context.Context, response proto.Message) (int, []byte, error)
		Err(ctx context.Context, err error) (int, []byte, error)
	}
)

var registeredTransformer = make(map[string]Transformer)

func RegisterTransformer(transformer Transformer) {
	if transformer == nil {
		panic("cannot register a nil transformer")
	}
	if transformer.Name() == "" {
		panic("cannot register transformer with empty string result for Name()")
	}
	name := transformer.Name()
	registeredTransformer[name] = transformer
}

func GetTransformer(name string) Transformer {
	if v, has := registeredTransformer[name]; has {
		return v
	}

	return registeredTransformer[DefaultTransformerName]
}
