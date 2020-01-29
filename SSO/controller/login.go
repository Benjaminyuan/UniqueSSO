package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"unique/jedi/common"
	"unique/jedi/service"
	"unique/jedi/session"
	"unique/jedi/util"
)

func Login(c *gin.Context){
	userName := c.PostForm(common.LoginUserNameKey)
	password := c.PostForm(common.LoginUserPasswordKey)
	if userName=="" || password=="" {
		c.JSON(http.StatusForbidden,common.ParameterErrorResponse)
		return
	}
	cipherStr := util.EncryptPassword(password)
	user,err := service.VerifyUser(c,userName,string(cipherStr))
	if err != nil {
		logrus.Errorf(" service fail to verifyUser, err:%v",err)
		c.JSON(http.StatusInternalServerError,common.ServerErrorResponse)
		return
	}
	userSession,err := session.GlobalSessionManager.SessionStart()
	if err != nil {
		logrus.Errorf("fail to init session,err:%v",err)
		c.JSON(http.StatusInternalServerError,common.ServerErrorResponse)
		return
	}
	common.InjectUserInfoToSession(userSession,user)
	_ = session.GlobalSessionManager.SessionUpdate(userSession.SessionID(),userSession)
	c.SetCookie(session.SID,userSession.SessionID(),session.DEFAULT_TIMEOUT,"/","",false,false)
	c.JSON(http.StatusOK,common.SuccessResponse)
}

func LoginHTML(c *gin.Context){

}