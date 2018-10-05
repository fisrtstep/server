package main

import (
	_ "firstStep/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "firstStep/common"
	"net/http"
	"encoding/json"
)
func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:@/c2h")
}

func page_not_found(rw http.ResponseWriter, r *http.Request){
	type PageNotFound struct {
		Status int `json"status"`
		Message string `json:"message"`
	}
	var resp PageNotFound
	resp.Status = 404
	resp.Message = "Requeste URL is invalid"
	response, _ := json.Marshal(resp)
	rw.Header().Set("status", "404")
	rw.Header().Set("ContentType", "application/json")
	rw.Write(response)
}

func main() {
	beego.ErrorHandler("404", page_not_found)
	beego.Run()
}

