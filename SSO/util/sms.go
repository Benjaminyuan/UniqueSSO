package util

import (
	"context"
	"time"
	uniquec "unique/jedi/common"
	"unique/jedi/database"
	"unique/jedi/pb/sms"
	"unique/jedi/pkg"
)

func GenerateSMSCode(ctx context.Context, phone string) (string, error) {
	code := NewSMSCode()
	err := database.RedisClient.Set(ctx, phone, code, uniquec.SMS_CODE_EXPIRES).Err()
	if err != nil {
		return "", err
	}
	return code, nil
}

func GetSMSCodeByPhone(ctx context.Context, phone string) (code string, err error) {
	value := database.RedisClient.GetDel(ctx, phone)
	if err = value.Err(); err != nil {
		return "", value.Err()
	}
	if err = value.Scan(&code); err != nil {
		return
	}
	return
}

// TODO: send sms by using Open-Platform
func SendSMS(ctx context.Context, phone string, code string, expire time.Duration) (*[]pkg.FailedSmsStatus, error) {
	resp, err := OpenClient.PushSMS(ctx, &sms.PushSMSRequest{})
	if err != nil {
		return nil, err
	}
	data := make([]pkg.FailedSmsStatus, 0, len(resp.SMSStatus))
	for i := range resp.SMSStatus {
		if resp.SMSStatus[i].ErrCode != "" {
			data = append(data, pkg.FailedSmsStatus{
				Phone:   resp.SMSStatus[i].PhoneNumber,
				Message: resp.SMSStatus[i].Message,
			})
		}
	}
	return &data, nil
}
