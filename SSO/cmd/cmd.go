package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
	"unique/jedi/common"
	"unique/jedi/conf"
	"unique/jedi/database"
	"unique/jedi/router"
	"unique/jedi/service"
	"unique/jedi/util"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/xylonx/zapx"
	zapxdecoder "github.com/xylonx/zapx/decoder"
	"go.uber.org/zap"
)

var rootCmd = &cobra.Command{
	Use:   "unique-sso",
	Short: "unique studio sso service",
	PreRun: func(c *cobra.Command, args []string) {
		setup()
	},
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

var configFilePath string

func init() {
	rootCmd.Flags().StringVarP(&configFilePath, "config", "c", "./conf/conf.yaml", "path to config file")
}

func Execute() error {
	return rootCmd.Execute()
}

func setup() {
	// init zapx
	logger, err := zap.NewProduction()
	if err != nil {
		os.Exit(1)
	}
	zapx.Use(logger, zapxdecoder.OpentelemetaryDecoder)

	if err := conf.InitConf(configFilePath); err != nil {
		zapx.Error("load config from file failed", zap.String("file", configFilePath), zap.Error(err))
		os.Exit(1)
	}

	if err := database.InitDB(); err != nil {
		os.Exit(1)
	}

	if err := database.InitTables(); err != nil {
		os.Exit(1)
	}

	if err := util.SetupUtils(); err != nil {
		os.Exit(1)
	}

	service.SetupAccessToken()
}

func run() {
	// setup otel tracing
	shutdown, err := util.SetupTracing()
	defer func() {
		zapx.Info("tracing reporter is shutting down...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := shutdown(ctx); err != nil {
			zapx.Error("tracing have been down down")
			return
		}
		zapx.Info("tracing reporter shut down successfully")
	}()

	if err != nil {
		zapx.Error("setup otel tracing failed", zap.Error(err))
		os.Exit(1)
	}

	if conf.SSOConf.Application.Mode == common.DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	router.InitRouter(r)

	addr := fmt.Sprintf("%s:%s",
		conf.SSOConf.Application.Host,
		conf.SSOConf.Application.Port)
	srv := http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  time.Second * time.Duration(conf.SSOConf.Application.ReadTimeout),
		WriteTimeout: time.Second * time.Duration(conf.SSOConf.Application.WriteTimeout),
	}

	go func() {
		zapx.Info("start http server", zap.String("host", addr))
		if err := srv.ListenAndServe(); err != nil {
			zapx.Error("http run error", zap.Error(err))
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		zapx.Error("shutdown http server failed", zap.Error(err))
		os.Exit(1)
	}
}
