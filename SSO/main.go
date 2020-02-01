package main
import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"unique/jedi/conf"
	"unique/jedi/controller"
	"unique/jedi/middleware"
	"unique/jedi/service"
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
	err = service.TestRpc()
	if err !=nil{
		log.Fatalf("rpc test failed ,err",err)
	}
	r := gin.Default()
	r.Use(middleware.Cors())
	r.Use(middleware.Sessions(session.GlobalSessionManager))
	r.GET("/template/signup",controller.SignUpHTML)
	r.GET("/template/login",controller.LoginHTML)
	r.POST("/signup",controller.SignUp)
	r.POST("/login",controller.Login)
	if err := r.Run();err != nil{
		log.Errorf("Fail to start the server,err: %v", err)
	}

}