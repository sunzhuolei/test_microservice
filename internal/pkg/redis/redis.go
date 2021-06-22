package redis

import (
	"closer_user/config"
	"closer_user/internal/pkg/global"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

/**
初始化redis
 */
func InitRedis()bool{
	configure := config.GetConfig()
	pass := ""
	host := configure.Data.Redis.Addr

	conf := struct {
		Addr    string
		Pass    string
		DB      int
		ReadTimeout time.Duration
		WriteTimeout time.Duration
	}{
		Addr:    host,
		Pass:    pass,
		DB:      configure.Data.Redis.DB,
		ReadTimeout: configure.Data.Redis.ReadTimeout,
		WriteTimeout: configure.Data.Redis.WriteTimeout,
	}

	fmt.Println(conf)
	options := &redis.Options{
		Addr:         conf.Addr,
		Password:     conf.Pass,
		DB:           conf.DB,
		ReadTimeout:  conf.ReadTimeout,
		WriteTimeout: conf.WriteTimeout,
	}
	client := redis.NewClient(options)

	if _, err := client.Ping(context.Background()).Result(); err != nil {

		fmt.Println("redis连接失败:",err.Error())
		return false
	}
	log.Println("set redis config success")
	global.Redis = client
	return true
}