package life

import (
	"life-service/const"
	"life-service/dao"
	"life-service/httpfly/httpres"
	"life-service/httpfly/jwt"
	"life-service/logx"
	"life-service/violet"
)

// Save 保存记录
func Save(c *violet.Context) {
	req := c.Request
	_, err := jwt.HelpCheck(req, consts.TOKEN_NAME, consts.SECKEY)
	if err != nil {
		c.WriteJSON(httpres.NotLogin())
		return
	}
	req.ParseForm()
	title := req.Form.Get("title")
	subtitle := req.Form.Get("subtitle")
	if title == "" || subtitle == "" {
		c.WriteJSON(httpres.Create(401, "参数错误", nil, false))
		return
	}

	// 插入事件记录
	lifeEntity := dao.Life{Title: title, Subtitle: subtitle}
	_, err = dao.InsertLife(lifeEntity)
	if err != nil {
		logx.LogInfo(err)
		c.WriteJSON(httpres.Fail())
		return
	}
	c.WriteJSON(httpres.Success(nil))
}

// All 查询记录
func All(c *violet.Context) {
	jwt, err := jwt.HelpCheck(c.Request, consts.TOKEN_NAME, consts.SECKEY)
	if err != nil {
		c.WriteJSON(httpres.NotLogin())
		return
	}
	uid := jwt.Payload.Data
	page := c.GetPathVar("page")
	lifes, err := dao.QueryLife(uid.(string), page)
	if err != nil {
		logx.LogInfo(err)
		c.WriteJSON(httpres.Create(500, err.Error(), nil, false))
	} else {
		c.WriteJSON(httpres.Create(200, "成功", lifes, true))
	}
}

// Delete 删除
func Delete(c *violet.Context) {
	_, err := jwt.HelpCheck(c.Request, consts.TOKEN_NAME, consts.SECKEY)
	if err != nil {
		c.WriteJSON(httpres.NotLogin())
		return
	}
	// id 是life的id
	id := c.GetPathVar("id")
	err = dao.DeleteLifeByID(id)
	if err != nil {
		logx.LogInfo(err)
		c.WriteJSON(httpres.Fail())
		return
	}
	c.WriteJSON(httpres.Success)
}
