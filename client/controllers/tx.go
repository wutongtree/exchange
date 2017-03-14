package controllers

import "github.com/wutongtree/exchange/client/models"

type TxController struct {
	BaseController
}

func (c *TxController) Exchange() {
	srcCurrency := c.GetString("srcCurrency")
	desCurrency := c.GetString("desCurrency")
	srcCount, err := c.GetFloat("srcCount")
	if err != nil {
		logger.Errorf("ParseFloat error: %v", err)
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "挂单失败：" + err.Error()}
		c.ServeJSON()
		return
	}
	desCount, err := c.GetFloat("desCount")
	if err != nil {
		logger.Errorf("ParseFloat error: %v", err)
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "挂单失败：" + err.Error()}
		c.ServeJSON()
		return
	}
	isBuyAll, err := c.GetBool("isBuyAll")
	if err != nil {
		logger.Errorf("ParseFloat error: %v", err)
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "挂单失败：" + err.Error()}
		c.ServeJSON()
		return
	}
	// expiredTime, err := c.GetInt64("expiredTime")
	// if err != nil {
	// 	logger.Errorf("ParseFloat error: %v", err)
	// 	c.Data["json"] = map[string]interface{}{"code": 0, "message": "挂单失败：" + err.Error()}
	// 	c.ServeJSON()
	// 	return
	// }

	// 挂单
	order := &models.Order{
		Account:     c.UserUserId,
		SrcCurrency: srcCurrency,
		SrcCount:    srcCount,
		DesCurrency: desCurrency,
		DesCount:    desCount,
		IsBuyAll:    isBuyAll,
		// ExpiredTime: expiredTime,
	}
	_, err = models.TxExchange(order)
	if err != nil {
		logger.Errorf("ParseFloat error: %v", err)

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "新建币失败：" + err.Error()}
		c.ServeJSON()

		return
	}
	c.Data["json"] = map[string]interface{}{"code": 1, "message": "挂单成功"} //uuid}
	c.ServeJSON()
}

func (c *TxController) ExchangeCheck() {
	uuid := c.Ctx.Input.Param(":uuid")
	code, err := models.CheckExchange(uuid)
	if err != nil {
		logger.Errorf("ExchangeCheck error: %v", err)

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "检测失败：" + err.Error()}
		c.ServeJSON()

		return
	}
	c.Data["json"] = map[string]interface{}{"code": 1, "message": code}
	c.ServeJSON()
}

func (c *TxController) Cancel() {
	id := c.GetString("id")

	_, err := models.TxCancel(id)
	if err != nil {
		logger.Errorf("ParseFloat error: %v", err)

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "撤单失败：" + err.Error()}
		c.ServeJSON()

		return
	}
	c.Data["json"] = map[string]interface{}{"code": 1, "message": "撤单成功"} // uuid}
	c.ServeJSON()
}

func (c *TxController) CancelCheck() {
	uuid := c.Ctx.Input.Param(":uuid")
	code, err := models.CheckCancel(uuid)
	if err != nil {
		logger.Errorf("CheckCancel error: %v", err)

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "检测失败：" + err.Error()}
		c.ServeJSON()

		return
	}
	c.Data["json"] = map[string]interface{}{"code": 1, "message": code}
	c.ServeJSON()
}
