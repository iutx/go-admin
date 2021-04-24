package models

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetRequestRecordTable(ctx *context.Context) table.Table {

	requestRecord := table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("mysql", "antiddos"))

	info := requestRecord.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).FieldHide()
	info.AddField("设备编号", "device_id", db.Varchar)
	info.AddField("应用提供商编号", "app_provider_id", db.Bigint)
	info.AddField("请求时间", "request_time", db.Datetime)
	info.AddField("客户端IP", "real_ip", db.Varchar)

	info.SetTable("request_record").SetTitle("请求记录")

	formList := requestRecord.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("设备编号", "device_id", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("应用提供商编号", "app_provider_id", db.Bigint, form.Number).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("请求时间", "request_time", db.Datetime, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("客户端IP", "real_ip", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()

	formList.SetTable("request_record").SetTitle("请求记录")

	return requestRecord
}
