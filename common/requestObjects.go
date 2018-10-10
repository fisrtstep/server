package common

type MaterialObj struct {
	MaterialName string `json:"materialName" valid:"Required;Match(/^[a-zA-Z0-9 ]{2,100}$/)"`
	MaterialType string `json:"materialType" valid:"Required;Match(/^[a-zA-Z0-9 ]{2,100}$/)"`
	MaterialGrade string `json:"materialGrade" valid:"Required;Match(/^[0-9]{1,2}$/)"`
}

type VendorObj struct {
	Name string `json:"name" valid:"Required;Match(/^[a-zA-Z ]{3,150}$/)"`
	Email string `json:"email" valid:"Required;Email"`
	Cell string `json:"cell" valid:"Required"`
	Phone string `json:"phone"`
	Fax string `json:"fax"`
	Address string `json:"address" valid:"Match(/^[a-zA-Z0-9 .,-]*$/)"`
}
type ClientObj struct {
	Name string `json:"name" valid:"Required;Match(/^[a-zA-Z ]{3,150}$/)"`
	Email string `json:"email" valid:"Required;Email"`
	Cell string `json:"cell" valid:"Required"`
	Phone string `json:"phone"`
	Fax string `json:"fax"`
	Address string `json:"address" valid:"Match(/^[a-zA-Z0-9 .,-]*$/)"`
}