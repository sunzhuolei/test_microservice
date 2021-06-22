package router

import (
	"closer_user/internal/controller"
	"github.com/gin-gonic/gin"
)

/**
初始化路由bff层
 */
func InitRouter(){
	engine := gin.Default()
	middleWares := make([]gin.HandlerFunc, 0)
	apiGroup := engine.Group("/api").Use(middleWares...)
	{
		apiGroup.GET("/get_user", controller.GetUserInfo)
	}
	err := engine.Run(":8081")
	if err != nil{
		panic(err.Error())
	}
}