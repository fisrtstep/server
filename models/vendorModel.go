package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"firstStep/common"
)

func QueryAllVendors ()(resp []orm.Params, error interface{}, count int64){
	beego.Info("Model: QueryAllVendors")
	o := orm.NewOrm()
	var result []orm.Params
	resultCount, queryError := o.Raw("SELECT * FROM vendor").Values(&result)
	beego.Info("Result set : ", result)
	if queryError == nil && resultCount > 0 {
		return result, nil, resultCount
	}else {
		beego.Error("Method: QueryAllVendors, Report : Failed to query the Vendor table -- ", queryError)
		return nil, queryError, resultCount
	}
}

func QueryVendorById(vendorId int) (resp []orm.Params, error interface{}, count int64){
	o := orm.NewOrm()
	var result []orm.Params
	resultCount, queryError := o.Raw("SELECT * FROM vendor WHERE vendorId = ?", vendorId).Values(&result)
	beego.Info("Result set : ", result)
	if queryError == nil && resultCount > 0 {
		return result, nil, resultCount
	}else {
		beego.Info("Method: QueryVendorById, Report : Failed to query the Vendor table -- ", queryError)
		return nil, queryError, resultCount
	}
}

func CreateRecordInVendorTable(vendor common.Vendor) (count int64, error common.ResposneStruct){
	beego.Info("Model: CreateRecordInMaterialTable")
	var duplicate common.ResposneStruct
	var tempResult []orm.Params
	o := orm.NewOrm()
	dupCount, dupError := o.Raw("SELECT * FROM vendor WHERE cell = ? AND email = ? ",
		vendor.Cell, vendor.Email).Values(&tempResult)
	if dupCount == 0 && dupError == nil{
		result, insertError := o.Insert(&vendor)
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
		beego.Error("Method: QueryVendor, Report : Failed to query the Vendor table -- ", dupError)
		duplicate.Status = 400
		duplicate.Message = "Unable to fetch recodes from material table"
		return 0, duplicate
	}
}

func DeleteVendorById(vendorId int) (resp common.ResposneStruct, error interface{}) {
	o := orm.NewOrm()
	var response common.ResposneStruct
	var result []orm.Params
	_, deleteError := o.Raw("DELETE FROM vendor WHERE vendorId = ?", vendorId).Values(&result)
	if deleteError == nil {
		response.Status = 200
		response.Message = "Record deleted!"
		return response, nil
	}else{
		beego.Info("Method: DeleteVendorById, Report : Failed to delete record vendor table -- ", deleteError)
		response.Status = 400
		response.Message = "Record deleted!"
		return response, deleteError
	}
}

func ModifyVendorById(vendor common.VendorObj, vendorId int) (count int64, error common.ResposneStruct){
	beego.Info("Model : ModifyMaterialById")
	var resp common.ResposneStruct
	o := orm.NewOrm()
	var tempResult []orm.Params
	count, selectError := o.Raw("SELECT * FROM vendor WHERE vendorId = ?", vendorId).Values(&tempResult)
	if count != 1 {
		beego.Error("Method: ModifyVendorById, Report : Failed to query the Material table -- ", selectError)
		resp.Status = 400
		resp.Message = "Unable to fetch recodes from material table"
		return 0, resp
	}else{
		_, modifyError := o.Raw("UPDATE vendor SET name = ?, email = ?, cell = ?, phone = ?, address = ?, fax = ? " +
			"WHERE vendorId = ?", vendor.Name, vendor.Email, vendor.Cell, vendor.Phone, vendor.Address, vendor.Fax, vendorId).Exec()
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




