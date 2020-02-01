package conf

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
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
	//MysqlPassword string `yaml:"MysqlPassword"`
}
type RPCConf struct {
	Addr string `yaml:"Addr"`
}
type Conf struct {
	OriginAllowedList []string `yaml:"OriginAllowedList"`
	MysqlConf MysqlConf `yaml:"MysqlConf"`
	RedisConf RedisConf `yaml:"RedisConf"`
	RPCConf RPCConf `yaml:"RPCConf"`
	MD5Sum string `yaml:"MD5Sum"`
}

var (
	SSOConf = &Conf{}
	RedisClient *redis.Client
	DB *gorm.DB
)
func ExampleNewClient(addr string,password string)(*redis.Client,error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       0,  // use default DB
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client,err
}
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
func InitDB()error{
	var err error
	RedisClient, err = ExampleNewClient(SSOConf.RedisConf.RedisAddr,SSOConf.RedisConf.RedisPassword)
	if err != nil {
		log.Fatalf("fail to init redis client, err:%v",err)
		return err
	}
	db,err := gorm.Open("mysql",SSOConf.MysqlConf.MysqlAddr)
	if err != nil{
		log.Fatalf("fail to init mysql client, err: %v",err)
		return err
	}
	DB = db
	return nil
}