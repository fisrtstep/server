package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"firstStep/common"
)

func CreateRecordInMaterialTable(material common.Material) (count int64, error common.ResposneStruct){
	beego.Info("Model: CreateRecordInMaterialTable")
	var duplicate common.ResposneStruct
	var tempResult []orm.Params
	o := orm.NewOrm()
	dupCount, dupError := o.Raw("SELECT * FROM material WHERE materialName = ? AND materialGrade = ? " +
		"AND materialType = ?", material.MaterialName, material.MaterialGrade, material.MaterialType).Values(&tempResult)
	if dupCount == 0 && dupError == nil{
		result, insertError := o.Insert(&material)
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
		beego.Error("Method: QueryAllMaterials, Report : Failed to query the Material table -- ", dupError)
		duplicate.Status = 400
		duplicate.Message = "Unable to fetch recodes from material table"
		return 0, duplicate
	}
}

func ModifyMaterialById(material common.MaterialObj, mid int) (count int64, error common.ResposneStruct){
	beego.Info("Model : ModifyMaterialById")
	var resp common.ResposneStruct
	o := orm.NewOrm()
	var tempResult []orm.Params
	count, selectError := o.Raw("SELECT * FROM material WHERE mid = ?", mid).Values(&tempResult)
	if count != 1 {
		beego.Error("Method: QueryAllMaterials, Report : Failed to query the Material table -- ", selectError)
		resp.Status = 400
		resp.Message = "Unable to fetch recodes from material table"
		return 0, resp
	}else{
		_, modifyError := o.Raw("UPDATE material SET materialName = ?, materialType = ?, materialGrade = ? " +
			"WHERE mid = ?", material.MaterialName, material.MaterialType, material.MaterialGrade, mid).Exec()
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

func QueryAllMaterials ()(resp []orm.Params, error interface{}, count int64){
	beego.Info("Model: QueryAllMaterials")
	o := orm.NewOrm()
	var result []orm.Params
	resultCount, queryError := o.Raw("SELECT * FROM material").Values(&result)
	beego.Info("Result set : ", result)
	if queryError == nil && resultCount > 0 {
		return result, nil, resultCount
	}else {
		beego.Error("Method: QueryAllMaterials, Report : Failed to query the Material table -- ", queryError)
		return nil, queryError, resultCount
	}
}

func QueryMaterialById(mid int) (resp []orm.Params, error interface{}, count int64){
	o := orm.NewOrm()
	var result []orm.Params
	resultCount, queryError := o.Raw("SELECT * FROM material WHERE mid = ?", mid).Values(&result)
	beego.Info("Result set : ", result)
	if queryError == nil && resultCount > 0 {
		return result, nil, resultCount
	}else {
		beego.Info("Method: QueryMaterialById, Report : Failed to query the Material table -- ", queryError)
		return nil, queryError, resultCount
	}
}

func DeleteMaterialById(mid int) (resp common.ResposneStruct, error interface{}) {
	o := orm.NewOrm()
	var response common.ResposneStruct
	var result []orm.Params
	_, deleteError := o.Raw("DELETE FROM material WHERE mid = ?", mid).Values(&result)
	if deleteError == nil {
		response.Status = 200
		response.Message = "Record deleted!"
		return response, nil
	}else{
		beego.Info("Method: DeleteMaterialById, Report : Failed to delete record material table -- ", deleteError)
		response.Status = 400
		response.Message = "Record deleted!"
		return response, deleteError
	}
}
