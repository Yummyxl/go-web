package response

type BaseResponseJson struct {
	Status int32       `json:"status"`
	Msg    interface{} `json:"msg"`
	Data   interface{} `json:"data"`
}

func BaseResponse(status int32, data interface{}, msg string) (json *BaseResponseJson) {
	json = &BaseResponseJson{Status: status, Data: data, Msg: msg}
	return
}
