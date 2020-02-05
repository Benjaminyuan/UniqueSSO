package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"unique/jedi/common"
	"unique/jedi/service"
	"unique/jedi/session"
	"unique/jedi/util"
)

func Login(c *gin.Context){
	redirectUrl ,ok := c.GetQuery(common.ServiceKey)
	if !ok {
		c.JSON(http.StatusBadRequest,common.ParameterErrorResponse)
		return
	}
	redirectUrl = fmt.Sprintf("%v?ticket=%v",redirectUrl,util.NewTicket())
	sid,err:= c.Cookie(session.SID)
	if err == nil && sid != ""{
		s , err := session.GlobalSessionManager.SessionRead(sid)
		if err == nil{
			 _   = session.GlobalSessionManager.SessionUpdate(sid,s)
			 c.SetCookie(session.SID,s.SessionID(),session.DEFAULT_TIMEOUT,"/","",false,false)
			 c.Redirect(http.StatusFound,redirectUrl)
			 //c.JSON(http.StatusOK,common.SuccessResponse)
			 return
		}
	}

	userName := c.PostForm(common.LoginUserNameKey)
	password := c.PostForm(common.LoginUserPasswordKey)
	if userName=="" || password=="" {
		c.JSON(http.StatusForbidden,common.ParameterErrorResponse)
		return
	}
	cipherStr := util.EncryptPassword(password)
	logrus.Debugf("password:%v,cipherStr:%v",password,cipherStr)
	user,err := service.VerifyUser(c,userName,cipherStr)
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
	InjectUserInfoToSession(userSession,user)
	_ = session.GlobalSessionManager.SessionUpdate(userSession.SessionID(),userSession)
	c.SetCookie(session.SID,userSession.SessionID(),session.DEFAULT_TIMEOUT,"/","",false,false)
	//c.JSON(http.StatusOK,common.SuccessResponse)
	c.Redirect(http.StatusFound,redirectUrl)
}

func LoginHTML(c *gin.Context){
	c.File("./HTML/login.html")
}