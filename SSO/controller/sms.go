package controller

import (
	"errors"
	"net/http"
	"unique/jedi/common"
	"unique/jedi/pkg"
	"unique/jedi/util"

	"github.com/gin-gonic/gin"
)

func SendSmsCode(ctx *gin.Context) {
	login := new(pkg.LoginUser)
	if err := ctx.ShouldBindJSON(login); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.InvalidRequest(err))
		return
	}
	if login.Phone == "" {
		ctx.JSON(http.StatusBadRequest, pkg.InvalidRequest(errors.New("no phone number specified")))
		return
	}

	code, err := util.GenerateSMSCode(ctx.Request.Context(), login.Phone)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.InternalError(err))
		return
	}

	status, err := util.SendSMS(login.Phone, code, common.SMS_CODE_EXPIRES)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.InternalError(err))
		return
	}
	if status != nil && len(*status) > 0 {
		ctx.JSON(http.StatusInternalServerError, pkg.InternalError(errors.New((*status)[0].Message)))
		return
	}

	ctx.JSON(http.StatusOK, pkg.CommonResponse{})
}
