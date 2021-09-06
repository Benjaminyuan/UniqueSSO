package middleware

import (
	"net/http"
	"strconv"
	"strings"
	"time"
	"unique/jedi/util"

	"github.com/SkyAPM/go2sky"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	v3 "skywalking.apache.org/repo/goapi/collect/language/agent/v3"
)

// defined to identity third-part plugins
// for go: [5000, 6000)
const compomentID = 5050

const (
	spanHost = "http.host"
	spanURI  = "http.uri"
)

func APMMiddleware() gin.HandlerFunc {
	return skywalkingAPM
}

func skywalkingAPM(ctx *gin.Context) {
	oname := ctx.Request.Method + "\t" + ctx.FullPath()
	span, swCtx, err := util.Tracer.CreateEntrySpan(ctx.Request.Context(), oname, func(key string) (string, error) {
		return ctx.Request.Header.Get(key), nil
	})
	if err != nil {
		logrus.WithError(err).Error("failed to create span for request")
		ctx.Next()
		return
	}

	// before hook
	span.SetComponent(compomentID)
	span.Tag(go2sky.TagHTTPMethod, ctx.Request.Method)
	span.Tag(spanHost, ctx.Request.Host)
	span.Tag(spanURI, ctx.Request.URL.Path)
	spanHeader(&span, ctx.Request)
	span.SetSpanLayer(v3.SpanLayer_Http)

	ctx.Request = ctx.Request.WithContext(swCtx)

	ctx.Next()

	// after hook
	if len(ctx.Errors) > 0 {
		span.Error(time.Now(), ctx.Errors.String())
	}
	span.Tag(go2sky.TagStatusCode, strconv.Itoa(ctx.Writer.Status()))
	span.End()
}

func spanHeader(s *go2sky.Span, req *http.Request) {
	span := *s
	for key := range req.Header {
		span.Tag(go2sky.Tag("http."+strings.ToLower(key)), strings.Join(req.Header[key], ", "))
	}
}
