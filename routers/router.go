package routers

import (
	"firstStep/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/v1/materials", &controllers.MainController{}, "get:GetMaterials")
	beego.Router("/v1/materials/:id", &controllers.MainController{}, "get:GetMaterialById")
	beego.Router("/v1/add/materials", &controllers.MainController{}, "post:AddMaterial")
	beego.Router("/v1/materials/:id", &controllers.MainController{}, "put:ModifyMaterialById")
}
