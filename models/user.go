package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int
	UserName string
	Password string
}

func init() {
	orm.RegisterModel(new(User))
}

func AddUser(u *User) (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(u)
	return
}
