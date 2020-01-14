package session

import (
	"github.com/sirupsen/logrus"
	"sync"
	"unique/jedi/conf"
)

var(
	GlobalSessionManager *Manager
)
func InitSession() error{
	var err error
	p := &RedisProvider{
		r:   conf.RedisClient,
		maxTime: 0,
		lock:    sync.Mutex{},
	}
	RegisterProvider("redis",p)
	GlobalSessionManager,err  = NewManager("redis",SID,DEFAULT_TIMEOUT)
	if err != nil {
		logrus.Fatalf("fail to new manager")
		return nil
	}
	return nil
}