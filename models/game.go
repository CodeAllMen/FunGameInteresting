package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
)

func GetGames() []Games {
	o := orm.NewOrm()
	var games []Games
	// o.QueryTable("games").All(&games)
	_, _ = o.Raw("select * from games where name_sk!=''").QueryRows(&games)
	return games
}

func GetGamesUrl(id string) string {
	o := orm.NewOrm()
	var games Games
	id_int, _ := strconv.Atoi(id)
	o.QueryTable("games").Filter("id", id_int).One(&games)
	return games.Url
}

func UpdateGames(id, nameSk string) {
	o := orm.NewOrm()
	var (
		err   error
		games Games
	)
	idInt, _ := strconv.Atoi(id)
	o.QueryTable("games").Filter("id", idInt).One(&games)

	games.NameSk = nameSk

	if _, err = o.Update(&games); err != nil {
		fmt.Println(err)
	}

}
