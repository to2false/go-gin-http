package message

import "strings"

type (
	Transformer interface {
		Name() string
		Transform(response any) any
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
	name := strings.ToLower(transformer.Name())
	registeredTransformer[name] = transformer
}

func GetTransformer(name string) Transformer {
	if v, has := registeredTransformer[name]; has {
		return v
	}

	return registeredTransformer[DefaultTransformerName]
}
