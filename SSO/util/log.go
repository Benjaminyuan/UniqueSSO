package util

import "github.com/sirupsen/logrus"

func InitLogrus() (err error) {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		DisableColors: true,
	})

	logrus.SetLevel(logrus.InfoLevel)

	return nil
}
