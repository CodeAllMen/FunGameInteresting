package controllers

import (
	"fmt"
	"github.com/MobileCPX/FunGameInteresting/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c MainController) Prepare() {
	var (
		err          error
		categoryList []models.CategoryModel
	)

	if categoryList, err = models.GetCategoryList(); err != nil {
		fmt.Println(err)
	}

	c.Data["CategoryList"] = categoryList
}

func (c *MainController) Get() {
	subID := c.GetString("subID")
	if subID != "" {
		c.SetSession("user", subID)
	}
	c.Data["games"] = models.GetGames()
	fmt.Println("asda")
	c.TplName = "index.html"
}

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.Data["status"] = "0"
	c.TplName = "login.html"
}

type PageController struct {
	beego.Controller
}

func (c *PageController) Get() {
	sp := c.Ctx.Input.Param(":sp")
	country := c.Ctx.Input.Param(":country")
	pagetype := c.Ctx.Input.Param(":type")

	c.TplName = sp + "/" + country + "/" + pagetype + ".html"
}

type GamePage struct {
	beego.Controller
}

func (c *GamePage) Get() {
	// id := c.Ctx.Input.Param(":id")
	// url := models.GetGamesUrl(id)
	// c.Redirect(url, 302)
	if c.GetSession("user") != nil {
		id := c.Ctx.Input.Param(":id")
		url := models.GetGamesUrl(id)
		c.Redirect(url, 302)
	} else {
		c.TplName = "login.html"
	}
}

func (c *MainController) CategoryPage() {
	var (
		err      error
		gameList []models.Games
		category *models.CategoryModel
	)

	categoryId := c.GetString("category_id")
	fmt.Println(categoryId)

	if gameList, err = models.GetGameListByCategory(categoryId); err != nil {
		fmt.Println(err)
	}

	if category, err = models.GetCategoryById(categoryId); err != nil {
		fmt.Println(err)
	}

	c.Data["games"] = gameList
	c.Data["Title"] = category.Name

	c.TplName = "category.html"
}