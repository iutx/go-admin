package models

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetProviderInfoTable(ctx *context.Context) table.Table {

	providerInfo := table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("mysql", "antiddos"))

	info := providerInfo.GetInfo().HideFilterArea().SetSortField("id")

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("应用提供商编号", "app_provider_id", db.Varchar)
	info.AddField("授权码", "app_auth_code", db.Varchar)
	info.AddField("服务商名称", "provider_name", db.Varchar)
	info.AddField("联系地址", "contact_address", db.Varchar)
	info.AddField("联系人", "contacts", db.Varchar)
	info.AddField("联系方式", "phone", db.Varchar)
	info.AddField("联系邮箱", "email", db.Varchar)
	info.AddField("SDK开关", "is_open_sdk", db.Tinyint)
	info.AddField("高防列表", "high_defense", db.Varchar)
	info.AddField("授权起始时间", "authorized_time", db.Datetime)
	info.AddField("授权结束时间", "expire_time", db.Datetime)

	info.SetTable("provider_info").SetTitle("服务商信息").SetDescription("服务商信息")

	formList := providerInfo.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("应用提供商编号", "app_provider_id", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("授权码", "app_auth_code", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("服务商名称", "provider_name", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("联系地址", "contact_address", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("联系人", "contacts", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("联系方式", "phone", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("联系邮箱", "email", db.Varchar, form.Email).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("SDK开关", "is_open_sdk", db.Tinyint, form.Number).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("高防列表", "high_defense", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("授权起始时间", "authorized_time", db.Datetime, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("授权结束时间", "expire_time", db.Datetime, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()

	formList.SetTable("provider_info").SetTitle("服务商信息").SetDescription("服务商信息")

	return providerInfo
}
