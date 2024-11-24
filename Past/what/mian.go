package main

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	h "github.com/go-kit/kit/transport/http"
)

// Service接口定义
type StringService interface {
	Uppercase(string) (string, error)
}

// 服务实现
type stringService struct{}

func (stringService) Uppercase(s string) (string, error) {
	return strings.ToUpper(s), nil
}

// 请求和响应类型
type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

// 创建endpoint
func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(uppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return uppercaseResponse{V: v, Err: err.Error()}, nil
		}
		return uppercaseResponse{V: v, Err: ""}, nil
	}
}

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	svc := stringService{}
	uppercaseHandler := h.NewServer(
		makeUppercaseEndpoint(svc),
		h.DecodeJSONRequest,
		h.EncodeJSONResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}
