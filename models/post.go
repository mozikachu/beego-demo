package models

import (
	"github.com/astaxie/beego/orm"
)

type Post struct {
	Id    int
	Title string
	User  *User `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Post))
}

func AddPost(p *Post) (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(p)
	return
}

func GetPost(id int) (p *Post, err error) {
	o := orm.NewOrm()
	p = &Post{Id: id}
	err = o.Read(p)
	return
}

func GetPostWithUsername(username string) (ps []*Post, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("post")
	_, err = qs.Filter("User__UserName", username).All(&ps)
	return
}

func UpdatePostWithUsername(params orm.Params, username string) (err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("post")
	_, err = qs.Filter("User__UserName", username).Update(params)
	return
}

func DeletePostWithUserName(username string) (err error) {
	qs := getTableQS("post")
	_, err = qs.Filter("User__UserName", username).Delete()
	return
}

func getTableQS(ptrStructOrTableName interface{}) (qs orm.QuerySeter) {
	o := orm.NewOrm()
	qs = o.QueryTable("post")
	return
}
