package controllers

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	logging "github.com/op/go-logging"
)

// global const
// const (
// 	Website = "https://wutongtree.com"
// 	Email   = "hyper@crypto2x.com"
// )

// var
var (
	logger = logging.MustGetLogger("exchange.client.controllers")
)

type BaseController struct {
	beego.Controller
	IsLogin bool

	UserUserId   string
	UserUsername string
	UserAvatar   string
}

func (this *BaseController) Prepare() {
	userLogin_exchange := this.GetSession("userLogin_exchange")
	if userLogin_exchange == nil {
		this.IsLogin = false
		// this.Redirect("/login", 302)
		fmt.Printf("BaseController: login=false\n")
	} else {
		fmt.Printf("BaseController: login=true\n")

		this.IsLogin = true
		tmp := strings.Split((this.GetSession("userLogin_exchange")).(string), "||")

		this.Data["LoginUserid"] = tmp[0]
		this.Data["LoginUsername"] = tmp[1]
		this.Data["LoginAvatar"] = tmp[2]

		this.UserUserId = tmp[0]
		this.UserUsername = tmp[1]
		this.UserAvatar = tmp[2]
	}
	this.Data["IsLogin"] = this.IsLogin
}
