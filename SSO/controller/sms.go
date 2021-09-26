package controller

import (
	"errors"
	"net/http"
	"unique/jedi/common"
	"unique/jedi/pkg"
	"unique/jedi/util"

	"github.com/gin-gonic/gin"
	"github.com/xylonx/zapx"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.uber.org/zap"
)

func SendSmsCode(ctx *gin.Context) {
	apmCtx, span := util.Tracer.Start(ctx.Request.Context(), "SendSMSCode")
	defer span.End()

	login := new(pkg.LoginUser)
	if err := ctx.ShouldBindJSON(login); err != nil {
		zapx.WithContext(apmCtx).Error("bind body failed", zap.Error(err))
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		ctx.JSON(http.StatusBadRequest, pkg.InvalidRequest(err))
		return
	}
	zapx.WithContext(apmCtx).Info("bind body successfully")
	span.SetAttributes(attribute.Any("body", login))

	if login.Phone == "" {
		ctx.JSON(http.StatusBadRequest, pkg.InvalidRequest(errors.New("no phone number specified")))
		return
	}

	code, err := util.GenerateSMSCode(apmCtx, login.Phone)
	if err != nil {
		zapx.WithContext(apmCtx).Error("generate sms code failed", zap.Error(err))
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		ctx.JSON(http.StatusBadRequest, pkg.InvalidRequest(err))
		ctx.JSON(http.StatusInternalServerError, pkg.InternalError(err))
		return
	}

	status, err := util.SendSMS(apmCtx, login.Phone, code, common.SMS_CODE_EXPIRES)
	if err != nil {
		zapx.WithContext(apmCtx).Error("send sms code failed", zap.Error(err))
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		ctx.JSON(http.StatusInternalServerError, pkg.InternalError(err))
		return
	}

	span.SetAttributes(attribute.Any("sms send status", status))

	if status != nil && len(*status) > 0 {
		ctx.JSON(http.StatusInternalServerError, pkg.InternalError(errors.New((*status)[0].Message)))
		return
	}

	ctx.JSON(http.StatusOK, pkg.CommonResponse{})
}
