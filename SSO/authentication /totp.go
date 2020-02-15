package authentication 
import(
	"time"
)
const (
	BASE_TIME = 1257894000
	INTERVAL int64 = 300
)
func GenerateTOTP(key string,time int64) string {
	var counter int64 = (time.Now().Unix()- BASE_TIME)/ INTERVAL
	return GenerateHOTP(key,counter)
}