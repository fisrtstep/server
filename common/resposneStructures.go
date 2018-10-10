package common

import (
	"regexp"
	"github.com/astaxie/beego"
)

type ResposneStruct struct {
	Status int `json:"status"`
	Message string `json:"message"`
}

func ValidateId (id string) (result bool){
	beego.Info(id)
	test, testError := regexp.MatchString("^[0-9]$", id)
	beego.Info(test, testError)
	if testError != nil {
		return false
	}else{
		return test
	}
}