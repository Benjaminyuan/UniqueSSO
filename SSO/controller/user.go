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

	"github.com/gin-gonic/gin"
	"github.com/xylonx/zapx"
	"go.uber.org/zap"
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
	apmCtx, span := util.Tracer.Start(ctx.Request.Context(), "Login")
	defer span.End()

	signType, ok := ctx.GetQuery("type")
	if !ok {
		zapx.WithContext(apmCtx).Error("sign type unsupported", zap.String("type", signType))
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
			zapx.WithContext(apmCtx).Error("failed to parse redirect url", zap.String("service", redirectUrl))
			ctx.JSON(http.StatusBadRequest, pkg.InvalidRequest(errors.New("service格式错误")))
			return
		}
		target = ru
	}

	data := new(pkg.LoginUser)
	if err := ctx.ShouldBindJSON(data); err != nil {
		zapx.WithContext(apmCtx).Error("post body format incorroct", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, pkg.InvalidRequest(errors.New("上传数据格式错误")))
		return
	}

	// validate user
	user, err := service.VerifyUser(ctx.Request.Context(), data, signType)
	if err != nil {
		zapx.WithContext(apmCtx).Error("validate user failed", zap.Error(err))
		ctx.JSON(http.StatusUnauthorized, pkg.InvalidTicketSpec(err))
		return
	}

	// new ticket, store and set cookie
	tgt := util.NewTGT()
	if err := service.StoreValue(ctx.Request.Context(), tgt, user.UID, common.CAS_TGT_EXPIRES); err != nil {
		zapx.WithContext(apmCtx).Error("store tgt failed", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, pkg.InternalError(errors.New("服务器错误，请稍后尝试")))
		return
	}
	ctx.SetCookie(common.CAS_COOKIE_NAME, tgt, int(common.CAS_TGT_EXPIRES/time.Second), "/", ctx.Request.Host, true, true)

	ticket := util.NewTicket()
	if err := service.StoreValue(ctx.Request.Context(), ticket, user.UID, common.CAS_TICKET_EXPIRES); err != nil {
		zapx.WithContext(apmCtx).Error("store ticket failed", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, pkg.InternalError(errors.New("服务器错误，请稍后尝试")))
		return
	}

	query := target.Query()
	query.Set("ticket", ticket)
	target.RawQuery = query.Encode()

	// append token

	ctx.Redirect(http.StatusFound, target.String())
}

// TODO: construct a watcher to implement logout function
func Logout(ctx *gin.Context) {

}

func GetWorkWxQRCode(ctx *gin.Context) {
	apmCtx, span := util.Tracer.Start(ctx.Request.Context(), "GetWorkWxQRCode")
	defer span.End()
	if conf.SSOConf.Application.Mode == "debug" {
		src := "https://open.work.weixin.qq.com/wwopen/sso/qrImg?key=2d2287cf9cc95a8"
		ctx.JSON(http.StatusOK, pkg.QrcodeSuccess(src))
		return
	}

	src, err := util.GetQRCodeSrc()
	if err != nil {
		zapx.WithContext(apmCtx).Error("get work wxQRCode failed", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, pkg.InternalError(errors.New("获取二维码错误")))
		return
	}
	ctx.JSON(http.StatusOK, pkg.QrcodeSuccess(src))
}
