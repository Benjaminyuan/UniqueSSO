package util

import (
	"context"
	"strconv"
	"time"
	uniquec "unique/jedi/common"
	"unique/jedi/conf"
	"unique/jedi/pkg"

	"github.com/sirupsen/logrus"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

func GenerateSMSCode(ctx context.Context, phone string) (string, error) {
	code := NewSMSCode()
	err := conf.RedisClient.Set(ctx, phone, code, uniquec.SMS_CODE_EXPIRES).Err()
	if err != nil {
		return "", err
	}
	return code, nil
}

func GetSMSCodeByPhone(ctx context.Context, phone string) (code string, err error) {
	value := conf.RedisClient.GetDel(ctx, phone)
	if err = value.Err(); err != nil {
		return "", value.Err()
	}
	if err = value.Scan(&code); err != nil {
		return
	}
	return
}

// TODO: send sms by using Open-Platform
func SendSMS(phone string, code string, expire time.Duration) (*[]pkg.FailedSmsStatus, error) {
	credentials := common.NewCredential(conf.SSOConf.Sms.SecretId, conf.SSOConf.Sms.SecretKey)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqTimeout = 5 // second

	smsClient, err := sms.NewClient(credentials, "ap-guangzhou", cpf)
	if err != nil {
		return nil, err
	}

	req := sms.NewSendSmsRequest()
	req.SmsSdkAppId = common.StringPtr(conf.SSOConf.Sms.AppId)
	req.SignName = common.StringPtr("UniqueStudio")

	req.TemplateId = common.StringPtr("1099470")
	req.TemplateParamSet = common.StringPtrs([]string{code, strconv.FormatInt(int64(expire)/int64(time.Minute), 10)})

	req.PhoneNumberSet = common.StringPtrs([]string{phone})
	resp, err := smsClient.SendSms(req)
	if err != nil {
		return nil, err
	}

	failed := make([]pkg.FailedSmsStatus, 0, len(phone))
	for i := range resp.Response.SendStatusSet {
		if resp.Response.SendStatusSet[i].Code != nil &&
			*resp.Response.SendStatusSet[i].Code != "Ok" {
			f := pkg.FailedSmsStatus{
				Phone:   *(resp.Response.SendStatusSet[i].PhoneNumber),
				Message: *(resp.Response.SendStatusSet[i].Message),
			}
			logrus.WithFields(logrus.Fields{
				"phone":   f.Phone,
				"message": f.Message,
			}).Error("send sms failed")
			failed = append(failed, f)
		}
	}
	return &failed, nil
}
