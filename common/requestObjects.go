package common

type MaterialObj struct {
	MaterialName string `json:"materialName" valid:"Required;Match(/^[a-zA-Z0-9 ]{2,100}$/)"`
	MaterialType string `json:"materialType" valid:"Required;Match(/^[a-zA-Z0-9 ]{2,100}$/)"`
	MaterialGrade string `json:"materialGrade" valid:"Required;Match(/^[0-9]{1,2}$/)"`
}