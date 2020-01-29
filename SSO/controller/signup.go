package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"unique/jedi/common"
	"unique/jedi/entity"
	"unique/jedi/service"
	"unique/jedi/session"
)

func SignUp(c *gin.Context){
	user := entity.User{}
	if err := c.ShouldBind(&user); err != nil {
		logrus.Errorf("Fail to bind user,err: %v", err)
		c.JSON(http.StatusOK, common.ParameterErrorResponse)
		return
	}
	//TODO: verify user info
	if err := service.CreateUser(c,&user);err != nil{
		logrus.Errorf("fail to createUser,err: %v",err)
		c.JSON(http.StatusInternalServerError,common.ServerErrorResponse)
		return
	}

	userSession,err := session.GlobalSessionManager.SessionStart()
	if err != nil {
		logrus.Errorf("fail to init session,err:%v",err)
		c.JSON(http.StatusInternalServerError,common.ServerErrorResponse)
		return
	}
	common.InjectUserInfoToSession(userSession,&user)
	_ = session.GlobalSessionManager.SessionUpdate(userSession.SessionID(),userSession)
	c.SetCookie(session.SID,userSession.SessionID(),session.DEFAULT_TIMEOUT,"/","",false,false)
	c.JSON(http.StatusOK,common.SuccessResponse)
	return
}
func SignUpHTML(c *gin.Context){

}