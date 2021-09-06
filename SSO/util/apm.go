package util

import (
	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
	"github.com/sirupsen/logrus"
)

var (
	SkyReporter go2sky.Reporter
	Tracer      *go2sky.Tracer
)

func SetupAPM() (err error) {
	// set up skywalking apm
	SkyReporter, err = reporter.NewGRPCReporter("192.168.0.230:11800")
	if err != nil {
		logrus.WithError(err).Error("init skywalking reporter failed")
		return err
	}
	logrus.Info("init skywalking reporter success")

	Tracer, err = go2sky.NewTracer("UniqueSSO", go2sky.WithReporter(SkyReporter))
	if err != nil {
		logrus.WithError(err).Error("init skywalking tracer failed")
		return err
	}

	return nil
}
