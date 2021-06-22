package main

import (
	api_user_v1 "closer_user/api/user/v1"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

func main(){

	serviceAddress := "192.168.2.85:50057"
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
	for{
		;
	}
}