package controllers

import (
	"github.com/wutongtree/exchange/client/models"
	"github.com/wutongtree/exchange/client/utils"
)

type CurrencyController struct {
	BaseController
}

func (c *CurrencyController) Index() {
	allCurrencys, _ := models.AllCurrencys()
	allCurrencyIds := []string{}
	for _, v := range allCurrencys {
		allCurrencyIds = append(allCurrencyIds, v.ID)
	}
	c.Data["allCurrencyIds"] = allCurrencyIds

	myCurrencys, _ := models.MyCurrencys(c.UserUserId)
	for k, v := range myCurrencys {
		myCurrencys[k].CreateDate = utils.GetDateMHS(v.CreateTime)
	}
	c.Data["myCurrencys"] = myCurrencys

	myCurrencyIds := []string{}
	myAssets, _ := models.MyAssets(c.UserUserId)
	for _, v := range myAssets {
		if v.Count > 0 {
			myCurrencyIds = append(myCurrencyIds, v.Currency)
		}
	}
	c.Data["myAssets"] = myAssets
	c.Data["myCurrencyIds"] = myCurrencyIds

	txs, _ := models.GetMyTxs(c.UserUserId)
	c.Data["txs"] = txs

	c.TplName = "currency/index.tpl"
}

func (c *CurrencyController) CreateGet() {
	c.TplName = "currency/create.tpl"
}

func (c *CurrencyController) CreatePost() {
	id := c.GetString("fundname")
	count, err := c.GetFloat("quotas")
	if err != nil {
		count = float64(0)
	}

	// 新建
	currency := &models.Currency{
		ID:    id,
		Count: count,
		User:  c.UserUserId,
	}
	_, err = models.CreateCurrency(currency)
	if err != nil {
		logger.Errorf("CreateCurrency error: %v", err)

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "新建币失败：" + err.Error()}
		c.ServeJSON()

		return
	}
	c.Data["json"] = map[string]interface{}{"code": 1, "message": "创建成功"} //txid}
	c.ServeJSON()
}

func (c *CurrencyController) CreateCheck() {
	txid := c.Ctx.Input.Param(":txid")
	code, err := models.CheckCreate(txid)
	if err != nil {
		logger.Errorf("CheckCreate error: %v", err)

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "检测失败：" + err.Error()}
		c.ServeJSON()

		return
	}
	c.Data["json"] = map[string]interface{}{"code": 1, "message": code}
	c.ServeJSON()
}

func (c *CurrencyController) ReleaseGet() {
	id := c.Ctx.Input.Param(":id")
	c.Data["id"] = id

	c.TplName = "currency/release.tpl"
}

func (c *CurrencyController) ReleasePost() {
	id := c.GetString("fundname")
	count, err := c.GetFloat("quotas")
	if err != nil {
		count = float64(0)
	}

	// 发布
	currency := &models.Currency{
		ID:    id,
		Count: count,
		User:  c.UserUserId,
	}
	_, err = models.ReleaseCurrency(currency)
	if err != nil {
		logger.Errorf("ReleaseCurrency error: %v", err)

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "发布基金失败：" + err.Error()}
		c.ServeJSON()

		return
	}
	c.Data["json"] = map[string]interface{}{"code": 1, "message": "发布成功"} // txid}
	c.ServeJSON()
}

func (c *CurrencyController) ReleaseCheck() {
	txid := c.Ctx.Input.Param(":txid")
	code, err := models.CheckRelease(txid)
	if err != nil {
		logger.Errorf("CheckRelease error: %v", err)

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "检测失败：" + err.Error()}
		c.ServeJSON()

		return
	}
	c.Data["json"] = map[string]interface{}{"code": 1, "message": code}
	c.ServeJSON()
}

func (c *CurrencyController) AssignGet() {
	id := c.Ctx.Input.Param(":id")
	c.Data["id"] = id

	users := []string{"lukas", "diego", "jim", "binhn", "alice", "bob", "assigner"}
	c.Data["users"] = users

	c.TplName = "currency/assign.tpl"
}

func (c *CurrencyController) AssignPost() {
	id := c.GetString("fundname")
	count, err := c.GetInt64("quotas")
	if err != nil {
		count = int64(0)
	}
	owner := c.GetString("user")

	// 分发
	assigns := &models.Assigns{
		User:     c.UserUserId,
		Currency: id,
		Assigns:  []models.Assign{models.Assign{Owner: owner, Count: count}},
	}

	_, err = models.AssignCurrency(assigns)
	if err != nil {
		logger.Errorf("AssignCurrency error: %v", err)

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "分发基金失败：" + err.Error()}
		c.ServeJSON()

		return
	}
	c.Data["json"] = map[string]interface{}{"code": 1, "message": "分发成功"} // txid}
	c.ServeJSON()
}

func (c *CurrencyController) AssignCheck() {
	txid := c.Ctx.Input.Param(":txid")
	code, err := models.CheckAssign(txid)
	if err != nil {
		logger.Errorf("CheckAssign error: %v", err)

		c.Data["json"] = map[string]interface{}{"code": 0, "message": "检测失败：" + err.Error()}
		c.ServeJSON()

		return
	}
	c.Data["json"] = map[string]interface{}{"code": 1, "message": code}
	c.ServeJSON()
}
