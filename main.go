package main

import (
	_ "github.com/MobileCPX/FunGameInteresting/initial"
	_ "github.com/MobileCPX/FunGameInteresting/routers"
	"github.com/astaxie/beego"
	// "github.com/MobileCPX/FunGame/models"
)

func main() {
	// models.FileToSql()
	beego.Run()
}
