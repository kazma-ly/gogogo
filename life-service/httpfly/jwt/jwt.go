package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"errors"
	"hash"
	"net/http"
	"strings"
	"time"
)

type (
	// JWT JSON WEB TOKEN
	JWT struct {
		Payload   Payload // 载荷
		Header    Header  // 头部
		Sign      string  // 签名
		secKey    string  // 配和加密的字符串
		jwtString string  // jwt
	}

	// Payload JWT的Payload
	Payload struct {
		Iss  string      `json:"iss"`   // 发行者
		Sub  string      `json:"sub"`   // Subject，主题
		Aud  string      `json:"aud"`   // Audience，观众
		Exp  int64       `json:"exp"`   // 过期时间
		Iat  int64       `json:"iat"`   // 发行时间
		Data interface{} `json:"other"` // 请求数据
	}

	// Header JWT的header
	Header struct {
		Typ string `json:"typ"` // 类型
		Alg string `json:"alg"` // 使用的加密算法
	}
)

var (
	HS512 = "HS512"
	HS256 = "HS256"
)

// Builder 构建JWT
func Builder(secKey string, exp int64, alg string, data interface{}) *JWT {
	return &JWT{
		Payload: Payload{Exp: exp, Data: data},
		Header:  Header{Typ: "jwt", Alg: alg},
		secKey:  secKey,
	}
}

// Decode 解码
func Decode(jwtStr, secKey string) (*JWT, error) {
	jwts := strings.Split(jwtStr, ".")
	if len(jwts) < 3 {
		return nil, errors.New("JWT 缺少足够的字符")
	}
	header := Header{}
	err := header.Decode(jwts[0])
	if err != nil {
		return nil, err
	}
	payload := Payload{}
	err = payload.Decode(jwts[1])
	if err != nil {
		return nil, err
	}
	jwt := &JWT{Payload: payload, Header: header, Sign: jwts[2], secKey: secKey, jwtString: jwtStr}
	return jwt, nil
}

// 得到JWT字符串
func (jwt *JWT) String() string {
	if jwt.jwtString != "" {
		return jwt.jwtString
	}
	encPl := jwt.Payload.Encode()
	encH := jwt.Header.Encode()

	jwt.Sign = jwt.HSEnc(encH, encPl)

	jwt.jwtString = encH + "." + encPl + "." + jwt.Sign

	return jwt.jwtString
}

// HSEnc SH256/SH512
func (jwt *JWT) HSEnc(encHeader, encPayload string) string {
	var hashMac hash.Hash
	if jwt.Header.Typ == "HS512" {
		hashMac = hmac.New(sha512.New, []byte(jwt.secKey)) // HS512
	} else {
		hashMac = hmac.New(sha256.New, []byte(jwt.secKey)) // 应该是HS256
	}
	hashMac.Write([]byte(encHeader + "." + encPayload))
	return base64.StdEncoding.EncodeToString(hashMac.Sum(nil))
}

// Check 检查Token状态
func (jwt *JWT) check() error {
	jwts := strings.Split(jwt.jwtString, ".")

	sign := jwt.HSEnc(jwts[0], jwts[1])
	if sign != jwt.Sign {
		return errors.New("Token签名错误")
	}
	if jwt.Payload.Exp > time.Now().Unix() {
		return errors.New("Token过期")
	}
	return nil
}

// HelpCheck 检测token 包含check检测
func HelpCheck(req *http.Request, cookieName string, secKey string) (*JWT, error) {
	cookie, err := req.Cookie(cookieName)
	if err != nil {
		return nil, err
	}
	jwt, err := Decode(cookie.Value, secKey)
	if err != nil {
		return nil, err
	}
	err = jwt.check()
	if err != nil {
		return nil, err
	}
	return jwt, nil
}

//////////////////////////////////////////// 编解码 /////////////////////////////////////////////////

// Encode base64编码
func (payload *Payload) Encode() string {
	pl, err := json.Marshal(payload)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(pl)
}

// Encode base64编码
func (header *Header) Encode() string {
	h, err := json.Marshal(header)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(h)
}

// Decode JWTPayload 解码
func (payload *Payload) Decode(origin string) error {
	if origin == "" {
		return errors.New("payload 为空")
	}
	payloadByte, err := base64.StdEncoding.DecodeString(origin)
	if err != nil {
		return err
	}
	return json.Unmarshal(payloadByte, payload)
}

// Decode JWTHeader 解码
func (header *Header) Decode(origin string) error {
	if origin == "" {
		return errors.New("header 为空")
	}
	headerByte, err := base64.StdEncoding.DecodeString(origin)
	if err != nil {
		return err
	}
	return json.Unmarshal(headerByte, header)
}
