package dao

import (
	"database/sql"
	"life-service/sqlx"
	"life-service/tool"
	"time"
)

type (
	// Life 实体
	Life struct {
		Id         string    `json:"id"`         // id
		UId        string    `json:"uid"`        // uid
		Title      string    `json:"title"`      // 标题
		Subtitle   string    `json:"subtitle"`   // 副标题
		CreateTime time.Time `json:"createTime"` // 插入时间
	}
)

// QueryLifeByPage 查询所有的Life
func QueryLife(id string, page string) ([]Life, error) {
	args := []interface{}{id, page}
	sqlStr := `select tl.id, title, subtitle, uid, tl.create_time
from t_life as tl
       left join t_user as tu on tu.id = tl.uid
where tu.id = ?
limit 10 offset ?`
	base := sqlx.New(sqlStr, args)
	var lifes []Life
	err := base.Querys(&lifes)
	if err != nil {
		return nil, err
	}
	return lifes, nil
}

// QueryLifeByID 通过id查询
func QueryLifeByID(id string) (*Life, error) {
	args := []interface{}{id}
	base := sqlx.New("select id, title, subtitle, pics, create_time from t_life where id = ?", args)

	var life Life
	err := base.Query(&life)
	if err != nil {
		return nil, err
	}
	return &life, nil
}

// InsertLife 插入life
func InsertLife(life Life) (r sql.Result, e error) {
	args := []interface{}{tool.MakeRandomNum(), life.Title, life.Subtitle, time.Now()}
	base := sqlx.New("insert into t_life (id, title, subtitle, create_time) values (?, ?, ?, ?)", args)
	return base.Exec()
}

// DeleteLifeByID 删除life 通过id
func DeleteLifeByID(id string) error {
	args := []interface{}{id}
	base := sqlx.New("delete from t_life where id = ?", args)
	_, err := base.Exec()
	return err
}
