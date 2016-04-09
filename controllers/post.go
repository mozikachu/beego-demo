package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/mozikachu/beego-demo/models"
)

type PostController struct {
	beego.Controller
}

//-------------------
// 向数据库添加新 post
//-------------------
// @router /:uid [post]
func (c *PostController) Post() {
	defer c.ServeJSON()

	uid, err := c.ParseParamInt(":uid")
	if err != nil {
		c.Data["json"] = err.Error()
		return
	}

	p := models.Post{User: &models.User{Id: uid}}
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &p)
	if err != nil {
		c.Data["json"] = err.Error()
		return
	}
	if err = models.AddPost(&p); err != nil {
		c.Data["json"] = err.Error()
		return
	}
	c.Data["json"] = "Succeed"
}

//------------------
// 根据 ID 查询 post
//------------------
// @router /:postId [get]
func (c *PostController) Get() {
	defer c.ServeJSON()

	postId, err := c.ParseParamInt(":postId")
	if err != nil {
		c.Data["json"] = err.Error()
		return
	}
	p, err := models.GetPost(postId)
	if err != nil {
		c.Data["json"] = err.Error()
		return
	}
	c.Data["json"] = p
}

//-----------------------------------
// 关联查询：查询指定用户名下的所有 posts
//-----------------------------------
// @router /username/:username [get]
func (c *PostController) GetWithUsername() {
	defer c.ServeJSON()

	username := c.Ctx.Input.Param(":username")
	ps, err := models.GetPostWithUsername(username)
	if err != nil {
		c.Data["json"] = err.Error()
		return
	}
	c.Data["json"] = ps
}

//-----------------------------------
// 关联更新：更新指定用户名下的所有 posts
//-----------------------------------
// @router /username/:username [put]
func (c *PostController) PutWithUsername() {
	defer c.ServeJSON()

	username := c.Ctx.Input.Param(":username")
	var p orm.Params
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &p); err != nil {
		c.Data["json"] = err.Error()
		return
	}
	if err := models.UpdatePostWithUsername(p, username); err != nil {
		c.Data["json"] = err.Error()
		return
	}
	c.Data["json"] = "Succeed"
}

//-----------------------------------
// 关联删除：删除指定用户名下的所有 posts
//-----------------------------------
// @router /username/:username [delete]
func (c *PostController) DeleteWithUsername() {
	defer c.ServeJSON()

	username := c.Ctx.Input.Param(":username")
	if err := models.DeletePostWithUserName(username); err != nil {
		c.Data["json"] = err.Error()
	}
	c.Data["json"] = "Succeed"
}

// util
func (c *PostController) ParseParamInt(param string) (ret int, err error) {
	raw := c.Ctx.Input.Param(param)
	ret, err = strconv.Atoi(raw)
	return
}
