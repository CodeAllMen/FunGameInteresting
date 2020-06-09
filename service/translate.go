/**
  create by yy on 2020/6/6
*/

package service

import (
	"fmt"
	"github.com/MobileCPX/FunGameInteresting/models"
	"github.com/astaxie/beego/httplib"
	"net/url"
	"time"
)

type responseStruct struct {
	Code   int    `json:"code"`
	Data   string `json:"data"`
	Status int    `json:"status"`
}

func TranslateGames() {
	var (
		nameSk string
	)

	games := models.GetGames()

	for _, game := range games {

		fmt.Println(fmt.Sprintf("translate: =====> %v", game.Name))
		nameSk = translateGames(game)

		models.UpdateGames(fmt.Sprintf("%v", game.Id), nameSk)

		time.Sleep(time.Second * 3)
	}
}

func translateGames(games models.Games) string {

	var (
		err error
	)

	requestUrl := "http://52.28.109.246:8077/api/translate_com?%v"

	queryData := url.Values{}
	queryData.Add("content", games.Name)
	queryData.Add("src", "auto")
	queryData.Add("dest", "sk")

	requestUrl = fmt.Sprintf(requestUrl, queryData.Encode())

	request := httplib.Get(requestUrl)

	response := &responseStruct{}

	if err = request.ToJSON(response); err != nil || response.Code != 0 {
		fmt.Println(err)
		fmt.Println(response)
		return ""
	}

	return response.Data
}
