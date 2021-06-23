package main

import (
	api_user_v1 "closer_user/api/user/v1"
	"closer_user/config"
	"context"
	"fmt"
	stdconsul "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"strconv"
	"time"
)

func main(){

	//载入配置
	config.LoadConfig()

	//服务发现
	configs := stdconsul.DefaultConfig()
	configs.Address = config.GetConfig().Server.Consul.Addr
	fmt.Println(configs.Address)
	client, err := stdconsul.NewClient(configs)
	if err != nil {
		fmt.Println("consul client error : ", err)
		return
	}
	serviceName := config.GetConfig().Service.Register.Name
	tags :=  config.GetConfig().Service.Register.Tag
	serviceHealthy, _, err := client.Health().ServiceMultipleTags(serviceName, tags, true, nil)
	if len(serviceHealthy) <=0 {
		fmt.Println("未找到相关服务 : ", err)
		return
	}

	//todo 负载均衡
	serviceAddress := serviceHealthy[0].Service.Address+":"+strconv.Itoa(serviceHealthy[0].Service.Port)
	conn,err := grpc.Dial(serviceAddress,grpc.WithInsecure())
	if err != nil {
		panic("connect error")
	}
	defer conn.Close()
	userClient := api_user_v1.NewUserServiceClient(conn)
	res,err := userClient.GetUserInfo(context.Background(),&api_user_v1.UserRequest{UserId: "05b8793ea4904896ac8fd3414afd0fd7"})
	fmt.Println("请求服务: ", serviceAddress, "当前时间: ", time.Now().Format("2006-01-02 15:04:05.99"))
	fmt.Println(res)
	fmt.Println(err)



}

