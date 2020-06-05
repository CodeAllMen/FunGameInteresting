package dimoco_pl

import (
	"github.com/astaxie/beego"
)

const (
	//UnsubSuccessCode 退订成功
	UnsubSuccessCode = "0"

	//MsisdnIsEmptyCode 退订电话号码为空
	MsisdnIsEmptyCode = "-1"

	//MsisdnNotExistCode 退订电话号码不存在
	MsisdnNotExistCode = "-2"

	// XMLErrorCode xml解析错误
	XMLErrorCode = "-3"

	//UnsubFaieldCode  退订失败
	UnsubFaieldCode = "-4"
)

type PLLPPageControllers struct {
	beego.Controller
}

func (c *PLLPPageControllers) Get() {
	trackID := c.GetString("track")
	if trackID == "" {
		c.Redirect("http://pl.leadernet-hk.com/identify", 302)
		return
	}
	c.Data["track"] = trackID
	c.Data["service_name"] = "Click4FunGame"
	c.TplName = "dm/pl/lp.html"
}

type PLWelcomePageControllers struct {
	beego.Controller
}

type UnsubPageControllers struct {
	beego.Controller
}

func (c *UnsubPageControllers) Get() {
	msisdn := ""
	if c.Ctx.Input.Session("user") != nil {
		msisdn = c.Ctx.Input.Session("user").(string)
	}
	c.Data["msisdn"] = msisdn
	c.TplName = "dm/pl/unsub.html"
}

type UnsubResultControllers struct {
	beego.Controller
}

func (c *UnsubResultControllers) Get() {
	code := c.GetString("code")
	if code == UnsubSuccessCode {
		c.Data["text"] = "Pomyślnie anulowano subskrypcję"
		c.Data["URL"] = "/"
	} else {
		c.Data["text"] = "Anulowanie subskrypcji nie powiodło się"
		c.Data["failed_desc"] = "Twój numer telefonu komórkowego nie może zostać rozpoznany. Podaj swój numer telefonu, aby zrezygnować z subskrypcji"
		c.Data["URL"] = "/unsub/msisdn"
	}
	c.TplName = "dm/pl/unsub_result.html"
}

type MsisdnUnsubPageControllers struct {
	beego.Controller
}

func (c *MsisdnUnsubPageControllers) Get() {
	c.TplName = "dm/pl/unsub_msisdn.html"
}

type DisclaimerController struct {
	beego.Controller
}

func (c *DisclaimerController) Get() {
	c.TplName = "dm/pl/disclaimer.html"
}
