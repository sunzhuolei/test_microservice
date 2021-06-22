package controller

import (
	"closer_user/internal/repo"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInfo(ctx * gin.Context){


	user,err := repo.GetUserByUserId("05b8793ea4904896ac8fd3414afd0fd7")
	if err != nil{
		fmt.Println(err)
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"code" : 0,
		"msg":"hello world",
		"data" : user,
	})
}