package rest

const CodeForSuccess = 0

type R struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}
