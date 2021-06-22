package repo

import (
	"closer_user/internal/dao"
	"closer_user/internal/model"
)

func GetUserByUserId(userId string) (*model.CloserUser,error){
	//todo 查询缓存是否有数据
	return dao.NewUserDao().GetUserByUserId(userId)
}