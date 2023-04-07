package message

import "context"

type (
	// DefaultTransformer Do nothing
	DefaultTransformer struct {
	}
)

var _ Transformer = (*DefaultTransformer)(nil)

const (
	DefaultTransformerName = "DefaultTransformerName"
)

func init() {
	RegisterTransformer(DefaultTransformer{})
}

func (t DefaultTransformer) Name() string {
	return DefaultTransformerName
}

func (DefaultTransformer) Transform(ctx context.Context, response any) any {
	return response
}
