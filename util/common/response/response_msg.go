package response

type Response struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(data interface{}) *Response {
	return &Response{
		Code: 200,
		Msg:  "success",
		Data: data,
	}
}

func Fail(msg string) *Response {
	return &Response{
		Code: -1,
		Msg:  msg,
	}
}

func FailCodeMsg(code int32, msg string) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
	}
}
