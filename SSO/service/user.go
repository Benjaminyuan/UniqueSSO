package service

import (
	"context"
	"errors"
	"net/url"
	"unique/jedi/common"
	"unique/jedi/conf"
	"unique/jedi/model"
	"unique/jedi/pkg"
	"unique/jedi/util"

	"github.com/sirupsen/logrus"
)

func VerifyUser(ctx context.Context, login *pkg.LoginUser, signType string) (*model.User, error) {
	switch signType {
	case common.SignTypeEmailPassword:
		return VerifyUserByEmail(login.Email, login.Password)
	case common.SignTypePhonePassword:
		return VerifyUserByPhone(login.Phone, login.Password)
	case common.SignTypePhoneSms:
		return VerifyUserBySMS(ctx, login.Phone, login.Code)
	case common.SignTypeWechat:
		return VerifyUserByQrcode(login.QrcodeSrc)
	default:
		return nil, errors.New("Invalid sign type")
	}
}

func GetUserById(uid string) (*model.User, error) {
	user := new(model.User)
	err := conf.DB.Table(user.TableName()).Where("uid = ?", uid).Scan(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func VerifyUserByEmail(email, password string) (*model.User, error) {
	lg := logrus.WithField("type", "email")
	user := new(model.User)
	err := conf.DB.Table(user.TableName()).Where("email = ?", email).Scan(user).Error
	if err != nil {
		lg.WithError(err).Info("get user by email error")
		return nil, err
	}
	if err := util.ValidatePassword(password, user.Password); err != nil {
		lg.WithError(err).Info("validate password error")
		return nil, err
	}
	return user, nil
}

func VerifyUserByPhone(phone, password string) (*model.User, error) {
	user := new(model.User)
	err := conf.DB.Table(user.TableName()).Where("phone = ?", phone).Scan(user).Error
	if err != nil {
		return nil, err
	}
	if err := util.ValidatePassword(password, user.Password); err != nil {
		return nil, err
	}
	return user, nil
}

func VerifyUserBySMS(ctx context.Context, phone, sms string) (*model.User, error) {
	user := new(model.User)
	code, err := util.GetSMSCodeByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}
	if code != sms {
		return nil, errors.New("sms code is wrong")
	}
	err = conf.DB.Table((user.TableName())).Where("phone = ?", phone).Scan(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func VerifyUserByQrcode(qrcode string) (*model.User, error) {
	src, err := url.Parse(qrcode)
	if err != nil {
		return nil, err
	}
	code, err := util.FetchAuthCode(src.Query().Get("key"))
	if err != nil {
		return nil, err
	}

	conf.SSOConf.WorkWx.AccessToken.RWLock.RLock()
	token := conf.SSOConf.WorkWx.AccessToken.Token
	conf.SSOConf.WorkWx.AccessToken.RWLock.RUnlock()
	userid, err := util.FetchWorkwxUserId(token, code)
	if err != nil {
		return nil, err
	}

	user := new(model.User)
	err = conf.DB.Table(user.TableName()).Where("workwx_user_id = ?", userid).Scan(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
