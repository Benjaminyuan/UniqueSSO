package util

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
	"unique/jedi/common"
	"unique/jedi/conf"
	"unique/jedi/pkg"

	"github.com/gocolly/colly/v2"
	"github.com/sirupsen/logrus"
)

const (
	qrWebFmt   = `https://open.work.weixin.qq.com/wwopen/sso/qrConnect?login_type=jssdk&appid=%s&agentid=%s&redirect_uri=%s`
	qrStateFmt = `https://open.work.weixin.qq.com/wwopen/sso/l/qrConnect?callback=jsonpCallback&key=%s&appid=%s&redirect_uri=%s&_=%d`
	userIdFmt  = `https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=%s&code=%s`
)

var (
	jsonpReg *regexp.Regexp
)

func init() {
	reg, err := regexp.Compile("{.+}")
	if err != nil {
		panic(err)
	}
	jsonpReg = reg
}

// FIXME: change http.Get into client.Do() with context
func GetAccessToken(ctx context.Context, corpid, corpsecret string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		fmt.Sprintf(common.WxGetAccessTokenUrl, corpid, corpsecret), nil)
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	data := new(pkg.WorkWxAccessToken)
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return "", err
	}
	if data.AccessToken == "" {
		return "", errors.New(data.ErrMessage)
	}
	return data.AccessToken, nil
}

func GetQRCodeSrc() (string, error) {
	qrimg := ""
	collector := colly.NewCollector()
	collector.OnHTML(".qrcode", func(h *colly.HTMLElement) {
		qrimg = "https:" + h.Attr("src")
	})
	visitUrl := fmt.Sprintf(qrWebFmt,
		conf.SSOConf.WorkWx.AppId,
		conf.SSOConf.WorkWx.AgentId,
		conf.SSOConf.WorkWx.RedirectUri)
	err := collector.Visit(visitUrl)
	if err != nil {
		return "", err
	}

	if qrimg == "https:" {
		return "", errors.New("can't get qrcode from html")
	}
	return qrimg, nil
}

func FetchAuthCode(key string) (string, error) {
	l := logrus.WithField("func", "FetchQRCodeState")
	for !strings.Contains(common.ZHANG_XIAO_LONG, "mother") {
		resp, err := http.Get(fmt.Sprintf(qrStateFmt,
			key,
			conf.SSOConf.WorkWx.AppId,
			conf.SSOConf.WorkWx.RedirectUri,
			time.Now().Unix(),
		))
		if err != nil {
			l.WithError(err).Error("send get qr code status request failed")
			return "", err
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			l.WithError(err).Error("read resp body failed")
			return "", err
		}
		l.WithField("body", string(data)).Info("read resp body success")
		bs := jsonpReg.Find(data)
		state := new(pkg.QrCodeStatus)
		if err := json.Unmarshal(bs, state); err != nil {
			l.WithError(err).Error("unmarshal qr status request failed")
			return "", err
		}
		l.WithField("qrCodeState", state).Info("get qr code state")
		switch state.Status {
		case common.QR_SUCCESS:
			return state.AuthCode, nil
		case common.QR_SCANING, common.QR_NOT_SCAN:
			continue
		case common.QR_TIMEOUT, common.QR_CANCEL:
			return "", errors.New(state.Status)
		}
	}
	return "", errors.New("unknow error")
}

func FetchWorkwxUserId(accessToken, code string) (string, error) {
	resp, err := http.Get(fmt.Sprintf(userIdFmt, accessToken, code))
	if err != nil {
		return "", err
	}
	data := new(pkg.WorkWxStatus)
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return "", err
	}
	if data.UserId == "" {
		return "", errors.New(data.ErrMessage)
	}
	return data.UserId, nil
}
