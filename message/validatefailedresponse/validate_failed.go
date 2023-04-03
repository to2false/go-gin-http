package validatefailedresponse

import (
	"context"
	"encoding/json"
	"github.com/to2false/go-gin-http/message"
	"net/http"
)

var (
	_ message.ResponseInterface = (*FailedResponse)(nil)
)

const (
	Name = "ValidateFailed"
)

func init() {
	message.RegisterDefinedResponse(FailedResponse{})
}

type (
	FailedResponse struct {
		ctx context.Context

		Code   int32       `json:"code"`
		Msg    string      `json:"msg"`
		Data   interface{} `json:"data"`
		Reason error       `json:"reason"`
	}
)

func (v FailedResponse) Name() string {
	return Name
}

func (v FailedResponse) StatusCode() int {
	return http.StatusOK
}

func (v FailedResponse) GetBody() ([]byte, error) {
	return json.Marshal(
		FailedResponse{
			ctx:    v.ctx,
			Code:   v.Code,
			Msg:    "参数验证不正确",
			Data:   make(map[any]any),
			Reason: v.Reason,
		},
	)
}

func (v FailedResponse) WithContext(ctx context.Context) message.ResponseInterface {
	v.ctx = ctx

	return v
}

func (v FailedResponse) WithReasonPhrase(reason error) message.ResponseInterface {
	v.Reason = reason

	return v
}
