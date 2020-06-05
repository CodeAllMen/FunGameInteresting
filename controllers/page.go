package controllers

import (
	"github.com/astaxie/beego"
)

type LpPage struct {
	beego.Controller
}

func (this *LpPage) Get() {
	this.TplName = "dm/nl/lp.html"
}

type HelpPage struct {
	beego.Controller
}

func (this *HelpPage) Get() {
	this.TplName = "dm/nl/help.html"
}

type PrivacyPage struct {
	beego.Controller
}

func (this *PrivacyPage) Get() {
	this.TplName = "dm/nl/privacy.html"
}

type TNCPage struct {
	beego.Controller
}

func (this *TNCPage) Get() {
	this.TplName = "dm/nl/tnc.html"
}
