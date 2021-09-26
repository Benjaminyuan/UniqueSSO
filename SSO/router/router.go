package router

import (
	"unique/jedi/controller"
	"unique/jedi/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.Use(middleware.TracingMiddleware())
	r.Use(middleware.Cors())
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	// CAS related
	sr := r.Group("/cas")
	sr.POST("/login", controller.Login)
	sr.POST("/logout", controller.Logout)
	sr.GET("/p3/serviceValidate", controller.ValidateTicket)

	smsrouter := r.Group("/sms")
	smsrouter.POST("code", controller.SendSmsCode)

	qrrouter := r.Group("/qrcode")
	qrrouter.GET("code", controller.GetWorkWxQRCode)
}
