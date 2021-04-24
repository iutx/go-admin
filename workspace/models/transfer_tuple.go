package models

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetTransferTupleTable(ctx *context.Context) table.Table {

	transferTuple := table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("mysql", "antiddos"))

	info := transferTuple.GetInfo().HideFilterArea().SetSortField("distribute_time")

	info.AddField("Id", "id", db.Bigint)
	info.AddField("机组ID", "tuple_id", db.Longtext)
	info.AddField("分配时间", "distribute_time", db.Datetime)
	info.AddField("设备ID", "device_id", db.Varchar)
	info.AddField("应用ID", "app_provider_id", db.Varchar)
	info.AddField("中转机信息", "tuple_transfers", db.Longtext)

	info.SetTable("transfer_tuple").SetTitle("中转机组")

	formList := transferTuple.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default).
		FieldDisableWhenCreate()
	formList.AddField("机组ID", "tuple_id", db.Longtext, form.RichText)
	formList.AddField("分配时间", "distribute_time", db.Datetime, form.Datetime)
	formList.AddField("设备ID", "device_id", db.Varchar, form.Text)
	formList.AddField("应用ID", "app_provider_id", db.Varchar, form.Text)
	formList.AddField("中转机信息", "tuple_transfers", db.Longtext, form.RichText)

	formList.SetTable("transfer_tuple").SetTitle("中转机组")

	return transferTuple
}
