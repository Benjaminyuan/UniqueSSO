package session

import "time"

const(
	COOKIE_NAME = "SSO_COOKIE"
	SID = "sid"
	DEFAULT_TIMEOUT = 10000
)

// session类的抽象
type Session interface {
	Set(key,value interface{})
	Get(key interface{}) interface{}
	Delete(key interface{}) interface{}
	Save()
	SessionID()string
}
// 提供session的底层存储
type Provider interface {
	SessionInit(sid string)(Session,error)
	SessionRead(sid string)(Session,error)
	SessionUpdate(sid string ,session Session)error
	SessionDestroy(sid string)error
	SetMaxLifeTime(maxTime time.Duration)
	SessionGC()
}
