/**
  create by yy on 2020/3/19
*/

package dimoco_pl

import "github.com/astaxie/beego"

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