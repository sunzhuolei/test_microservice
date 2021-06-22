package endpoint

import (
	api_user_v1 "closer_user/api/user/v1"
	"closer_user/internal/service"
	"context"
	"github.com/go-kit/kit/endpoint"
)

type EndPoints struct {
	GetUserInfoEndPoint endpoint.Endpoint
}

/**
实例化端点
 */
func NewUserInfoEndpoints(svc service.UserServiceInterface) EndPoints {
	var getUserInfoEndPoint endpoint.Endpoint
	{
		getUserInfoEndPoint = makeGetUserInfoEndPoint(svc)
	}
	return EndPoints{GetUserInfoEndPoint: getUserInfoEndPoint}
}

/**
创建端点
 */
func makeGetUserInfoEndPoint(s service.UserServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{},error) {
		req := request.(*api_user_v1.UserRequest)
		resp,err := s.GetUserInfo(ctx, req)
		return resp, err
	}
}