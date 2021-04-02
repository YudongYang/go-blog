package controllers

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var (
	SUCCESS_CODE = 200
	FAIL_CODE    = 500
)

func (resp *Response) Success(data interface{}) {
	resp.Code = SUCCESS_CODE
	resp.Data = data
}

func (resp *Response) Fail(msg string) {
	resp.Code = FAIL_CODE
	resp.Msg = msg
}
