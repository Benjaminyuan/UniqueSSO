package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"unique/jedi/session"
)

func Sessions( manager *session.Manager) gin.HandlerFunc{
	return func(c *gin.Context) {
		sid,err:= c.Cookie(session.SID)
		if err != nil || sid == ""{
			logrus.Fatalf("fail to get sid from cookie,err: %v",err)
			c.Redirect(http.StatusFound,"/login")
			return
		}
		_ , err = manager.SessionRead(sid)
		if err != nil {
			logrus.Fatalf("fail to read session,err: %v",err)
			c.Redirect(http.StatusFound,"/login")
		}
		c.Set(session.SID,sid)
		c.Next()
	}
}
