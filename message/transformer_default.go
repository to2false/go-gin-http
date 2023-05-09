package message

import (
	"context"
	"github.com/to2false/go-gin-http/encoding"
	"github.com/to2false/go-gin-http/encoding/json"
	"google.golang.org/protobuf/proto"
	"net/http"
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

func (t DefaultTransformer) ContentType() string {
	return "application/json"
}

func (t DefaultTransformer) PreProcessRequest(r *http.Request) error {
	return nil
}

func (DefaultTransformer) Transform(ctx context.Context, response proto.Message) (int, []byte, error) {
	data, err := encoding.GetCodec(json.Name).Marshal(response)

	return http.StatusOK, data, err
}

func (t DefaultTransformer) Err(ctx context.Context, err error) (int, []byte, error) {
	data, e := encoding.GetCodec(json.Name).Marshal(err)

	return http.StatusOK, data, e
}
