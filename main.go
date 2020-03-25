package main

import (
	"fmt"
	"github.com/droundy/goopt"
	"strings"
	"gtool/core"
	"gtool/templates"
)

var (
	sqlTable    = goopt.String([]string{"-t", "--table"}, "", "Table to build struct from")
	sqlDatabase = goopt.String([]string{"-d", "--database"}, "", "database, database name")
	routerName  = goopt.String([]string{"-r", "--router"}, "", "routerName, if have this conditon, will create router file")
	controller  = goopt.String([]string{"-c", "--controller"}, "false", "controller, if \"true\" will create controller file.")
	fileDir     = goopt.String([]string{"-f", "--file-path"}, "", "file dir,if not use 'output' ")
)

func init() {
	goopt.Description = func() string {
		return "Gtool is automaticlly generate clue-api model"
	}
	goopt.Version = "0.1"
	goopt.Summary = `Gtool --table tableName --database dbName --router routeName --controller true --file-path ./`
	goopt.Parse(nil)
}

func main() {
	if *sqlTable == "" {
		fmt.Println("table can not is empty")
		return
	}
	if *sqlDatabase == "" {
		fmt.Println("database can not is empty")
		return
	}
	if *routerName != "" {
		generateRoute()
	}
	if *controller == "true" {
		generateController()
	}
	generateModel()
	generateService()
}

func generateController()  {
	var iCore core.ICore
	var ce core.Core
	ce.Name = core.LcFirst(*sqlTable)
	ce.FileDir = *fileDir
	ce.Format = map[string]string{
		"{{StructFcName}}" 	: ce.Name,
		"{{StructName}}"	: core.UcFirst(*sqlTable),
		"{{StructDB}}"		: core.UcFirst(*sqlDatabase),
		"{{StructRoute}}" 	: *routerName,
		"{{StrcutF}}" 		: strings.ToLower(ce.Name[:1]),
	}
	iCore = ce
	iCore.GenerateFile(templates.ControllerTpl, "controllers")
}

func generateRoute() {
	var iCore core.ICore
	var ce core.Core
	ce.FileDir = *fileDir
	ce.Name = core.LcFirst(*sqlTable)
	ce.Format = map[string]string{
		"{{StructFcName}}" 	: ce.Name,
		"{{StructName}}"	: core.UcFirst(*sqlTable),
		"{{StructDB}}"		: core.UcFirst(*sqlDatabase),
		"{{StructRoute}}" 	: *routerName,
	}
	iCore = ce
	iCore.GenerateFile(templates.RouterTpl, "router")
}

func generateService() {
	var iCore core.ICore
	var ce core.Core
	ce.FileDir = *fileDir
	ce.Name = core.LcFirst(*sqlTable)
	ce.Format = map[string]string{
		"{{StructLcName}}": ce.Name,
		"{{StructName}}": core.UcFirst(*sqlTable),
		"{{StructDB}}"	: core.UcFirst(*sqlDatabase),
	}
	iCore = ce
	iCore.GenerateFile(templates.ServiceTpl, "service")
}

func generateModel() {
	var iCore core.ICore
	var ce core.Core
	ce.FileDir = *fileDir
	ce.Name = core.LcFirst(*sqlTable)
	ce.Format = map[string]string{
		"{{StructLcName}}": ce.Name,
		"{{StructName}}": core.UcFirst(*sqlTable),
		"{{StructDB}}"	: core.UcFirst(*sqlDatabase),
	}
	iCore = ce
	iCore.GenerateFile(templates.ModelTpl, "model")
}