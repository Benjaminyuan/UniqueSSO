package common


type CommonResponse struct {
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
var (
	SuccessResponse   = &CommonResponse{Msg: "OK", Code: 0, Data: ""}
	ParameterErrorResponse = &CommonResponse{Msg: "Parameter invalid", Code: 1, Data: ""}
	VerifyFailedResponse   = &CommonResponse{Msg: "Verify failed", Code: 2, Data: ""}
	ServerErrorResponse    = &CommonResponse{Msg: "Server Error", Code: 4, Data: ""}
)
func NewErrorResponse(msg string)*CommonResponse{
	return &CommonResponse{
		Msg:  msg,
		Code: 1,
		Data:"",
	}
}