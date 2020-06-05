package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
)

func GetGames() []Games {
	o := orm.NewOrm()
	var games []Games
	o.QueryTable("games").All(&games)
	return games
}

func GetGamesUrl(id string) string {
	o := orm.NewOrm()
	var games Games
	id_int, _ := strconv.Atoi(id)
	o.QueryTable("games").Filter("id", id_int).One(&games)
	return games.Url
}
