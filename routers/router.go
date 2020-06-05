package routers

import (
	"github.com/MobileCPX/FunGameInteresting/controllers"
	"github.com/MobileCPX/FunGameInteresting/controllers/dimoco_pl"
	"github.com/MobileCPX/FunGameInteresting/controllers/sk"
	"github.com/MobileCPX/FunGameInteresting/controllers/slovenia"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})

	beego.Router("/registere/username", &controllers.RegisteredControllers{}) // 电话号码注册用户

	// Dimoco
	// beego.Router("/dm/pl/lp", &dimoco_pl.PLLPPageControllers{})
	// beego.Router("/dm/pl/welcome", &dimoco_pl.SubResultControllers{})
	// beego.Router("/sub/result/page", &dimoco_pl.SubResultControllers{})
	beego.Router("/cz/welcome", &dimoco_pl.SubPage{}, "*:Welcome")

	beego.Router("/lp", &dimoco_pl.SubPage{}, "*:Lp")
	beego.Router("/tan", &dimoco_pl.SubPage{}, "*:Tan")
	beego.Router("/confirm", &dimoco_pl.SubPage{}, "*:Confirm")
	beego.Router("/tnc", &dimoco_pl.SubPage{}, "*:Condition")
	beego.Router("/help", &dimoco_pl.SubPage{}, "*:Help")
	beego.Router("/privacy", &dimoco_pl.SubPage{}, "*:Privacy")

	beego.Router("/sv/lp", &slovenia.SubPage{}, "*:Lp")
	beego.Router("/sv/tan", &slovenia.SubPage{}, "*:Tan")
	beego.Router("/sv/confirm", &slovenia.SubPage{}, "*:Confirm")
	beego.Router("/sv/tnc", &slovenia.SubPage{}, "*:Condition")
	beego.Router("/sv/help", &slovenia.SubPage{}, "*:Help")
	beego.Router("/sv/privacy", &slovenia.SubPage{}, "*:Privacy")

	beego.Router("/?:sp/?:country/?:type", &controllers.PageController{})
	beego.Router("/game/?:id", &controllers.GamePage{})

	beego.Router("/user/?:method", &controllers.UserController{})

	//sk
	beego.Router("/sk/lp", &sk.SubPage{}, "*:Lp")
	beego.Router("/sk/terms", &sk.SubPage{}, "*:Terms")
	beego.Router("/sk/privacy", &sk.SubPage{}, "*:Privacy")
	beego.Router("/sk/help", &sk.SubPage{}, "*:Help")

	// 波兰退订页面
	beego.Router("/unsub/page", &dimoco_pl.UnsubPageControllers{})
	beego.Router("/unsub/result", &dimoco_pl.UnsubResultControllers{})
	beego.Router("/unsub/msisdn", &dimoco_pl.MsisdnUnsubPageControllers{})

	// beego.Router("/pl/start",&dimoco_pl.)

	// 页面
	beego.Router("/l", &controllers.LpPage{})
	beego.Router("/h", &controllers.HelpPage{})
	beego.Router("/p", &controllers.PrivacyPage{})
	beego.Router("/t", &controllers.TNCPage{})
}
