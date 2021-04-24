package models

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetVerifyServerTable(ctx *context.Context) table.Table {

	verifyServer := table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("mysql", "antiddos"))

	info := verifyServer.GetInfo().HideFilterArea().SetSortAsc()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("应用服务商编号", "app_provider_id", db.Varchar)
	info.AddField("验证端地址", "verify_address", db.Varchar)
	info.AddField("附加地址", "additional_address", db.Varchar)
	info.AddField("SDK版本", "sdk_version", db.Float)
	info.AddField("权重", "weight", db.Tinyint)

	info.SetTable("verify_server").SetTitle("验证端列表").SetDescription("验证端列表")

	formList := verifyServer.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("应用服务商编号", "app_provider_id", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("验证端地址", "verify_address", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("附加地址", "additional_address", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("SDK版本", "sdk_version", db.Float, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("权重", "weight", db.Tinyint, form.Number).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()

	formList.SetTable("verify_server").SetTitle("验证端列表").SetDescription("验证端列表")

	return verifyServer
}
