package service

import (
	"context"
	"time"
	"unique/jedi/conf"
	"unique/jedi/util"

	"github.com/sirupsen/logrus"
)

// update accessToken periodically
func SetupAccessToken() {
	go func() {
		for {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			newToken, err := util.GetAccessToken(ctx, conf.SSOConf.WorkWx.CorpId, conf.SSOConf.WorkWx.CorpSecret)
			cancel()
			if err != nil {
				logrus.WithError(err).Error("get new access token failed")
				time.Sleep(time.Minute * 5)
				continue
			}
			conf.SSOConf.WorkWx.AccessToken.RWLock.Lock()
			conf.SSOConf.WorkWx.AccessToken.Token = newToken
			conf.SSOConf.WorkWx.AccessToken.RWLock.Unlock()
			logrus.WithField("access token", newToken).Info("get access token succeeded")
			time.Sleep(time.Hour + time.Minute*30)
		}
	}()
}
