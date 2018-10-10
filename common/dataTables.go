package common

import (
	"github.com/astaxie/beego/orm"
)

type Material struct {
	Mid int `orm:"pk"`
	MaterialName string `orm:"column(materialName)"`
	MaterialType string `orm:"column(materialType)"`
	MaterialGrade string `orm:"column(materialGrade)"`
}

type Vendor struct {
	VendorId int `orm:"pk;column(vendorId)"`
	Name string `orm:"column(name)"`
	Email string `orm:"column(email)"`
	Cell string `orm:"column(cell)"`
	Phone string `orm:"column(phone)"`
	Fax string `orm:"column(fax)"`
	Address string `orm:"column(address)"`
}

type Client struct {
	ClientId int `orm:"pk;column(clientId)"`
	Name string `orm:"column(name)"`
	Email string `orm:"column(email)"`
	Cell string `orm:"column(cell)"`
	Phone string `orm:"column(phone)"`
	Fax string `orm:"column(fax)"`
	Address string `orm:"column(address)"`
}

func init(){
	orm.RegisterModel(new(Material), new(Vendor), new(Client))
}