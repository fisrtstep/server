package routers

import (
	"firstStep/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//All the routes to the material CRUD
    beego.Router("/v1/materials", &controllers.MaterialController{}, "get:GetMaterials")
	beego.Router("/v1/material/:id", &controllers.MaterialController{}, "get:GetMaterialById")
	beego.Router("/v1/addMaterials", &controllers.MaterialController{}, "post:AddMaterial")
	beego.Router("/v1/material/:id", &controllers.MaterialController{}, "put:ModifyMaterialById")
    beego.Router("/v1/material/:id", &controllers.MaterialController{}, "delete:DeleteMaterial")
    //All the routes to the vendor CRUD
    beego.Router("/v1/vendors", &controllers.VendorController{}, "get:GetAllVendors")
	beego.Router("/v1/vendor/:id", &controllers.VendorController{}, "get:GetVendorById")
    beego.Router("/v1/addVendor", &controllers.VendorController{}, "post:AddVendor")
    beego.Router("/v1/vendor/:id", &controllers.VendorController{}, "put:ModifyVendorById")
    beego.Router("/v1/vendor/:id", &controllers.VendorController{}, "delete:DeleteVendor")
	//All the routes to the client CRUD
	beego.Router("/v1/clients", &controllers.ClientController{}, "get:GetAllClients")
	beego.Router("/v1/client/:id", &controllers.ClientController{}, "get:GetClientById")
	beego.Router("/v1/addClient", &controllers.ClientController{}, "post:AddClient")
	beego.Router("/v1/client/:id", &controllers.ClientController{}, "put:ModifyClientById")
	beego.Router("/v1/client/:id", &controllers.ClientController{}, "delete:DeleteClient")
}
