package models

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetTransferListTable(ctx *context.Context) table.Table {

	transferList := table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("mysql", "antiddos"))

	info := transferList.GetInfo().HideFilterArea().SetSortAsc()

	info.AddField("ID", "id", db.Bigint).FieldHide()
	info.AddField("IP地址", "ip_address", db.Varchar).FieldFilterable()
	info.AddField("端口号", "port_number", db.Smallint)
	info.AddField("密码", "password", db.Varchar)
	info.AddField("区域", "area", db.Varchar).FieldSortable()
	info.AddField("状态", "status", db.Tinyint).FieldSortable().FieldDisplay(func(value types.FieldModel) interface{} {
		switch value.Value {
		case "1":
			return "<span class='label label-success'>存活</span>"
		case "0":
			return "<span class='label label-danger'>死亡</span>"
		default:
			return "<span class='label label-default'>未知</span>"
		}
	}).FieldFilterable()
	info.AddField("池", "pool", db.Tinyint).FieldSortable().FieldDisplay(func(value types.FieldModel) interface{} {
		switch value.Value {
		case "1":
			return "<span class='label label-success'>可信</span>"
		case "2":
			return "<span class='label label-warning'>临时</span>"
		case "3":
			return "<span class='label label-danger'>备用</span>"
		default:
			return "<span class='label label-default'>未知</span>"
		}
	})

	//info.AddButton("导入", icon.SignIn, action.PopUpWithForm(action.PopUpData{
	//	Id:     "/ops/transfer/popup/upload",
	//	Title:  "导入中转机",
	//	Width:  "560px",
	//	Height: "280px",
	//}, func(panel *types.FormPanel) *types.FormPanel {
	//	panel.HideResetButton()
	//	panel.AddField("文件", "file", db.Varchar, form.File)
	//	panel.EnableAjax("成功", "失败")
	//	return panel
	//}, "/ops/transfer/popup/upload"))

	info.AddSelectBox("池", types.FieldOptions{
		{Value: "1", Text: "可信池"},
		{Value: "2", Text: "临时池"},
		{Value: "3", Text: "备用池"},
	}, action.FieldFilter("pool"))

	info.AddSelectBox("状态", types.FieldOptions{
		{Value: "0", Text: "死亡"},
		{Value: "1", Text: "存活"},
	}, action.FieldFilter("status"))

	info.SetTable("transfer_list").SetTitle("中转机列表")

	formList := transferList.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default).
		FieldHideWhenCreate().FieldHideWhenUpdate()
	formList.AddField("IP地址", "ip_address", db.Varchar, form.Ip).FieldMust()
	formList.AddField("端口号", "port_number", db.Smallint, form.Number)
	formList.AddField("密码", "password", db.Varchar, form.Password)
	formList.AddField("区域", "area", db.Varchar, form.SelectSingle).FieldOptions(types.FieldOptions{
		{Text: "华北", Value: "1"},
		{Text: "华东", Value: "2"},
		{Text: "华南", Value: "3"},
		{Text: "国内其他地区", Value: "4"},
		{Text: "海外", Value: "5"},
	})
	formList.AddField("状态", "status", db.Tinyint, form.SelectSingle).FieldOptions(types.FieldOptions{
		{Text: "死亡", Value: "0"},
		{Text: "存活", Value: "1"},
	})
	formList.AddField("池", "pool", db.Tinyint, form.SelectSingle).FieldOptions(types.FieldOptions{
		{Text: "可信池", Value: "1"},
		{Text: "临时池", Value: "2"},
		{Text: "备用池", Value: "3"},
	})

	formList.SetTable("transfer_list").SetTitle("中转机列表")

	return transferList
}
