package controllers

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"firstStep/models"
	"encoding/json"
	"firstStep/common"
	"strconv"
	"github.com/astaxie/beego/validation"
)

type MaterialController struct {
	beego.Controller
}

//Create a new record in the material table
func (this *MaterialController) AddMaterial(){
	//Read the input parameters in the object, since it is unmarshaled only the required fields will be read
	var material common.MaterialObj
	unmarshalError := json.Unmarshal(this.Ctx.Input.RequestBody, &material)
	//Check unmarshal errors, if there are errors then send 400 bad request
	if unmarshalError != nil {
		beego.Error("Method: AddMaterial, Report: Error reading request : ", unmarshalError)
		common.Return400(this.Ctx.ResponseWriter)
		return
	}
	//No unmarshal error then validate all the input parameters
	valid := validation.Validation{}
	validated, _ := valid.Valid(material)
	//If there are validation errors then send a 211 error response
	if !validated{
		beego.Error("Method: AddMaterial, Report: Input Validation failed, cannot determine which field")
		common.Return417ErrorResponse(this.Ctx.ResponseWriter)
		return
	}
	//Now everything looks good call the model to insert the data in the table
	var materialDBObj common.Material
	json.Unmarshal(this.Ctx.Input.RequestBody, &materialDBObj)
	_, createError := models.CreateRecordInMaterialTable(materialDBObj)
	if createError.Status == 409 {
		beego.Error("Method: AddMaterial, Report: Create new record failed : ", createError)
		common.Return409Error(this.Ctx.ResponseWriter, createError)
		return
	}else if createError.Status == 400 {
		beego.Error("Method: AddMaterial, Report: Create new record failed : ", createError)
		common.Return400(this.Ctx.ResponseWriter)
		return
	}else{
		var resp common.ResposneStruct
		resp.Status = 201
		resp.Message = "New record created"
		response, _ := json.Marshal(resp)
		common.Return201(this.Ctx.ResponseWriter, response)
		return
	}
}

//Get the list of all the entries from the material table
func (this *MaterialController) GetMaterials(){
	result, error, count := models.QueryAllMaterials()
	if error == nil {
		if count == 0{
			common.EmptyResult(this.Ctx.ResponseWriter)
			return
		}
		resp, marshalErr := json.Marshal(result)
		if marshalErr == nil {
			common.Return200(this.Ctx.ResponseWriter, resp)
			return
		}else{
			common.Return210ErrorResponse(this.Ctx.ResponseWriter)
			return
		}
	}else{
		errResp, errMarshal := json.Marshal(error)
		if errMarshal == nil {
			common.Return200(this.Ctx.ResponseWriter, errResp)
			return
		}else{
			common.Return210ErrorResponse(this.Ctx.ResponseWriter)
			return
		}
	}
}

//Get a specific entry from the material table
func (this *MaterialController) GetMaterialById(){
	id := this.Ctx.Input.Param(":id")
	test := common.ValidateId(id)
	if !test {
		beego.Error("Method: DeleteMaterial, Report: Invalid ID format")
		common.Return400(this.Ctx.ResponseWriter)
		return
	}else{
		mid, _ := strconv.Atoi(id)
		result, selectError, count := models.QueryMaterialById(mid)
		if selectError == nil {
			if count == 0 {
				common.EmptyResult(this.Ctx.ResponseWriter)
				return
			}
			resp, marshalErr := json.Marshal(result)
			if marshalErr == nil {
				common.Return200(this.Ctx.ResponseWriter, resp)
			}
		}else{
			common.Return400(this.Ctx.ResponseWriter)
		}
	}
	return
}

func (this *MaterialController) ModifyMaterialById(){
	id := this.Ctx.Input.Param(":id")
	test := common.ValidateId(id)
	if !test {
		beego.Error("Method: ModifyMaterialById, Report: Invalid ID format")
		common.Return400(this.Ctx.ResponseWriter)
		return
	}else{
		//Read the input parameters and unmarshal the request body
		mid, _ := strconv.Atoi(id)
		var material common.MaterialObj
		unmarshalError := json.Unmarshal(this.Ctx.Input.RequestBody, &material)
		if unmarshalError != nil {
			beego.Error("Method: ModifyMaterialById, Report: Error reading request : ", unmarshalError)
			common.Return400(this.Ctx.ResponseWriter)
			return
		}
		//Validate the input fields before making the modifications to the DB
		valid := validation.Validation{}
		validated, _ := valid.Valid(&material)
		if !validated {
			beego.Error("Method: ModifyMaterialById, Report: Input Validation failed, cannot determine which field")
			common.Return417ErrorResponse(this.Ctx.ResponseWriter)
			return
		}
		//Call the model method to make the necessary changes in the DB
		_, modifyError := models.ModifyMaterialById(material, mid)
		if modifyError.Status != 200 {
			beego.Error("Method: ModifyMaterialById, Report: Modify request failed : ", modifyError)
			common.Return400(this.Ctx.ResponseWriter)
			return
		}else{
			response, _ := json.Marshal(modifyError)
			common.Return200(this.Ctx.ResponseWriter, response)
		}
	}
}

func (this *MaterialController) DeleteMaterial(){
	id := this.Ctx.Input.Param(":id")
	test := common.ValidateId(id)
	if !test {
		beego.Error("Method: DeleteMaterial, Report: Invalid ID format")
		common.Return400(this.Ctx.ResponseWriter)
		return
	}else{
		mid, _ := strconv.Atoi(id)
		result, deleteError := models.DeleteMaterialById(mid)
		if deleteError == nil {
			response, _ := json.Marshal(result)
			common.Return200(this.Ctx.ResponseWriter, response)
			return
		}else{
			beego.Error("Method: DeleteMaterial, Report: Failed to delete the record : ", deleteError)
			common.Return400(this.Ctx.ResponseWriter)
			return
		}
	}
}