package models

import (
	"github.com/astaxie/beego/orm"
)

type Games struct {
	Id      int64  `orm:"pk;auto"`
	Url     string `orm:"size(100)"`
	Img     string `orm:"size(100)"`
	Name    string `orm:"size(100)"`
	NameSk  string `orm:"size(100)" json:"name_sk"`
	PlayNum int
}

type Users struct {
	Id       int64 `orm:"pk;auto"`
	UserName string
	Password string
	Sp       string
	Country  string
}

func init() {
	orm.RegisterModel(new(Games), new(Users))
}
