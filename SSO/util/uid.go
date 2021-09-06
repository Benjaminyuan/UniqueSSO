package util

import (
	"crypto/md5"
	//"crypto/rand"
	//"encoding/base64"
	"encoding/hex"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetMD5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func UUID() string {
	uid := uuid.New()
	return uid.String()
}

func GetUUID() string {
	return primitive.NewObjectID().Hex()
}
