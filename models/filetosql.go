package models

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"

	"github.com/astaxie/beego/orm"
)

func FileToSql() {
	f, err := os.Open("log")
	defer f.Close()
	if err != nil {

	} else {
		rd := bufio.NewReader(f)
		for {
			line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
			if err != nil || io.EOF == err {
				break
			}
			line = strings.Replace(line, "\n", "", -1)
			data := strings.Split(line, "####")
			o := orm.NewOrm()
			fmt.Println(data)
			var game Games
			o.QueryTable("games").Filter("img", data[2]).One(&game)
			fmt.Println(game)
			if game.Id == 0 {
				game.Img = data[2]
				game.Name = data[1]
				game.Url = data[0]
				game.PlayNum = int(rand.Int63n(1765237))
				o.Insert(&game)
			}

		}
	}
}
