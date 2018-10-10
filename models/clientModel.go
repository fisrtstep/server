package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"firstStep/common"
)

func QueryAllClients ()(resp []orm.Params, error interface{}, count int64){
	beego.Info("Model: QueryAllClients")
	o := orm.NewOrm()
	var result []orm.Params
	resultCount, queryError := o.Raw("SELECT * FROM client").Values(&result)
	beego.Info("Result set : ", result)
	if queryError == nil && resultCount > 0 {
		return result, nil, resultCount
	}else {
		beego.Error("Method: QueryAllClients, Report : Failed to query the client table -- ", queryError)
		return nil, queryError, resultCount
	}
}

func QueryClientById(clientId int) (resp []orm.Params, error interface{}, count int64){
	o := orm.NewOrm()
	var result []orm.Params
	resultCount, queryError := o.Raw("SELECT * FROM client WHERE clientId = ?", clientId).Values(&result)
	beego.Info("Result set : ", result)
	if queryError == nil && resultCount > 0 {
		return result, nil, resultCount
	}else {
		beego.Info("Method: QueryClientById, Report : Failed to query the Client table -- ", queryError)
		return nil, queryError, resultCount
	}
}

func CreateRecordInClientTable(client common.Client) (count int64, error common.ResposneStruct){
	beego.Info("Model: CreateRecordInClientTable")
	var duplicate common.ResposneStruct
	var tempResult []orm.Params
	o := orm.NewOrm()
	dupCount, dupError := o.Raw("SELECT * FROM client WHERE cell = ? AND email = ? ",
		client.Cell, client.Email).Values(&tempResult)
	if dupCount == 0 && dupError == nil{
		result, insertError := o.Insert(&client)
		if insertError == nil {
			return result, duplicate
		}else{
			duplicate.Status = 400
			duplicate.Message = insertError.Error()
			return 0, duplicate
		}
	}else if dupCount != 0 {
		duplicate.Status = 409
		duplicate.Message = "Duplicate entry, please check the existing entries and try again"
		return 0, duplicate
	}else{
		beego.Error("Method: CreateRecordInClientTable, Report : Failed to query the Client table -- ", dupError)
		duplicate.Status = 400
		duplicate.Message = "Unable to fetch recodes from Client table"
		return 0, duplicate
	}
}

func DeleteClientById(clientId int) (resp common.ResposneStruct, error interface{}) {
	o := orm.NewOrm()
	var response common.ResposneStruct
	var result []orm.Params
	_, deleteError := o.Raw("DELETE FROM vendor WHERE vendorId = ?", clientId).Values(&result)
	if deleteError == nil {
		response.Status = 200
		response.Message = "Record deleted!"
		return response, nil
	}else{
		beego.Info("Method: DeleteClientById, Report : Failed to delete record client table -- ", deleteError)
		response.Status = 400
		response.Message = "Record deleted!"
		return response, deleteError
	}
}

func ModifyClientById(client common.ClientObj, clientId int) (count int64, error common.ResposneStruct){
	beego.Info("Model : ModifyMaterialById")
	var resp common.ResposneStruct
	o := orm.NewOrm()
	var tempResult []orm.Params
	count, selectError := o.Raw("SELECT * FROM client WHERE clientId = ?", clientId).Values(&tempResult)
	if count != 1 {
		beego.Error("Method: ModifyClientById, Report : Failed to query the Material table -- ", selectError)
		resp.Status = 400
		resp.Message = "Unable to fetch recodes from client table"
		return 0, resp
	}else{
		_, modifyError := o.Raw("UPDATE client SET name = ?, email = ?, cell = ?, phone = ?, address = ?, fax = ? " +
			"WHERE clientId = ?", client.Name, client.Email, client.Cell, client.Phone, client.Address, client.Fax, clientId).Exec()
		if modifyError == nil {
			resp.Status = 200
			resp.Message = "Record modified"
			return 1, resp
		}else{
			resp.Status = 400
			resp.Message = "Modify failed"
			return 0, resp
		}

	}
}




