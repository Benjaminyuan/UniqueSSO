package middleware

import(
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"unique/jedi/conf"
)
func originAllowed(origin string )bool{
	for _,val := range conf.SSOConf.OriginAllowedList{
		if val == origin{
			return true
		}
	}
	return false
}
func Cors() gin.HandlerFunc{
	return func(c *gin.Context){
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		log.Infof("Handle request from origin: %v",origin)
		if origin != "" && originAllowed(origin){
			c.Writer.Header().Set("Access-Control-allow-Origin",origin)
			c.Header("Access-Control-Allow-Methods","GET, POST, OPTIONS")
			c.Header("Access-Control-Allow-Headers","*")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar,Authorization")      // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")        // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")
		}
		if method == "OPTIONS"{
			c.String(http.StatusOK,"")
		}
		c.Next()
	}
}
