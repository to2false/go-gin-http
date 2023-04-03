package internalerrresponse

import (
	"context"
	"encoding/json"
	"github.com/to2false/go-gin-http/message"
	"net/http"
)

var _ message.ResponseInterface = (*InternalErr)(nil)

const (
	Name = "InternalErr"
)

func init() {
	message.RegisterDefinedResponse(InternalErr{})
}

type InternalErr struct {
	ctx context.Context

	Code   int32       `json:"code"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
	Reason error       `json:"reason"`
}

func (s InternalErr) Name() string {
	return Name
}

func (s InternalErr) StatusCode() int {
	return http.StatusInternalServerError
}

func (s InternalErr) WithContext(ctx context.Context) message.ResponseInterface {
	s.ctx = ctx

	return s
}

func (s InternalErr) WithReasonPhrase(reason error) message.ResponseInterface {
	s.Reason = reason

	return s
}

func (s InternalErr) GetBody() ([]byte, error) {
	return json.Marshal(
		InternalErr{
			ctx:    s.ctx,
			Code:   s.Code,
			Msg:    "内部错误",
			Data:   make(map[any]any),
			Reason: s.Reason,
		},
	)
}
