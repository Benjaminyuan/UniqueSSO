package conf

import (
	"context"
	"net/url"
	"regexp"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Conf struct {
	Application ApplicationConf `mapstructure:"application"`
	Database    DatabaseConf    `mapstructure:"database"`
	Redis       RedisConf       `mapstructure:"redis"`
	Sms         SmsConf         `mapstructure:"sms"`
	WorkWx      WorkWxConf      `mapstructure:"work_wx"`
}
type ApplicationConf struct {
	Host            string           `mapstructure:"host"`
	Port            string           `mapstructure:"port"`
	Mode            string           `mapstructure:"mode"`
	ReadTimeout     int              `mapstructure:"read_timeout"`
	WriteTimeout    int              `mapstructure:"write_timeout"`
	AllowService    []string         `mapstructure:"allow_service"`
	AllowServiceReg []*regexp.Regexp `mapstructure:"-"`
}

type DatabaseConf struct {
	PostgresDSN string `mapstructure:"postgres_dsn"`
}

type RedisConf struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type SmsConf struct {
	SecretId  string `mapstructure:"secret_id"`
	SecretKey string `mapstructure:"secret_key"`
	AppId     string `mapstructure:"app_id"`
}

type WorkWxConf struct {
	AppId       string `mapstructure:"app_id"`
	AgentId     string `mapstructure:"agent_id"`
	RedirectUri string `mapstructure:"redirect_uri"`
	CorpId      string `mapstructure:"corpid"`
	CorpSecret  string `mapstructure:"corpsecret"`
	AccessToken struct {
		RWLock sync.RWMutex
		Token  string
	} `mapstructure:"-"`
}

var (
	SSOConf     = &Conf{}
	DB          *gorm.DB
	RedisClient *redis.Client
)

func InitConf(confFilepath string) error {
	viper.SetConfigFile(confFilepath)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(SSOConf)
	if err != nil {
		return err
	}

	SSOConf.Application.AllowServiceReg = make([]*regexp.Regexp, len(SSOConf.Application.AllowService))
	for i, service := range SSOConf.Application.AllowService {
		reg, err := regexp.Compile(service)
		if err != nil {
			return err
		}
		SSOConf.Application.AllowServiceReg[i] = reg
	}

	if SSOConf.Application.Mode == "debug" {
		logrus.WithField("config", SSOConf).Debug("load config")
	}

	SSOConf.WorkWx.RedirectUri = url.PathEscape(SSOConf.WorkWx.RedirectUri)

	return nil
}

func InitDB(ctx context.Context) (err error) {
	// connect postgres
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: SSOConf.Database.PostgresDSN,
	}))
	if err != nil {
		logrus.WithError(err).Error("open gorm error")
		return err
	}
	logrus.Info("conect to postgres success")
	DB = db

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	RedisClient, err = initRedis(ctx)
	if err != nil {
		return err
	}

	return nil
}

func initRedis(ctx context.Context) (*redis.Client, error) {
	// init redis
	rclient := redis.NewClient(&redis.Options{
		Addr:     SSOConf.Redis.Addr,
		Password: SSOConf.Redis.Password,
		DB:       SSOConf.Redis.DB,
	})
	pong, err := rclient.Ping(ctx).Result()
	if err != nil {
		logrus.WithError(err).Error("ping redis error")
		return nil, err
	}
	logrus.WithField("result", pong).Info("ping redis success")
	return rclient, nil
}
