package session

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
	"unique/jedi/util"
)

var (
	providers = make(map[string]Provider)
)

type RedisData struct {
	Sid        string `json:"sid"`
	AccessTime int64  `json:"access_time"`
	Name       string `json:"name"`
	College    string `json:"college"`
	UID 	   int64   `json:"uid"`
}
type MySession struct {
	sid          string
	timeAccessed int64
	c            map[interface{}]interface{}
}
type RedisProvider struct {
	r       *redis.Client
	maxTime time.Duration
	lock    sync.Mutex
	c       *gin.Context
}

// session的管理器
type Manager struct {
	cookieName  string
	lock        sync.Mutex
	provider    Provider
	maxLifeTime int64
}

func (s *MySession) Set(key, value interface{}) {
	s.c[key] = value
}
func (s *MySession) Get(key interface{}) interface{} {
	value, exist := s.c[key]
	if !exist {
		return nil
	}
	return value
}
func (s *MySession) Delete(key interface{}) interface{} {
	value := s.Get(key)
	delete(s.c, key)
	return value
}
func (s *MySession) Save() {

}

func (s *MySession) SessionID() string {
	return s.sid
}
func (s *MySession) String() string {
	r := RedisData{
		Sid:        s.sid,
		AccessTime: s.timeAccessed,
	}
	if s.c != nil {
		if key,ok := s.c[SessionUserUIDKey]; ok {
			r.UID = key.(int64)
		}
		if key,ok := s.c[SessionUserNameKey];ok {
			r.Name = key.(string)
		}
		if key,ok := s.c[SessionUserCollegeKey];ok {
			r.College = key.(string)
		}
	}
	data, err := json.Marshal(r)
	if err != nil {
		logrus.Errorf("fail to marshal,err: %v", err)
	}
	logrus.Debugf("data:%v", string(data))
	return string(data)
}
func (s *MySession) Deserialization(data []byte) error {
	r := &RedisData{}
	err := json.Unmarshal(data, r)
	if err != nil {
		logrus.Errorf("fail to deserialization:%v,err")
		return err
	}
	s.sid = r.Sid
	s.timeAccessed = r.AccessTime
	if s.c == nil {
		s.c = make(map[interface{}]interface{})
	}
	s.c[SessionUserUIDKey] = r.UID
	s.c[SessionUserNameKey] = r.Name
	s.c[SessionUserCollegeKey] = r.College
	return nil
}
func (p *RedisProvider) SessionInit(sid string) (Session, error) {
	s := &MySession{sid: sid, c: make(map[interface{}]interface{}), timeAccessed: time.Now().Unix()}
	s.Set(SID, sid)
	logrus.Debugf("session:%v", s.String())
	_, err := p.r.Set(sid, s.String(), p.maxTime).Result()
	if err != nil {
		logrus.Fatalf(" Session Init,fail to init session, err: %v", err)
		return nil, err
	}
	return s, nil
}
func (p *RedisProvider) SessionRead(sid string) (Session, error) {
	res, err := p.r.Get(sid).Result()
	if err == redis.Nil {
		logrus.Fatalf("sid not exist")
		return nil, errors.New("sid not exist ")
	} else if err != nil {
		return nil, err
	}
	s := &MySession{}
	err = s.Deserialization([]byte(res))
	if err != nil {
		return nil, err
	}
	return s, nil
}
func (p *RedisProvider) SessionUpdate(sid string, session Session) error {
	_, err := p.r.Get(sid).Result()
	if err == redis.Nil {
		logrus.Fatalf("sid not exist")
		return errors.New("sid not exist")
	}
	s, ok := session.(*MySession)
	logrus.Debugf("session:%+v", s.String())
	if !ok {
		logrus.Fatalf("fail to cast session")
		return errors.New("cast error ")
	}
	_, err = p.r.Set(sid, s.String(), p.maxTime).Result()
	if err != nil {
		logrus.Fatalf("fail to update session")
		return err
	}
	return nil
}
func (p *RedisProvider) SessionDestroy(sid string) error {
	res, err := p.r.Del(sid).Result()
	if err != nil {
		logrus.Fatalf("fail to del sid : %v,err: %v", err)
	}
	logrus.Debugf("del val : %v", res)
	return nil
}
func (p *RedisProvider) SetMaxLifeTime(maxTime time.Duration) {
	if maxTime > 0 {
		p.maxTime = maxTime
	} else {
		p.maxTime = DEFAULT_TIMEOUT
	}
}
func (p *RedisProvider) SessionGC() {

}

func (m *Manager) sessionID() string {
	return util.UUID()
}

func (m *Manager) SessionStart() (Session, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	id := m.sessionID()
	s, err := m.provider.SessionInit(id)
	if err != nil {
		logrus.Fatalf("fail to init session: %v", err)
		return nil, err
	}
	return s, nil
}
func (m *Manager) SessionUpdate(sid string, session Session) error {
	if err := m.provider.SessionUpdate(sid, session); err != nil {
		logrus.Fatalf("fail to update session, err: %v", err)
		return err
	}
	return nil
}
func (m *Manager) SessionRead(sid string) (Session, error) {
	s, err := m.provider.SessionRead(sid)
	if err != nil {
		logrus.Fatalf("fail to find session, err: %v", err)
		return nil, err
	}
	return s, nil
}
func (m *Manager) SessionDestroy(sid string) error {
	if err := m.provider.SessionDestroy(sid); err != nil {
		logrus.Fatalf("fail to destroy session,err: %v", err)
		return err
	}
	return nil
}
func NewManager(providerName, cookieName string, maxLifeTime int64) (*Manager, error) {
	provider, ok := providers[providerName]
	if !ok {
		logrus.Fatal("provider not found")
		return nil, errors.New("provider not found")
	}
	provider.SetMaxLifeTime(time.Duration(maxLifeTime) * time.Second)
	return &Manager{
		cookieName:  cookieName,
		lock:        sync.Mutex{},
		provider:    provider,
		maxLifeTime: 0,
	}, nil
}
func RegisterProvider(provideName string, provider Provider) {
	if _, ok := providers[provideName]; ok {
		logrus.Infof("update exist provider")
	}
	providers[provideName] = provider
}
