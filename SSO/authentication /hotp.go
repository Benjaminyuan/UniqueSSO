package authentication

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"math"
	"strconv"
)

const (
	// KEY    = "HOTPkey"
	CODE_SIZE = 6
)

func GenerateHOTP(key string,counter int64) string {
	counterStr := strconv.FormatInt(counter,10)
	return Truncate(HMAC_SHA_1(key,counterStr))
}
func HMAC_SHA_1(key, input string) []byte {
	keyByte := []byte(key)
	h := hmac.New(sha1.New, keyByte)
	h.Write([]byte(input))
	return h.Sum(nil)
}
func Truncate(data []byte)string {
	offset := int(data[len(data)-1] & 0x0f)
	truncateBytes := data[offset : offset+4]
	num := binary.BigEndian.Uint32(truncateBytes)
	num = num % uint32(math.Pow10(CODE_SIZE))
	strNum := strconv.Itoa(int(num))
	for len(strNum)<6 {
		strNum = "0" + strNum
	}
	return strNum
}