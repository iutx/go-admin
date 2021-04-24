package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin" // web framework adapter
	"github.com/GoAdminGroup/go-admin/engine"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql" // sql driver
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/go-admin/workspace/models"
	"github.com/GoAdminGroup/go-admin/workspace/pages"
	"github.com/GoAdminGroup/go-admin/workspace/tables"
	_ "github.com/GoAdminGroup/themes/sword" // ui theme
)

func main() {
	startServer()
}

func startServer() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	r := gin.Default()

	template.AddComp(chartjs.NewChart())
	//template.AddLoginComp(login.Get())

	eng := engine.Default()

	if err := eng.AddConfigFromJSON("./workspace/config.json").
		AddGenerators(tables.Generators).
		Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", "./uploads")

	eng.HTML("GET", "/ops", pages.GetDashBoard)

	models.Init(eng.MysqlConnection())

	_ = r.Run(":80")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	eng.MysqlConnection().Close()
}
