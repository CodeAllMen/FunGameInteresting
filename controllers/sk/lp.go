/**
  create by yy on 2020/3/27
*/

package sk

import "github.com/astaxie/beego"

type SubPage struct {
	beego.Controller
}

func (s *SubPage) Lp() {
	s.TplName = "sk/lp.html"
}

func (s *SubPage) Terms() {
	s.TplName = "sk/terms.html"
}

func (s *SubPage) Privacy() {
	s.TplName = "sk/privacy.html"
}

func (s *SubPage) Help() {
	s.TplName = "sk/help.html"
}
