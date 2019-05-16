package defs

import (
	"encoding/json"
	"net/http"
)

type (
	// HTTPResponse 响应
	HTTPResponse struct {
		Code    string      `json:"code"`
		Msg     string      `json:"msg"`
		Data    interface{} `json:"data"`
		Success bool        `json:"success"`
	}
	// ErrorResponse 错误响应 包含Http状态码
	ErrorResponse struct {
		Code int
		Res  HTTPResponse
	}
	// MSResponse MS RESPONSE
	MSResponse struct {
		TokenType    string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in"`
		Scope        string `json:"scope"`
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
)

var (
	// ErrorNotGetInfo 没有获得认证信息
	ErrorNotGetInfo = ErrorResponse{Code: 400, Res: HTTPResponse{Code: "001", Msg: "can`t get info", Success: false}}
	// ErrorJSONParse JSON解析失败
	ErrorJSONParse = ErrorResponse{Code: 500, Res: HTTPResponse{Code: "002", Msg: "internal server error", Success: false}}
	// ErrorInternalServer 内部服务器错误
	ErrorInternalServer = ErrorResponse{Code: 500, Res: HTTPResponse{Code: "003", Msg: "internal server error", Success: false}}
	// ErrorCanNotGetAccessToken 无法获取AccessToken
	ErrorCanNotGetAccessToken = ErrorResponse{Code: 500, Res: HTTPResponse{Code: "004", Msg: "cant not get access_token", Success: false}}
	// ErrorDownload 下载失败
	ErrorDownload = ErrorResponse{Code: 500, Res: HTTPResponse{Code: "005", Msg: "cant not download", Success: false}}
)

// SendSuccessResponse 发送成功的消息
func SendSuccessResponse(w http.ResponseWriter, code, msg string, state int, data interface{}) {
	res := HTTPResponse{
		Code:    code,
		Msg:     msg,
		Data:    data,
		Success: true,
	}
	resByte, _ := json.Marshal(res)
	w.WriteHeader(state)
	w.Write(resByte)
}

// SendErrorResponse 发送错误
func SendErrorResponse(w http.ResponseWriter, res ErrorResponse) {
	resByte, _ := json.Marshal(res.Res)
	w.WriteHeader(res.Code)
	w.Write(resByte)
}
