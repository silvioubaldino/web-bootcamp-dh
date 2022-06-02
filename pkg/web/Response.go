package web

type Response struct {
	Code  int    `json:"Code"`
	Data  any    `json:"Data,omitempty"`
	Error string `json:"Error,omitempty"`
}

func NewSuccesResp(code int, data any) Response {
	return Response{
		Code:  code,
		Data:  data,
		Error: "",
	}
}

func NewErrorResp(code int, err string) Response {
	return Response{
		Code:  code,
		Data:  nil,
		Error: err,
	}
}
