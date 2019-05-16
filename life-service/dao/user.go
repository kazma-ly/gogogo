package dao

import (
	"errors"
	"life-service/sqlx"
	"life-service/tool"
	"time"
)

type (
	// User 用户实体
	User struct {
		Id         string    `json:"id"`                 // ID
		Username   string    `json:"username"`           // 用户名
		Nickname   string    `json:"nickname"`           // 用户昵称
		Password   string    `json:"password,omitempty"` // 密码 omitempty 允许不输出json
		CreateTime time.Time `json:"createTime"`         // 加入时间
	}
)

// Login 登陆
func Login(username, password string) (*User, error) {
	args := []interface{}{username, password}
	base := sqlx.New("SELECT id, username, nickname, password, create_time FROM t_user WHERE username = ? and password = ?", args)

	user := User{}
	err := base.Query(&user)
	return &user, err
}

// Register 注册 没有error就是ok
func Register(username, nickname, password string) error {
	exits, e := findByUserName(username)
	if e != nil {
		return e
	}
	if exits {
		return errors.New("user exits")
	}
	rNum := tool.MakeRandomNum()
	args := []interface{}{rNum, username, nickname, password, time.Now()}
	base := sqlx.New("INSERT INTO t_user (id, username, nickname, password, create_time) VALUES (?, ?, ?, ?, ?)", args)
	_, err := base.Exec()
	return err
}

//
func findByUserName(username string) (bool, error) {
	args := []interface{}{username}
	base := sqlx.New("SELECT id, username, nickname, password, create_time FROM t_user WHERE username = ?", args)
	user := User{}
	err := base.Query(&user)
	if err != nil {
		return false, err
	}
	return user.Id != "", nil
}
