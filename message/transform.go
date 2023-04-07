package message

import "context"

type (
	Transformer interface {
		Name() string
		Transform(ctx context.Context, response any) any
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
