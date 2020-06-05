package controllers

import (
	"github.com/MobileCPX/FunGameInteresting/models"
	"github.com/astaxie/beego"
)

type RegisteredControllers struct {
	beego.Controller
}

func (c *RegisteredControllers) Get() {
	var user models.Users
	user.UserName = c.GetString("user_name")
	user.Password = c.GetString("password")
	user.Country = c.GetString("country")
	user.Sp = c.GetString("sp")
	if user.UserName != "" {
		models.RegistereUser(user)
	}
	c.Ctx.WriteString("OK")
}
