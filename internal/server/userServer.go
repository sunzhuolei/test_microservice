package server

import (
	api_user_v1 "closer_user/api/user/v1"
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type UserServer struct {
	grpctransport.Handler
	//GetUserInfoHandler grpctransport.Handler
}



func (h  UserServer) Check(context.Context, *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING} , nil
}

func (h UserServer) Watch(*grpc_health_v1.HealthCheckRequest, grpc_health_v1.Health_WatchServer) error{
	return nil
}


/**
获取用户信息
 */
func(s *UserServer)GetUserInfo(ctx context.Context, req *api_user_v1.UserRequest) (*api_user_v1.UserResponse, error) {
	_, rsp, err := s.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rsp.(*api_user_v1.UserResponse),err
}