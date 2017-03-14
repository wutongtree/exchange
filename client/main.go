package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	logging "github.com/op/go-logging"

	_ "github.com/wutongtree/exchange/client/bootstrap"
	_ "github.com/wutongtree/exchange/client/routers"
)

// var config
var (
	logger = logging.MustGetLogger("exchange.client")
)

func writeHyperledgerExplorer() {
	hyperledger_explorer := beego.AppConfig.String("hyperledger_explorer")
	filename := "static/explorer/hyperledger.js"
	fout, err := os.Create(filename)
	defer fout.Close()

	if err != nil {
		fmt.Printf("Write hyperledger exploer error: %v\n", err)
	} else {
		content := fmt.Sprintf("const REST_ENDPOINT = \"%v\";", hyperledger_explorer)
		fout.WriteString(content)
		fmt.Printf("Write hyperledger explorer with: %v\n", hyperledger_explorer)
	}
}

func main() {
	beego.SetStaticPath("/static", "static")
	beego.BConfig.WebConfig.DirectoryIndex = true
	// Write hyperledger explorer config
	writeHyperledgerExplorer()

	beego.InsertFilter("/*", beego.BeforeRouter, filterUser)
	beego.ErrorHandler("404", pageNotFound)
	beego.ErrorHandler("401", pageNoPermission)
	beego.Run()
}

var filterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("userLogin_exchange").(string)
	if !ok && ctx.Request.RequestURI != "/login" {
		ctx.Redirect(302, "/login")
	}
}

func pageNotFound(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("404.tpl").ParseFiles("views/404.tpl")
	data := make(map[string]interface{})
	t.Execute(rw, data)
}

func pageNoPermission(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("401.tpl").ParseFiles("views/401.tpl")
	data := make(map[string]interface{})
	t.Execute(rw, data)
}
