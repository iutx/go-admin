package models

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetClientWeightTable(ctx *context.Context) table.Table {

	clientWeight := table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("mysql", "antiddos"))

	info := clientWeight.GetInfo().HideFilterArea().SetSortAsc()

	info.AddField("Id", "id", db.Bigint).
		FieldHide()
	info.AddField("应用编号", "app_provider_id", db.Varchar).FieldFilterable()
	info.AddField("设备标识", "device_id", db.Varchar).FieldFilterable()
	info.AddField("权重值", "weight", db.Smallint).FieldEditAble()
	info.AddField("创建时间", "created_at", db.Datetime)
	info.AddField("更新时间", "updated_at", db.Datetime)

	info.SetTable("client_weight").SetTitle("设备权重")

	formList := clientWeight.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("应用编号", "app_provider_id", db.Varchar, form.Text).
		FieldDisableWhenUpdate()
	formList.AddField("设备标识", "device_id", db.Varchar, form.Text).
		FieldDisableWhenUpdate()
	formList.AddField("权重值", "weight", db.Smallint, form.Number)
	formList.AddField("创建时间", "created_at", db.Datetime, form.Datetime).
		FieldDisableWhenUpdate().
		FieldHide().FieldNowWhenInsert()
	formList.AddField("更新时间", "updated_at", db.Datetime, form.Datetime).
		FieldDisableWhenUpdate().
		FieldHide().FieldNowWhenUpdate()

	formList.SetTable("client_weight").SetTitle("设备权重")

	return clientWeight
}
