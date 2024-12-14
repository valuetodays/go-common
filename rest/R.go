package rest

const CodeForSuccess = 0

type R struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func OfError(code int32, msg string) *R {
	return &R{
		Msg:  msg,
		Code: code,
	}
}

func OfSuccess(data any) *R {
	return &R{
		Code: CodeForSuccess,
		Data: data,
	}
}
