package user

import (
	"encoding/json"
	"io/ioutil"
	"life-service/const"
	"life-service/dao"
	"life-service/httpfly/httpres"
	"life-service/httpfly/jwt"
	"life-service/logx"
	"life-service/session"
	"life-service/tool"
	"life-service/violet"
	"net/http"
	"time"
)

// Login 登陆
func Login(c *violet.Context) {
	req := c.Request
	bs, _ := ioutil.ReadAll(req.Body)
	var u dao.User
	json.Unmarshal(bs, &u)
	user, err := dao.Login(u.Username, tool.Md5(u.Password))
	if err != nil {
		logx.LogInfo(err)
		c.WriteJSON(httpres.Create(500, "登录失败", nil, false))
		return
	}
	if user == nil || user.Id == "" {
		logx.LogInfo(err)
		c.WriteJSON(httpres.Create(301, "请检查用户名或密码", nil, false))
	} else {
		user.Password = "" // 置空密码

		jwt := jwt.Builder(consts.SECKEY, 7200, "HS256", user.Id)
		c.AddCookie(&http.Cookie{
			Name:     "jwt",
			Value:    jwt.String(),
			Path:     "/",
			Expires:  time.Now().Add(time.Second * 60 * 30),
			HttpOnly: true,
		})
		c.WriteJSON(httpres.Create(200, "登陆成功", user, true))
	}
}

// Register 注册
func Register(c *violet.Context) {
	req := c.Request
	bs, _ := ioutil.ReadAll(req.Body)
	var u dao.User
	json.Unmarshal(bs, &u)
	err := dao.Register(u.Username, "", tool.Md5(u.Password))
	if err == nil {
		logx.LogInfo(err)
		c.WriteJSON(httpres.Create(200, "注册成功", nil, true))
	} else {
		logx.LogInfo(err)
		c.WriteJSON(httpres.Create(301, "注册失败: "+err.Error(), nil, true))
	}
}

// Logout 注销
func Logout(c *violet.Context) {
	sess := c.Session()
	sessionID := sess.SessionID()
	if sessionID != "" && len(sessionID) > 0 {
		sess.DestroySession(sessionID)
	}
	c.AddCookie(&http.Cookie{
		Name:     session.SESSION_NAME,
		MaxAge:   -1,
		Path:     "/",
		HttpOnly: true,
	})
	c.AddCookie(&http.Cookie{
		Name:     consts.TOKEN_NAME,
		MaxAge:   -1,
		Path:     "/",
		HttpOnly: true,
	})
	c.WriteFile("static/views/login.html", "text/html")
}
