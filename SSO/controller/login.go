package controller

import (
	"crypto/md5"
	"github.com/gin-gonic/gin"
	"net/http"
	"unique/jedi/common"
	"unique/jedi/conf"
	"unique/jedi/service"
	"unique/jedi/session"
)

func Login(c *gin.Context){
	userName := c.PostForm(common.LoginUserNameKey)
	password := c.PostForm(common.LoginUserPasswordKey)
	if userName=="" || password=="" {
		c.JSON(http.StatusForbidden,common.ParameterErrorResponse)
		return
	}
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(password))
	cipherStr := md5Ctx.Sum([]byte(conf.SSOConf.MD5Sum))
	res,err := service.VerifyUser(c,userName,string(cipherStr))
	if err != nil {
		c.JSON(http.StatusInternalServerError,common.ServerErrorResponse)
		return
	}
	if res.Basic.Code != 0 || res.User == nil {
		c.JSON(http.StatusForbidden,common.NewErrorResponse(res.Basic.Info))
		return
	}
	user := res.User
	userSession,err := session.GlobalSessionManager.SessionStart()
	if err != nil {
		c.JSON(http.StatusInternalServerError,common.ServerErrorResponse)
		return
	}
	userSession.Set("name",user.Name)
	userSession.Set("uid",user.Uid)
	userSession.Set("college",user.College)
	_ = session.GlobalSessionManager.SessionUpdate(userSession.SessionID(),userSession)
	c.SetCookie(session.SID,userSession.SessionID(),session.DEFAULT_TIMEOUT,"/","",false,false)
	c.JSON(http.StatusOK,common.SuccessResponse)
}