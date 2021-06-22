package service

import (
	api_user_v1 "closer_user/api/user/v1"
	"closer_user/internal/dao"
	"context"
	"encoding/json"
	"fmt"
)

/**
定义用户服务的接口
 */
type UserServiceInterface interface {
	GetUserInfo(ctx context.Context, in *api_user_v1.UserRequest) (*api_user_v1.UserResponse,error)
}

type UserService struct {
	dao *dao.UserDao
}

func NewUserService(dao *dao.UserDao) *UserService  {
	return &UserService {dao: dao}
}

func (s *UserService) GetUserInfo(ctx context.Context, in *api_user_v1.UserRequest)(*api_user_v1.UserResponse,error) {
	userModel,err := s.dao.GetUserByUserId(in.UserId)
	fmt.Println(userModel)
	if err != nil{
		return nil,err
	}

	bytes,err := json.Marshal(userModel)
	if err != nil{
		return nil,err
	}
	var res api_user_v1.UserResponse
	err = json.Unmarshal(bytes,&res)
	if err != nil{
		return nil,err
	}
	return &res,nil
}