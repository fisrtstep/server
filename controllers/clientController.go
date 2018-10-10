package controllers

import (
	"github.com/astaxie/beego"
	"firstStep/models"
	"encoding/json"
	"firstStep/common"
	"strconv"
	"github.com/astaxie/beego/validation"
	"regexp"
)

type ClientController struct {
	beego.Controller
}

//Get the list of all the vendors that we buy from
func (this *ClientController) GetAllClients(){
	result, getError, count := models.QueryAllClients()
	if getError == nil {
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
		errResp, errMarshal := json.Marshal(getError)
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
func (this *ClientController) GetClientById(){
	id := this.Ctx.Input.Param(":id")
	test := common.ValidateId(id)
	if !test {
		beego.Error("Method: GetClientById, Report: Invalid ID format")
		common.Return400(this.Ctx.ResponseWriter)
		return
	}else{
		vendorId, _ := strconv.Atoi(id)
		result, selectError, count := models.QueryClientById(vendorId)
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

//Add a vendor to the DB
func (this ClientController) AddClient(){
	//Read the input parameters in the object, since it is unmarshaled only the required fields will be read
	var client common.ClientObj
	unmarshalError := json.Unmarshal(this.Ctx.Input.RequestBody, &client)
	//Check unmarshal errors, if there are errors then send 400 bad request
	if unmarshalError != nil {
		beego.Error("Method: AddClient, Report: Error reading request : ", unmarshalError)
		common.Return400(this.Ctx.ResponseWriter)
		return
	}
	//No unmarshal error then validate all the input parameters
	valid := validation.Validation{RequiredFirst:true}
	validated, _ := valid.Valid(&client)
	//If there are validation errors then send a 211 error response
	if !validated{
		beego.Error("Method: AddClient, Report: Input Validation failed, cannot determine which field")
		common.Return417ErrorResponse(this.Ctx.ResponseWriter)
		return
	}
	cellResult := valid.Match(client.Cell, regexp.MustCompile(`^(\+?[0-9]{10,14})((,\+?[0-9]{10,14})?){1,20}$`), "")
	if !cellResult.Ok {
		beego.Error("Method: AddClient, Report: Error reading request : Invalid cell phone num format")
		common.Return400(this.Ctx.ResponseWriter)
		return
	}
	if client.Phone != "" {
		phoneResult := valid.Match(client.Phone, regexp.MustCompile(`^^(\+?[0-9]{10,14})((,\+?[0-9]{10,14})?){1,20}$`), "")
		if !phoneResult.Ok {
			beego.Error("Method: AddClient, Report: Error reading request : Invalid phone num format")
			common.Return400(this.Ctx.ResponseWriter)
			return
		}
	}
	if client.Fax != ""{
		faxResult := valid.Match(client.Fax, regexp.MustCompile(`^^(\+?[0-9]{10,14})((,\+?[0-9]{10,14})?){1,20}$`), "")
		if !faxResult.Ok {
			beego.Error("Method: AddClient, Report: Error reading request : Invalid fax num format")
			common.Return400(this.Ctx.ResponseWriter)
			return
		}
	}
	//Now everything looks good call the model to insert the data in the table
	var clientDBObj common.Client
	json.Unmarshal(this.Ctx.Input.RequestBody, &clientDBObj)
	_, createError := models.CreateRecordInClientTable(clientDBObj)
	if createError.Status == 409 {
		beego.Error("Method: AddClient, Report: Create new record failed : ", createError)
		common.Return409Error(this.Ctx.ResponseWriter, createError)
		return
	}else if createError.Status == 400 {
		beego.Error("Method: AddVendor, Report: Create new record failed : ", createError)
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

func (this *ClientController) DeleteClient(){
	id := this.Ctx.Input.Param(":id")
	test := common.ValidateId(id)
	if !test {
		beego.Error("Method: DeleteClient, Report: Invalid ID format")
		common.Return400(this.Ctx.ResponseWriter)
		return
	}else{
		clientId, _ := strconv.Atoi(id)
		result, deleteError := models.DeleteClientById(clientId)
		if deleteError == nil {
			response, _ := json.Marshal(result)
			common.Return200(this.Ctx.ResponseWriter, response)
			return
		}else{
			beego.Error("Method: DeleteClient, Report: Failed to delete the record : ", deleteError)
			common.Return400(this.Ctx.ResponseWriter)
			return
		}
	}
}

func (this *ClientController) ModifyClientById(){
	id := this.Ctx.Input.Param(":id")
	test := common.ValidateId(id)
	if !test {
		beego.Error("Method: ModifyClientById, Report: Invalid ID format")
		common.Return400(this.Ctx.ResponseWriter)
		return
	}else{
		//Read the input parameters and unmarshal the request body
		clientId, _ := strconv.Atoi(id)
		var client common.ClientObj
		unmarshalError := json.Unmarshal(this.Ctx.Input.RequestBody, &client)
		if unmarshalError != nil {
			beego.Error("Method: ModifyClientById, Report: Error reading request : ", unmarshalError)
			common.Return400(this.Ctx.ResponseWriter)
			return
		}
		//Validate the input fields before making the modifications to the DB
		valid := validation.Validation{}
		validated, _ := valid.Valid(&client)
		if !validated {
			beego.Error("Method: ModifyClientById, Report: Input Validation failed, cannot determine which field")
			common.Return417ErrorResponse(this.Ctx.ResponseWriter)
			return
		}
		cellResult := valid.Match(client.Cell, regexp.MustCompile(`^^(\+?[0-9]{10,14})((,\+?[0-9]{10,14})?){1,20}$`), "")
		if !cellResult.Ok {
			beego.Error("Method: ModifyClientById, Report: Error reading request : Invalid cell phone num format")
			common.Return400(this.Ctx.ResponseWriter)
			return
		}
		if client.Phone != "" {
			phoneResult := valid.Match(client.Phone, regexp.MustCompile(`^^(\+?[0-9]{10,14})((,\+?[0-9]{10,14})?){1,20}$`), "")
			if !phoneResult.Ok {
				beego.Error("Method: ModifyClientById, Report: Error reading request : Invalid phone num format")
				common.Return400(this.Ctx.ResponseWriter)
				return
			}
		}
		if client.Fax != ""{
			faxResult := valid.Match(client.Fax, regexp.MustCompile(`^^(\+?[0-9]{10,14})((,\+?[0-9]{10,14})?){1,20}$`), "")
			if !faxResult.Ok {
				beego.Error("Method: ModifyClientById, Report: Error reading request : Invalid fax num format")
				common.Return400(this.Ctx.ResponseWriter)
				return
			}
		}
		//Call the model method to make the necessary changes in the DB
		_, modifyError := models.ModifyClientById(client, clientId)
		if modifyError.Status != 200 {
			beego.Error("Method: ModifyVendorById, Report: Modify request failed : ", modifyError)
			common.Return400(this.Ctx.ResponseWriter)
			return
		}else{
			response, _ := json.Marshal(modifyError)
			common.Return200(this.Ctx.ResponseWriter, response)
		}
	}
}
