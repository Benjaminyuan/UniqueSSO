package conf

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type RedisConf struct {
	RedisAddr     string `yaml:"RedisAddr"`
	RedisPassword string `yaml:"RedisPassword"`
}
type MysqlConf struct {
	MysqlAddr     string `yaml:"MysqlAddr"`
	MysqlPassword string `yaml:"MysqlPassword"`
}
type Conf struct {
	OriginAllowedList []string `yaml:"OriginAllowedList"`
	MysqlConf MysqlConf `yaml:"MysqlConf"`
	RedisConf RedisConf `yaml:"RedisConf"`
}
var (
	SSOConf = &Conf{}
)
func InitConf() error {
	fileName,_ := filepath.Abs("./conf/conf.yaml")
	yamlData, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Fail to load yaml file ,err: %v", err)
		return err
	}
	err = yaml.Unmarshal(yamlData,SSOConf)
	if err != nil{
		log.Fatalf("Fail to unmarshal yaml data,err:%v",err)
		return err
	}
	return nil
}
