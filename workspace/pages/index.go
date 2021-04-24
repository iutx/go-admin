package pages

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/workspace/models"
	"html/template"
	"strconv"
	"time"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	goadmintemplate "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/themes/sword/components/progress_group"
)

const (
	QueryStartTime = "start_time"
	QueryEndTime   = "end_time"
	actionSVG      = `
	<i aria-label="图标: info-circle-o" class="anticon anticon-info-circle-o">
		<svg viewBox="64 64 896 896" focusable="false" class="" data-icon="info-circle" width="1em" height="1em" fill="currentColor" aria-hidden="true">
			<path d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448 448-200.6 448-448S759.4 64 512 64zm0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372 372 166.6 372 372-166.6 372-372 372z"></path><path d="M464 336a48 48 0 1 0 96 0 48 48 0 1 0-96 0zm72 112h-48c-4.4 0-8 3.6-8 8v272c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V456c0-4.4-3.6-8-8-8z">
			</path>
		</svg>
	</i>`
)

var (
	rgbMap = map[string]string{
		"2": "rgb(215,231,232,222)",
		"3": "rgb(0,23,123,34)",
	}
)

type ShowInfo struct {
	transferInfo  *TransferInfo
	providersInfo *[]ProviderInfo
}

type TransferInfo struct {
	total     int
	trustPool PoolInfo
	tempPool  PoolInfo
	backPool  PoolInfo
}

type PoolInfo struct {
	name         template.HTML
	color        template.HTML
	Total        int `gorm:"column:total"`
	HealthyCount int `gorm:"column:healthy_count"`
	DeathCount   int `gorm:"column:death_count"`
}

type ProviderInfo struct {
	models.ProviderInfo
}

func GetDashBoard(ctx *context.Context) (types.Panel, error) {

	ti := *GetShowInfo().transferInfo
	psi := GetShowInfo().providersInfo

	queryInfo := ctx.Request.URL.Query()

	components := goadmintemplate.Get(config.GetTheme())
	colComp := components.Col()

	transfersPool := colComp.
		SetContent(GetProcessGroup(ti)).
		SetSize(types.SizeMD(4)).
		GetContent()

	dailyRequest := colComp.SetContent(GetLineChart(queryInfo.Get(QueryStartTime), queryInfo.Get(QueryEndTime), psi)).SetSize(types.SizeMD(12)).GetContent()
	boxInternalRow := components.Row().SetContent(dailyRequest).GetContent()

	box := components.Box().WithHeadBorder().SetHeader("请求数据").
		SetBody(boxInternalRow).
		GetContent()

	boxcol := colComp.SetContent(box).SetSize(types.SizeMD(12)).GetContent()
	row2 := components.Row().SetContent(boxcol).GetContent()

	boxPool := components.Box().WithHeadBorder().SetHeader("中转机池状态").
		SetBody(transfersPool).
		GetContent()

	poolCol := colComp.SetContent(boxPool).SetSize(types.SizeMD(12)).GetContent()
	row3 := components.Row().SetContent(poolCol).GetContent()

	return types.Panel{
		Content: row3 + row2,
		Title:   "仪表盘",
	}, nil
}

func GetShowInfo() ShowInfo {
	showInfo := ShowInfo{
		transferInfo: &TransferInfo{
			trustPool: PoolInfo{
				name:  "可信池",
				color: "#ace0ae",
			},
			tempPool: PoolInfo{
				name:  "临时池",
				color: "#fdd698",
			},
			backPool: PoolInfo{
				name:  "备用池",
				color: "#f17c6e",
			},
		},
		providersInfo: &[]ProviderInfo{},
	}

	tlCon := models.ORM.Table("transfer_list")
	if err := tlCon.Count(&showInfo.transferInfo.total).Error; err != nil {
		logger.Error(err)
	}

	queryParam := "COUNT(1) as total, COUNT(case when status=1 then 1 else null end) as healthy_count, COUNT(case when status=0 then 1 else null end) as death_count"

	if err := tlCon.Select(queryParam).Where("pool = 1").Scan(&showInfo.transferInfo.trustPool).Error; err != nil {
		logger.Error(err)
	}

	if err := tlCon.Select(queryParam).Where("pool = 2").Scan(&showInfo.transferInfo.tempPool).Error; err != nil {
		logger.Error(err)
	}

	if err := tlCon.Select(queryParam).Where("pool = 3").Scan(&showInfo.transferInfo.backPool).Error; err != nil {
		logger.Error(err)
	}

	if err := models.ORM.Table("provider_info").Scan(&showInfo.providersInfo).Error; err != nil {
		logger.Error(err)
	}

	return showInfo
}

func convertToIntPercent(info PoolInfo) int {
	res, err := strconv.Atoi(fmt.Sprintf("%.0f", float64(info.HealthyCount)/float64(info.Total)*100))
	if err != nil {
		logger.Error(err)
	}
	return res
}

func GetProcessGroup(ti TransferInfo) template.HTML {
	title := `<p class="text-center"><strong>中转机池详情</strong></p>`

	tempPoolPG := progress_group.New().
		SetTitle(ti.tempPool.name).
		SetColor(ti.tempPool.color).
		SetDenominator(ti.tempPool.Total).
		SetMolecular(ti.tempPool.HealthyCount).
		SetPercent(convertToIntPercent(ti.tempPool)).
		GetContent()

	trustPoolPG := progress_group.New().
		SetTitle(ti.trustPool.name).
		SetColor(ti.trustPool.color).
		SetDenominator(ti.trustPool.Total).
		SetMolecular(ti.trustPool.HealthyCount).
		SetPercent(convertToIntPercent(ti.trustPool)).
		GetContent()

	backPoolPG := progress_group.New().
		SetTitle(ti.backPool.name).
		SetColor(ti.backPool.color).
		SetDenominator(ti.backPool.Total).
		SetMolecular(ti.backPool.HealthyCount).
		SetPercent(convertToIntPercent(ti.backPool)).
		GetContent()
	return template.HTML(title) + tempPoolPG + trustPoolPG + backPoolPG
}

func GetLineChart(inStartTime, inEndTime string, psi *[]ProviderInfo) template.HTML {
	layoutDay := "2006-01-02"

	nt := time.Now()
	endTime := nt
	startTime := nt.AddDate(0, 0, -6)

	if inStartTime != "" && inEndTime == "" {
		if tmpTime, err := time.Parse(inStartTime, layoutDay); err == nil {
			startTime = tmpTime
		}
		if tmpTime, err := time.Parse(inEndTime, layoutDay); err == nil {
			endTime = tmpTime
		}
	}

	changeTime := startTime

	labels := make([]string, 0)

	for !changeTime.After(endTime) {
		labels = append(labels, changeTime.Format("01-02"))
		changeTime = changeTime.AddDate(0, 0, 1)
	}

	charObj := chartjs.Line().
		SetID("requestChart").
		SetHeight(180).
		SetTitle(template.HTML(fmt.Sprintf("%v - %v", startTime.Format(layoutDay), endTime.Format(layoutDay)))).
		SetLabels(labels)

	for _, provider := range *psi {
		charObj = charObj.AddDataSet(provider.ProviderName).
			DSData([]float64{65, 59, 80, 81, 56, 55, 40}).
			DSFill(false).
			DSBorderColor(chartjs.Color(rgbMap[provider.AppProviderID])).
			DSLineTension(0.1)
	}

	return charObj.GetContent()
}
