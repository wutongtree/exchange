package controllers

import (
	"fmt"

	"github.com/wutongtree/exchange/client/models"
)

// LoginController main controller
type LoginController struct {
	BaseController
}

// Get default url
func (c *LoginController) Get() {
	if c.IsLogin {
		c.Abort("401")
	} else {
		c.TplName = "login.tpl"
	}
}

// Login to system
func (c *LoginController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")

	if "" == username {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写用户名"}
		c.ServeJSON()
	}

	if "" == password {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写密码"}
		c.ServeJSON()
	}

	logined := models.Login(username, password)
	if logined {
		c.SetSession("userLogin_exchange", username+"||"+username+"||"+username)

		fmt.Printf("Login successful: %s\n", username)

		c.Data["json"] = map[string]interface{}{"code": 1, "message": "贺喜你，登录成功！"}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "登录失败，请输入正确的用户名和密码！"}
	}
	c.ServeJSON()
}

// Logout to system
func (c *LoginController) Logout() {
	c.DelSession("userLogin_exchange")

	c.Redirect("/", 302)
	return
}
