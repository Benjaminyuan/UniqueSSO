package util

import (
	"unique/jedi/conf"
	"unique/jedi/pb/sms"

	"github.com/xylonx/zapx"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var OpenClient sms.SMSServiceClient

func SetupUtils() error {
	if err := setupOpenPlatformGrpc(); err != nil {
		zapx.Error("set up open platform grpc client failed", zap.Error(err))
		return err
	}
	return nil
}

func setupOpenPlatformGrpc() (err error) {
	var c *grpc.ClientConn
	var creds credentials.TransportCredentials
	if conf.SSOConf.Application.Mode == "dev" {
		creds, err = credentials.NewClientTLSFromFile(
			conf.SSOConf.OpenPlatform.GrpcCert,
			conf.SSOConf.OpenPlatform.GrpcServerName,
		)
		if err != nil {
			zapx.Error("create client tls from file failed", zap.Error(err))
			return err
		}
		c, err = grpc.Dial(
			conf.SSOConf.OpenPlatform.GrpcAddr,
			grpc.WithTransportCredentials(creds),
			grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
			grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
		)
	} else {
		c, err = grpc.Dial(
			conf.SSOConf.OpenPlatform.GrpcAddr,
			grpc.WithInsecure(),
			grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
			grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
		)
	}
	if err != nil {
		zapx.Error("dial to open platform failed", zap.Error(err))
		return err
	}

	OpenClient = sms.NewSMSServiceClient(c)
	return nil
}
