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

type MainController struct {
	beego.Controller
}

//Create a new record in the material table
func (m *MainController) AddMaterial(){
	//Read the input parameters in the object, since it is unmarshaled only the required fields will be read
	var material common.MaterialObj
	unmarshalError := json.Unmarshal(m.Ctx.Input.RequestBody, &material)
	//Check unmarshal errors, if there are errors then send 400 bad request
	if unmarshalError != nil {
		beego.Error("Method: AddMaterial, Report: Error reading request : ", unmarshalError)
		common.Return400(m.Ctx.ResponseWriter)
		return
	}
	//No unmarshal error then validate all the input parameters
	valid := validation.Validation{}
	validated, _ := valid.Valid(material)
	//If there are validation errors then send a 211 error response
	if !validated{
		beego.Error("Method: AddMaterial, Report: Input Validation failed, cannot determine which field")
		common.Return417ErrorResponse(m.Ctx.ResponseWriter)
		return
	}
	//Now everything looks good call the model to insert the data in the table
	var materialDBObj common.Material
	json.Unmarshal(m.Ctx.Input.RequestBody, &materialDBObj)
	_, createError := models.CreateRecordInMaterialTable(materialDBObj)
	if createError.Status == 409 {
		beego.Error("Method: AddMaterial, Report: Create new record failed : ", createError)
		common.Return409Error(m.Ctx.ResponseWriter, createError)
		return
	}else if createError.Status == 400 {
		beego.Error("Method: AddMaterial, Report: Create new record failed : ", createError)
		common.Return400(m.Ctx.ResponseWriter)
		return
	}else{
		var resp common.ResposneStruct
		resp.Status = 201
		resp.Message = "New record created"
		response, _ := json.Marshal(resp)
		common.Return201(m.Ctx.ResponseWriter, response)
		return
	}
}

//Get the list of all the entries from the material table
func (m *MainController) GetMaterials(){
	// TODO: Make the sanity check and then call model
	result, error, count := models.QueryAllMaterials()
	if error == nil {
		if count == 0{
			common.EmptyResult(m.Ctx.ResponseWriter)
			return
		}
		resp, marshleErr := json.Marshal(result)
		if marshleErr == nil {
			common.Return200(m.Ctx.ResponseWriter, resp)
		}else{
			common.Return210ErrorResponse(m.Ctx.ResponseWriter)
		}
	}else{
		errResp, errMarshal := json.Marshal(error)
		if errMarshal == nil {
			common.Return200(m.Ctx.ResponseWriter, errResp)
		}else{
			common.Return210ErrorResponse(m.Ctx.ResponseWriter)
		}
	}
	return
}

//Get a specific entry from the material table
func (m *MainController) GetMaterialById(){
	id := m.Ctx.Input.Param(":id")
	//TODO: Check the ID with a regular expression and report error if validation fails
	if id == "" {
		common.Return400(m.Ctx.ResponseWriter)
	}else{
		mid, _ := strconv.Atoi(id)
		result, selectError, count := models.QueryMaterialById(mid)
		if selectError == nil {
			if count == 0 {
				common.EmptyResult(m.Ctx.ResponseWriter)
				return
			}
			resp, marshalErr := json.Marshal(result)
			if marshalErr == nil {
				common.Return200(m.Ctx.ResponseWriter, resp)
			}
		}else{
			common.Return400(m.Ctx.ResponseWriter)
		}
	}
	return
}

func (m *MainController) ModifyMaterialById(){
	id := m.Ctx.Input.Param(":id")
	//TODO: Check the ID with a regular expression and report error if validation fails
	if id == "" {
		common.Return400(m.Ctx.ResponseWriter)
	}else{
		//Read the input parameters and unmarshal the request body
		mid, _ := strconv.Atoi(id)
		var material common.MaterialObj
		unmarshalError := json.Unmarshal(m.Ctx.Input.RequestBody, &material)
		if unmarshalError != nil {
			beego.Error("Method: ModifyMaterialById, Report: Error reading request : ", unmarshalError)
			common.Return400(m.Ctx.ResponseWriter)
			return
		}
		//Validate the input fields before making the modifications to the DB
		valid := validation.Validation{}
		validated, _ := valid.Valid(&material)
		if !validated {
			beego.Error("Method: ModifyMaterialById, Report: Input Validation failed, cannot determine which field")
			common.Return417ErrorResponse(m.Ctx.ResponseWriter)
			return
		}
		//Call the model method to make the necessary changes in the DB
		_, modifyError := models.ModifyMaterialById(material, mid)
		if modifyError.Status != 200 {
			beego.Error("Method: ModifyMaterialById, Report: Modify request failed : ", modifyError)
			common.Return400(m.Ctx.ResponseWriter)
			return
		}else{
			resposne, _ := json.Marshal(modifyError)
			common.Return200(m.Ctx.ResponseWriter, resposne)
		}
	}
}