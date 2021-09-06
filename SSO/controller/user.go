package controller

import (
	"errors"
	"net/http"
	"net/url"
	"time"
	"unique/jedi/common"
	"unique/jedi/conf"
	"unique/jedi/pkg"
	"unique/jedi/service"
	"unique/jedi/util"

	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

/*
	query param:
	type: phone / sms / email / wechat
	service[option]

	1. phone number with password
    body:
    {
        "phone": "",
        "password": ""
    }

2. phone sms
    body:
    {
        "phone": "",
        "code": ""
    }

3. email address with password
    body:
    {
        "email": "",
        "password": ""
    }
*/

func Login(ctx *gin.Context) {

	r, err := reporter.NewGRPCReporter("")
	t, err := go2sky.NewTracer("", go2sky.WithReporter(r))
	t.CreateLocalSpan(context.Background())

	signType, ok := ctx.GetQuery("type")
	if !ok {
		ctx.JSON(http.StatusBadRequest, pkg.InvalidRequest(errors.New("unsupported login type: "+signType)))
		return
	}

	target := &url.URL{
		Path: "/",
	}
	if redirectUrl, ok := ctx.GetQuery("service"); ok && redirectUrl != "" {
		if service.VerifyService(redirectUrl) != nil {
			ctx.JSON(http.StatusUnauthorized, pkg.InvalidService(errors.New("unsupported service: "+redirectUrl)))
			return
		}
		ru, err := url.Parse(redirectUrl)
		if err != nil {
			logrus.WithField("service", redirectUrl).WithError(err).Error("failed to parse service url")
			ctx.JSON(http.StatusBadRequest, pkg.InvalidRequest(errors.New("service格式错误")))
			return
		}
		target = ru
	}

	data := new(pkg.LoginUser)
	if err := ctx.ShouldBindJSON(data); err != nil {
		logrus.WithError(err).Error("post body format failed")
		ctx.JSON(http.StatusBadRequest, pkg.InvalidRequest(errors.New("上传数据格式错误")))
		return
	}

	// validate user
	user, err := service.VerifyUser(ctx.Request.Context(), data, signType)
	if err != nil {
		logrus.WithError(err).Error("validate user failed")
		ctx.JSON(http.StatusUnauthorized, pkg.InvalidTicketSpec(err))
		return
	}

	// new ticket, store and set cookie
	tgt := util.NewTGT()
	if err := service.StoreValue(ctx.Request.Context(), tgt, user.UID, common.CAS_TGT_EXPIRES); err != nil {
		logrus.WithError(err).Error("store tgt failed")
		ctx.JSON(http.StatusInternalServerError, pkg.InternalError(errors.New("服务器错误，请稍后尝试")))
		return
	}
	ctx.SetCookie(common.CAS_COOKIE_NAME, tgt, int(common.CAS_TGT_EXPIRES/time.Second), "/", ctx.Request.Host, true, true)

	ticket := util.NewTicket()
	if err := service.StoreValue(ctx.Request.Context(), ticket, user.UID, common.CAS_TICKET_EXPIRES); err != nil {
		logrus.WithError(err).Error("store ticket failed")
		ctx.JSON(http.StatusInternalServerError, pkg.InternalError(errors.New("服务器错误，请稍后尝试")))
		return
	}

	query := target.Query()
	query.Set("ticket", ticket)
	target.RawQuery = query.Encode()
	ctx.JSON(http.StatusOK, pkg.RedirectSuccess(target.String()))

	// FIXME
	logrus.Info("function return")
}

// TODO: construct a watcher to implement logout function
func Logout(ctx *gin.Context) {

}

func Register(ctx *gin.Context) {

}

func GetWorkWxQRCode(ctx *gin.Context) {
	if conf.SSOConf.Application.Mode == "debug" {
		src := "https://open.work.weixin.qq.com/wwopen/sso/qrImg?key=2d2287cf9cc95a8"
		ctx.JSON(http.StatusOK, pkg.QrcodeSuccess(src))
		return
	}

	src, err := util.GetQRCodeSrc()
	if err != nil {
		logrus.WithError(err).Error("get work wx qrcode failed")
		ctx.JSON(http.StatusInternalServerError, pkg.InternalError(errors.New("获取二维码错误")))
		return
	}
	ctx.JSON(http.StatusOK, pkg.QrcodeSuccess(src))
}
