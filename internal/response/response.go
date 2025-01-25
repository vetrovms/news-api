package response

// Response Структура відповіді.
type Response struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
	Data  any    `json:"data"`
}

// NewResponse Повертає структуру відповіді.
func NewResponse(code int, error string, data any) *Response {
	return &Response{
		Code:  code,
		Error: error,
		Data:  data,
	}
}
