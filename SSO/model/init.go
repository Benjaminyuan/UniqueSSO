package model

import (
	"unique/jedi/conf"

	"github.com/sirupsen/logrus"
)

func InitTables() (err error) {
	err = conf.DB.AutoMigrate(&User{})
	if err != nil {
		logrus.WithField("table", (&User{}).TableName()).Error("create table failed")
		return err
	}
	return nil
}
