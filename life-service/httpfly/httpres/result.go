package httpres

import "encoding/json"

type (
	// HTTPResult http的响应结构体 用于同意返回结果
	HTTPResult struct {
		Code    int         `json:"code"`
		Msg     string      `json:"msg"`
		Data    interface{} `json:"data"`
		Success bool        `json:"success"`
	}
)

// Create 创建一个响应
func Create(code int, msg string, data interface{}, success bool) *HTTPResult {
	return &HTTPResult{code, msg, data, success}
}

// Success 普通的成功返回
func Success(data interface{}) *HTTPResult {
	return &HTTPResult{
		Code:    200,
		Msg:     "成功",
		Data:    data,
		Success: true,
	}
}

// Fail 普通的失败返回结果
func Fail() *HTTPResult {
	return &HTTPResult{
		Code:    400,
		Msg:     "失败",
		Data:    nil,
		Success: false,
	}
}

// NotLogin 未登录
func NotLogin() *HTTPResult {
	return &HTTPResult{
		Code:    403,
		Msg:     "请先登录",
		Data:    nil,
		Success: false,
	}
}

// Bytes Bytes
func (r *HTTPResult) Bytes() []byte {
	bs, _ := json.Marshal(r)
	return bs
}
