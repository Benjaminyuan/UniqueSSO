package router

import (
	"unique/jedi/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// CAS related
	sr := r.Group("/cas")
	sr.POST("/login", controller.Login)
	sr.POST("/logout", controller.Logout)
	sr.GET("/p3/serviceValidate", controller.ValidateTicket)

	// register function
	r.POST("/register", controller.Register)

	smsrouter := r.Group("/sms")
	smsrouter.POST("code", controller.SendSmsCode)

	qrrouter := r.Group("/qrcode")
	qrrouter.GET("code", controller.GetWorkWxQRCode)
}
