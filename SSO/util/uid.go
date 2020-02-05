package util

import (
	"crypto/md5"
	//"crypto/rand"
	//"encoding/base64"
	"encoding/hex"
	"github.com/google/uuid"
)

func GetMD5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
func UUID() string{
	uid := uuid.New()
	return 	uid.String()
	//b := make([]byte,48)
	//if _,err := rand.Read(b);err != nil{
	//	return ""
	//}
	//return GetMD5String(base64.URLEncoding.EncodeToString(b))
}