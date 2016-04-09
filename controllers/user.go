package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/mozikachu/beego-demo/models"
)

type UserController struct {
	beego.Controller
}

//-------------------
// 向数据库添加新 user
//-------------------
// @router / [post]
func (c *UserController) Post() {
	defer c.ServeJSON()

	u := models.User{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &u)
	if err != nil {
		c.Data["json"] = err.Error()
		return
	}
	if err = models.AddUser(&u); err != nil {
		c.Data["json"] = err.Error()
		return
	}
	c.Data["json"] = "Succeed"
}
