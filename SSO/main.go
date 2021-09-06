package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
	"unique/jedi/common"
	"unique/jedi/conf"
	"unique/jedi/middleware"
	"unique/jedi/model"
	"unique/jedi/router"
	"unique/jedi/service"
	"unique/jedi/util"

	"github.com/SkyAPM/go2sky"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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

var Tracer *go2sky.Tracer

func init() {
	rootCmd.Flags().StringVarP(&configFilePath, "config", "c", "./conf/conf.yaml", "path to config file")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatalf("[]Failed to execute rootCmd")
	}
}

func setup() {
	if err := conf.InitConf(configFilePath); err != nil {
		os.Exit(1)
	}

	if err := util.InitLogrus(); err != nil {
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := conf.InitDB(ctx); err != nil {
		os.Exit(1)
	}

	if err := model.InitTables(); err != nil {
		os.Exit(1)
	}

	service.SetupAccessToken()

	// if no apm access, should not panic
	// if err := util.SetupAPM(); err != nil {
	// }
}

func run() {
	r := gin.Default()

	if conf.SSOConf.Application.Mode == common.DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// TODO: intergrating the APM System
	// r.Use(middleware.APMMiddleware())
	r.Use(middleware.Cors())

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
		logrus.WithFields(logrus.Fields{
			"Host": addr,
		}).Info("start HTTP server")
		if err := srv.ListenAndServe(); err != nil {
			logrus.WithError(err).Error("http server run error")
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logrus.WithError(err).Error("shutdown service error")
		os.Exit(1)
	}
}
