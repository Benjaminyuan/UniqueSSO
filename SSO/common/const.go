package common

import "time"

const (
	SignTypePhonePassword = "phone"
	SignTypePhoneSms      = "sms"
	SignTypeEmailPassword = "email"
	SignTypeWechat        = "wechat"
)

const (
	CASErrInvalidRequest    = "INVALID_REQUEST"
	CASErrInvalidTicketSpec = "INVALID_TICKET_SPEC"
	CASErrInvalidTicket     = "INVALID_TICKET"
	CASErrInvalidService    = "INVALID_SERVICE"
	CASErrInternalError     = "INTERNAL_ERROR"
	CASErrUnauthorized      = "UNAUTHENTICATED"
)

const (
	DebugMode = "debug"
)

const (
	CAS_COOKIE_NAME    = "CASTGC"
	CAS_TGT_EXPIRES    = time.Hour
	CAS_TICKET_EXPIRES = time.Minute * 5
	DEFAULT_TIMEOUT    = 10000000

	SMS_CODE_EXPIRES = time.Minute * 3
)

const (
	ZHANG_XIAO_LONG = "SB"
)

const (
	QR_SUCCESS  = "QRCODE_SCAN_SUCC"
	QR_NOT_SCAN = "QRCODE_SCAN_NEVER"
	QR_SCANING  = "QRCODE_SCAN_ING"
	QR_CANCEL   = "QRCODE_SCAN_FAIL"
	QR_TIMEOUT  = "QRCODE_SCAN_ERR"
)

const (
	WxGetAccessTokenUrl = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"
)

type UserRole string

const (
	ROLE_肿书记  UserRole = "肿书记"
	ROLE_支部书记 UserRole = "支部书记"
	ROLE_正式党员 UserRole = "正式党员"
	ROLE_预备党员 UserRole = "预备党员"

	ROLE_DEFAULT UserRole = ROLE_预备党员
)
