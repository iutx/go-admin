package tables

import (
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/workspace/models"
)

var Generators = map[string]table.Generator{
	"sys_config":     models.GetSysConfigTable,
	"provider_info":  models.GetProviderInfoTable,
	"request_record": models.GetRequestRecordTable,
	"transfer_list":  models.GetTransferListTable,
	"verify_server":  models.GetVerifyServerTable,
	"client_weight":  models.GetClientWeightTable,
	"transfer_tuple": models.GetTransferTupleTable,
}
