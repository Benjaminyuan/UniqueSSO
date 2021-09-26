package pkg

import (
	"unique/jedi/common"
	"unique/jedi/database"
)

type CommonResponse struct {
	ServiceResponse serviceRespnse `json:"serviceResponse,omitempty"`
}

type serviceRespnse struct {
	AuthenticationSuccess *authenticationSuccess `json:"authenticationSuccess,omitempty"`
	AuthenticationFailure *authenticationFailure `json:"authenticationFailure,omitempty"`
}

type authenticationSuccess struct {
	RedirectService *string        `json:"redirectService,omitempty"`
	UserId          *string        `json:"user,omitempty"`
	QrcodeSrc       *string        `json:"qrcodeSrc,omitempty"`
	Attributes      *database.User `json:"attributes,omitempty"`
}

type authenticationFailure struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type FailedSmsStatus struct {
	Phone   string
	Message string
}

func InvalidRequest(err error) *CommonResponse {
	return &CommonResponse{
		ServiceResponse: serviceRespnse{
			AuthenticationSuccess: nil,
			AuthenticationFailure: &authenticationFailure{Code: common.CASErrInvalidRequest, Description: err.Error()},
		},
	}
}

func InvalidTicketSpec(err error) *CommonResponse {
	return &CommonResponse{
		ServiceResponse: serviceRespnse{
			AuthenticationSuccess: nil,
			AuthenticationFailure: &authenticationFailure{Code: common.CASErrInvalidTicketSpec, Description: err.Error()}},
	}
}

func InvalidTicket(err error) *CommonResponse {
	return &CommonResponse{
		ServiceResponse: serviceRespnse{
			AuthenticationSuccess: nil,
			AuthenticationFailure: &authenticationFailure{Code: common.CASErrInvalidTicket, Description: err.Error()},
		},
	}
}

func InvalidService(err error) *CommonResponse {
	return &CommonResponse{
		ServiceResponse: serviceRespnse{
			AuthenticationSuccess: nil,
			AuthenticationFailure: &authenticationFailure{Code: common.CASErrInvalidService, Description: err.Error()},
		},
	}
}

func InternalError(err error) *CommonResponse {
	return &CommonResponse{
		ServiceResponse: serviceRespnse{
			AuthenticationSuccess: nil,
			AuthenticationFailure: &authenticationFailure{Code: common.CASErrInternalError, Description: err.Error()},
		},
	}
}

func AuthSuccess(userInfo *database.User) *CommonResponse {
	return &CommonResponse{
		ServiceResponse: serviceRespnse{
			AuthenticationSuccess: &authenticationSuccess{
				UserId:     &userInfo.UID,
				Attributes: userInfo,
			},
			AuthenticationFailure: nil,
		},
	}
}

func RedirectSuccess(service string) *CommonResponse {
	return &CommonResponse{
		ServiceResponse: serviceRespnse{
			AuthenticationSuccess: &authenticationSuccess{
				RedirectService: &service,
			},
			AuthenticationFailure: nil,
		},
	}
}

func QrcodeSuccess(imgsrc string) *CommonResponse {
	return &CommonResponse{
		ServiceResponse: serviceRespnse{
			AuthenticationSuccess: &authenticationSuccess{
				QrcodeSrc: &imgsrc,
			},
			AuthenticationFailure: nil,
		},
	}
}
