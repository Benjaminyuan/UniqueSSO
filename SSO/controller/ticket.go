package controller

import (
	"errors"
	"net/http"
	"unique/jedi/pkg"
	"unique/jedi/service"

	"github.com/gin-gonic/gin"
)

func ValidateTicket(ctx *gin.Context) {
	ticket, ok := ctx.GetQuery("ticket")
	if !ok {
		ctx.JSON(http.StatusBadRequest, pkg.InvalidRequest(errors.New("no ticket available")))
		return
	}
	serviceName, ok := ctx.GetQuery("service")
	if !ok {
		ctx.JSON(http.StatusBadRequest, pkg.InvalidRequest(errors.New("no service available")))
		return
	}
	if service.VerifyService(serviceName) != nil {
		ctx.JSON(http.StatusUnauthorized, pkg.InvalidService(errors.New("unsupported service: "+serviceName)))
		return
	}

	uid, err := service.GetDelValue(ctx.Request.Context(), ticket)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, pkg.InvalidTicket(err))
		return
	}

	// return success json data
	user, err := service.GetUserById(uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.InternalError(err))
		return
	}

	ctx.JSON(http.StatusOK, pkg.AuthSuccess(user))
}
