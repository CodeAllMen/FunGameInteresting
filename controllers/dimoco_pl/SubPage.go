/**
  create by yy on 2020/3/19
*/

package dimoco_pl

import (
	"github.com/MobileCPX/FunGameInteresting/models"
	"github.com/astaxie/beego"
)

type SubPage struct {
	beego.Controller
}

func (s *SubPage) Lp() {
	// 获取 track_id
	trackId := s.GetString("tid")

	// 如果 trackId 不存在，则 获取
	if trackId == "" {
	}

	s.Data["TrackId"] = trackId
	s.TplName = "dm/ae/lp.html"
}

func (s *SubPage) Tan() {
	// 获取 track_id
	trackId := s.GetString("tid")
	s.Data["TrackId"] = trackId
	s.Data["service_name"] = "Interesting GAME"
	s.TplName = "dm/ae/tan.html"
}

func (s *SubPage) Confirm() {
	s.TplName = "dm/ae/confirm.html"
}

func (s *SubPage) Condition() {
	s.TplName = "dm/ae/tnc.html"
}

func (s *SubPage) Help() {
	s.TplName = "dm/ae/help.html"
}

func (s *SubPage) Privacy() {
	s.TplName = "dm/ae/privacy.html"
}

func (s *SubPage) Welcome() {
	s.TplName = "dm/ae/welcome.html"
}

func (s *SubPage) ConfirmSMS() {
	msisdn := s.GetString("msisdn")
	s.Data["Msisdn"] = msisdn
	s.TplName = "dm/ae/confirm_sms.html"
}

func (s *SubPage) ConfirmSMSAjax() {
	msisdn := s.GetString("msisdn")

	uid := models.CheckUser(msisdn)
	if uid == "" {
		s.Data["json"] = failedReply("Zkontrolujte svou SMS a dokončete registraci")
	} else {
		s.Data["json"] = successReply("")
	}

	s.ServeJSON()
}

func (s *SubPage) Register() {
	msisdn := s.GetString("msisdn")

	// 将电话号码或者那个id存到数据库
	user := models.Users{
		UserName: msisdn,
		Password: msisdn,
		Sp:       "Dimoco",
		Country:  "CZ",
	}
	models.RegistereUser(user)

	s.Data["json"] = msisdn
	s.ServeJSON()
}

type reply struct {
	Status int    `json:"status"`
	Desc   string `json:"desc"`
}

func successReply(desc string) *reply {
	return &reply{
		Status: 1,
		Desc:   desc,
	}
}

func failedReply(desc string) *reply {
	return &reply{
		Status: 0,
		Desc:   desc,
	}
}
