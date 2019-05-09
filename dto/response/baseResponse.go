package response

type BaseResponseJson struct {
	Code int32       `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func BaseResponse(code int32, data interface{}, msg string) (json *BaseResponseJson) {
	json = &BaseResponseJson{Code: code, Data: data, Msg: msg}
	return
}

func Success(data interface{}) (json *BaseResponseJson) {
	json = &BaseResponseJson{Code: 0, Data: data, Msg: "成功"}
	return
}
