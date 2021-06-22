package transport

import (
	endpoints "closer_user/internal/endpoint"
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

func MewGrpcHandler(endpoints endpoints.EndPoints) grpctransport.Handler {

	handler := grpctransport.NewServer(
		endpoints.GetUserInfoEndPoint,
		DecodeRequest,
		EncodeResponse,
	)

	return handler
}

/**
decode请求
 */
func DecodeRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

/**
encode响应
 */
func EncodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp, nil
}
