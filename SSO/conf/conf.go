package conf

import (
	"net/url"
	"regexp"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"github.com/xylonx/zapx"
	"go.uber.org/zap"
)

type Conf struct {
	Application  ApplicationConf  `mapstructure:"application"`
	Database     DatabaseConf     `mapstructure:"database"`
	Redis        RedisConf        `mapstructure:"redis"`
	Sms          []SMSOptions     `mapstructure:"sms"`
	WorkWx       WorkWxConf       `mapstructure:"work_wx"`
	OpenPlatform OpenPlatformConf `mapstructure:"openplat_form"`
	APM          APMConf          `mapstructure:"apm"`
}
type ApplicationConf struct {
	Host            string           `mapstructure:"host"`
	Port            string           `mapstructure:"port"`
	Name            string           `mapstructure:"name"`
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

type SMSOptions struct {
	Name       string `mapstructure:"name" validator:"oneof='verificationCode'"`
	TemplateId string `mapstructure:"template_id"`
	SignName   string `mapstructure:"sign_name"`
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

type OpenPlatformConf struct {
	GrpcAddr       string `mapstructure:"grpc_addr"`
	GrpcCert       string `mapstructure:"grpc_cert"`
	GrpcServerName string `mapstructure:"grpc_server_name"`
}

type APMConf struct {
	ReporterBackground string `mapstructure:"reporter_backend"`
}

var (
	SSOConf = &Conf{}
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

	validate := validator.New()
	if err := validate.Struct(SSOConf); err != nil {
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
		zapx.Info("run mode", zap.String("mode", SSOConf.Application.Mode))
	}

	SSOConf.WorkWx.RedirectUri = url.PathEscape(SSOConf.WorkWx.RedirectUri)

	return nil
}
