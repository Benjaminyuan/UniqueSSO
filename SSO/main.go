package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"unique/jedi/conf"
	"unique/jedi/controller"
	"unique/jedi/middleware"
	"unique/jedi/session"
)


func main(){
	err := conf.InitConf()
	if err != nil{
		log.Fatalf("Fail to init conf, err: %v",err)
	}
	err = conf.InitDB()
	if err != nil{
		log.Fatalf("fail to init DB, err: %v",err)
	}
	err = session.InitSession()
	if err != nil {
		log.Fatalf("fail to init session,err",err)
	}
	log.SetLevel(log.DebugLevel)
	log.Debugf("Conf:%v",*conf.SSOConf)
	r := gin.Default()
	r.Use(middleware.Cors())
	r.Use(middleware.Sessions(session.GlobalSessionManager))
	r.POST("/signup",controller.SignUp)
}