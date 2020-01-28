package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"unique/jedi/entity"
	pb "unique/jedi/protos/gencode"
	r "unique/jedi/rpc"
)
func CreateUser(c *gin.Context,user *entity.User)( *pb.CreateUserResponse,error) {
	req := &pb.CreateUserRequest{Name:user.Name,Phone: user.Phone,EMail: user.EMail,College:user.College}
	client,err := r.NewUserClient()
	if err != nil {
		logrus.Fatal("fail to NewUserClient, err: %v",err)
		return nil, err
	}
	res,err := client.CreateUser(c,req)
	if err != nil {
		logrus.Fatalf("fail to CreateUser,err:%v",err)
		return nil,err
	}
	return res,nil
}
func VerifyUser(c *gin.Context,userName string,password string)(*pb.VerifyUserIdentityResponse,error){
	req := &pb.VerifyUserIdentityRequest{
		UserName:             userName,
		Password:             password,
	}
	client,err := r.NewUserClient()
	if err != nil{
		logrus.Fatal("fail to NewUserClient, err: %v",err)
		return nil,err
	}
	res,err := client.VerifyUserIdentity(c,req)
	if err != nil {
		logrus.Fatalf("fail to VerifyUser,err:%v",err)
		return nil,err
	}
	return res,err
}
