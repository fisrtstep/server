package common

import "github.com/astaxie/beego/orm"

type Material struct {
	Mid int `orm:"pk"`
	MaterialName string `orm:"column(materialName)"`
	MaterialType string `orm:"column(materialType)"`
	MaterialGrade string `orm:"column(materialGrade)"`
}

func init(){
	orm.RegisterModel(new(Material))
}