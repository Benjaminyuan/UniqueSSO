package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"unique/jedi/common"
	"unique/jedi/entity"
	"unique/jedi/service"
	"unique/jedi/session"
	"unique/jedi/util"
)

func SignUp(c *gin.Context){
	user := entity.User{}
	if err := c.ShouldBind(&user); err != nil {
		logrus.Errorf("Fail to bind user,err: %v", err)
		c.JSON(http.StatusBadRequest, common.ParameterErrorResponse)
		return
	}
	//TODO: verify user info
	user.Password = util.EncryptPassword(user.Password)
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
	InjectUserInfoToSession(userSession,&user)
	_ = session.GlobalSessionManager.SessionUpdate(userSession.SessionID(),userSession)
	c.SetCookie(session.SID,userSession.SessionID(),session.DEFAULT_TIMEOUT,"/","",false,false)
	c.JSON(http.StatusOK,common.SuccessResponse)
	return
}
func InjectUserInfoToSession(userSession session.Session,user *entity.User){
	userSession.Set(session.SessionUserNameKey,user.Name)
	userSession.Set(session.SessionUserUIDKey,user.UID)
	userSession.Set(session.SessionUserCollegeKey,user.College)
	userSession.Set(session.SessionUserEmailKey,user.EMail)
}
func SignUpHTML(c *gin.Context){
	c.File("./HTML/signup.html")

}