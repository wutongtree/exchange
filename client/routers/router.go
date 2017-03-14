package routers

import (
	"github.com/astaxie/beego"
	"github.com/wutongtree/exchange/client/controllers"
)

func init() {
	// 登录
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")

	beego.Router("/my/index", &controllers.CurrencyController{}, "get:Index")

	beego.Router("/currency/create", &controllers.CurrencyController{}, "post:CreatePost;get:CreateGet")
	beego.Router("/currency/create/check/:txid(.*)", &controllers.CurrencyController{}, "get:CreateCheck")

	beego.Router("/currency/release", &controllers.CurrencyController{}, "post:ReleasePost")
	beego.Router("/currency/release/:id(.*)", &controllers.CurrencyController{}, "get:ReleaseGet")
	beego.Router("/currency/release/check/:txid(.*)", &controllers.CurrencyController{}, "get:CreateCheck")

	beego.Router("/currency/assign", &controllers.CurrencyController{}, "post:AssignPost")
	beego.Router("/currency/assign/:id(.*)", &controllers.CurrencyController{}, "get:AssignGet")
	beego.Router("/currency/assign/check/:txid(.*)", &controllers.CurrencyController{}, "get:AssignCheck")

	beego.Router("/tx/exchange", &controllers.TxController{}, "post:Exchange")
	beego.Router("/currency/exchange/check/:uuid(.*)", &controllers.TxController{}, "get:ExchangeCheck")

	beego.Router("/tx/cancel", &controllers.TxController{}, "post:Cancel")
	beego.Router("/currency/cancel/check/:uuid(.*)", &controllers.TxController{}, "get:CancelCheck")
}
