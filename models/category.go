/**
  create by yy on 2020/6/17
*/

package models

import "github.com/astaxie/beego/orm"

type CategoryModel struct {
	Id   int    `orm:"pk;auto" json:"id"`
	Name string `json:"name"`
}

func (c *CategoryModel) TableName() string {
	return "category"
}

func getCategoryTableName() string  {
	return "category"
}

func GetCategoryList() (categoryList []CategoryModel, err error) {

	db := orm.NewOrm()

	_, err = db.QueryTable(getCategoryTableName()).All(&categoryList)

	return
}

func GetCategoryById(categoryId string) (category *CategoryModel, err error) {
	var (
		categoryModel CategoryModel
	)

	db := orm.NewOrm()

	err = db.QueryTable(getCategoryTableName()).Filter("id", categoryId).One(&categoryModel)

	category = &categoryModel

	return
}