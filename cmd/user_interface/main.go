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
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	stdconsul "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"os"
	"strconv"
	"strings"
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

	//服务注册
	consulAddr := config.GetConfig().Server.Consul.Addr
	consulSplits := strings.Split(consulAddr,":")
	consulHost := consulSplits[0]
	consulPort := consulSplits[1]
	serviceAddr := config.GetConfig().Server.Grpc.Addr
	serviceSplits := strings.Split(serviceAddr,":")
	serviceHost := serviceSplits[0]
	servicePort := serviceSplits[1]


	fmt.Println(serviceHost)
	registrar :=  NewRegistar(consulHost,consulPort,serviceHost,servicePort,log.NewNopLogger())
	registrar.Register()
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


func NewRegistar(consulHost string,consulPort string,serviceHost string,servicePort string,logger log.Logger)(registar sd.Registrar){
	//获取consul配置
	var client consul.Client
	{
		configs := stdconsul.DefaultConfig()
		configs.Address = consulHost + ":" + consulPort
		consulClient,err := stdconsul.NewClient(configs)
		if err != nil{
			logger.Log("create consul client error:", err)
			os.Exit(1)
		}
		client = consul.NewClient(consulClient)

	}

	// 设置Consul对服务健康检查的参数
	check := stdconsul.AgentServiceCheck{
		GRPC: fmt.Sprintf("%v:%v", serviceHost, servicePort),    //grpc健康检测 grpc.health.v1.Health
		Interval:  config.GetConfig().Service.Check.Interval,
		Timeout:  config.GetConfig().Service.Check.Timeout,
		Notes:    config.GetConfig().Service.Check.Notes,
		DeregisterCriticalServiceAfter:config.GetConfig().Service.Check.DeregisterCriticalServiceAfter,
	}

	servicePortInt,_ := strconv.Atoi(servicePort)
	//设置微服务想Consul的注册信息
	reg := stdconsul.AgentServiceRegistration{
		ID: config.GetConfig().Service.Register.Id,
		Name:    config.GetConfig().Service.Register.Name,    // 服务名称,
		Address: serviceHost,
		Port:    servicePortInt,
		Tags:    config.GetConfig().Service.Register.Tag,
		Check:   &check,
	}

	// 执行注册
	registar = consul.NewRegistrar(client, &reg, logger)
	return

}