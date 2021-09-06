package pkg

type QrCodeStatus struct {
	Status   string `json:"status"`
	AuthCode string `json:"auth_code"`
}

type LoginUser struct {
	Phone     string `json:"phone,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	Code      string `json:"code,omitempty"`
	QrcodeSrc string `json:"qrcode_src"`
}

type WorkWxStatus struct {
	ErrCode    int    `json:"errcode"`
	ErrMessage string `json:"errmsg"`
	UserId     string `json:"userId"`
}

type WorkWxAccessToken struct {
	ErrCode     int    `json:"errcode"`
	ErrMessage  string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
