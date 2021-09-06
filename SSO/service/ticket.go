package service

import (
	"context"
	"errors"
	"time"
	"unique/jedi/conf"
)

func StoreValue(ctx context.Context, key, value string, expires time.Duration) error {
	return conf.RedisClient.Set(ctx, key, value, expires).Err()
}

func GetValue(ctx context.Context, key string) (string, error) {
	return conf.RedisClient.Get(ctx, key).Result()
}

func GetDelValue(ctx context.Context, key string) (string, error) {
	return conf.RedisClient.GetDel(ctx, key).Result()
}

func VerifyService(service string) error {
	for i := range conf.SSOConf.Application.AllowServiceReg {
		if conf.SSOConf.Application.AllowServiceReg[i].MatchString(service) {
			return nil
		}
	}
	return errors.New("service not allow")
}
