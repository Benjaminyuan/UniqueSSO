package util

import (
	"crypto/md5"
	"encoding/hex"
)

func EncryptPassword(password string)string{
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(password))
	//return string(md5Ctx.Sum([]byte(conf.SSOConf.MD5Sum)))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}