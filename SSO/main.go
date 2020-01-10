package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"unique/jedi/conf"
	"unique/jedi/middleware"
)


func main(){
	err := conf.InitConf()

	if err != nil{
		log.Fatalf("Fail to init conf,err: %v",err)
	}
	log.SetLevel(log.DebugLevel)
	log.Debugf("Conf:%v",*conf.SSOConf)
	r := gin.Default()
	r.Use(middleware.Cors())
}