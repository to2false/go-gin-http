package message

import (
	"context"
	"google.golang.org/protobuf/proto"
)

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

func (DefaultTransformer) Transform(ctx context.Context, response proto.Message) any {
	return response
}
