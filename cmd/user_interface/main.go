package main

import (
	api_user_v1 "closer_user/api/user/v1"
	"closer_user/config"
	"closer_user/internal/dao"
	endpoints "closer_user/internal/endpoint"
	"closer_user/internal/pkg/database"
	"closer_user/internal/server"
	"closer_user/internal/service"
	"closer_user/internal/transport"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
)

func main(){
	//载入配置文件
	config.LoadConfig()
	////链接数据库
	database.InitDataBase()
	////链接redis
	//redis.InitRedis()
	////api路由
	//router.InitRouter()

	//todo 注册
	userServer := server.UserServer{}
	userEndpoint := endpoints.NewUserInfoEndpoints(service.NewUserService(dao.NewUserDao()))
	userServer.Handler = transport.MewGrpcHandler(userEndpoint)
	//fmt.Println(userServer)
	serviceAddress := config.GetConfig().Server.Grpc.Addr
	fmt.Println(serviceAddress)
	ls,_ := net.Listen("tcp",serviceAddress)
	gs := grpc.NewServer()
	api_user_v1.RegisterUserServiceServer(gs,&userServer)
	grpc_health_v1.RegisterHealthServer(gs, health.NewServer())
	gs.Serve(ls)

}