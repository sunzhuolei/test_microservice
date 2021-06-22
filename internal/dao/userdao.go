package dao

import (
	"closer_user/internal/model"
	"closer_user/internal/pkg/global"
)

type UserDao struct{

}

func NewUserDao() *UserDao{
	return &UserDao{}
}

/**
根据userId获取用户信息
 */
func (user *UserDao)GetUserByUserId(userId string)(model.CloserUser,error){
	modelUser := model.CloserUser{}
	err := global.Db.Table("closer_user").
		Where("user_id =?",userId).
		Where("is_delete = ?",0).
		First(&modelUser).Error
	return modelUser,err
}