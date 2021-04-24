package models

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetSysConfigTable(ctx *context.Context) table.Table {

	sysConfig := table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("mysql", "antiddos"))

	info := sysConfig.GetInfo().HideFilterArea().SetSortAsc().
		HideNewButton()

	info.AddField("Id", "id", db.Bigint).FieldHide()
	info.AddField("参数项", "key", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		switch value.Value {
		case "replay_attack_time":
			return "重放判定间隔"
		case "distribute_cycle_time":
			return "分配周期"
		case "transfer_gc_interval":
			return "中转机GC扫描间隔"
		case "cache_expire_time":
			return "缓存过期时间"
		case "distribute_count":
			return "中转机组成员数"
		case "back_pool_threshold":
			return "备用池告警阈值"
		default:
			return "unknown"
		}
	})
	info.AddField("参数值", "value", db.Varchar).FieldEditAble()

	info.SetTable("sys_config").SetTitle("系统配置")

	formList := sysConfig.GetForm()
	formList.AddField("参数项", "key", db.Varchar, form.Text).FieldDisableWhenUpdate()
	formList.AddField("参数值", "value", db.Varchar, form.Text).FieldOptions(types.FieldOptions{
		{Text: "true", Value: "true"},
		{Text: "false", Value: "false"},
	})

	formList.SetTable("sys_config").SetTitle("系统配置")

	return sysConfig
}
