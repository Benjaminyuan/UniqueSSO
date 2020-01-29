package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"unique/jedi/entity"
	pb "unique/jedi/protos/gencode"
	r "unique/jedi/rpc"
)

func CreateUser(c *gin.Context, user *entity.User)error{
	req := &pb.CreateUserRequest{Name: user.Name, Phone: user.Phone, EMail: user.EMail, College: user.College}
	client, err := r.NewUserClient()
	if err != nil {
		logrus.Errorf("fail to NewUserClient, err: %v", err)
		return  err
	}
	res, err := client.CreateUser(c, req)
	if err != nil {
		logrus.Errorf(" rpc fail to CreateUser,err:%v", err)
		return  err
	}
	if res.Basic.Code != 0 {
		logrus.Errorf(" service fail to CreateUser,err:%v",res.Basic.Info)
		return errors.New(res.Basic.Info)
	}
	user.UID = res.User.Uid
	return nil
}
func VerifyUser(c *gin.Context, userName string, password string) (*entity.User, error) {
	req := &pb.VerifyUserIdentityRequest{
		UserName: userName,
		Password: password,
	}
	client, err := r.NewUserClient()
	if err != nil {
		logrus.Errorf("fail to NewUserClient, err: %v", err)
		return nil, err
	}
	res, err := client.VerifyUserIdentity(c, req)
	if err != nil {
		logrus.Errorf("fail to VerifyUser,err:%v", err)
		return nil, err
	}
	if res.Basic.Code != 0 || res.User == nil {
		logrus.Errorf("fail to find user,err:%v", res.Basic.Info)
		return nil, errors.New(res.Basic.Info)
	}
	return &entity.User{UID: res.User.Uid, Name: res.User.Name, Phone: res.User.Phone, EMail: res.User.EMail, College: res.User.EMail}, err
}
