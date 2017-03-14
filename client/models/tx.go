package models

import (
	"encoding/json"
	"fmt"
)

// Order Order
type Order struct {
	UUID         string  `json:"uuid"`        //UUID
	Account      string  `json:"account"`     //账户
	SrcCurrency  string  `json:"srcCurrency"` //源币种代码
	SrcCount     float64 `json:"srcCount"`    //源币种交易数量
	DesCurrency  string  `json:"desCurrency"` //目标币种代码
	DesCount     float64 `json:"desCount"`    //目标币种交易数量
	IsBuyAll     bool    `json:"isBuyAll"`    //是否买入所有，即为true是以目标币全部兑完为主,否则算部分成交,买完为止；为false则是以源币全部兑完为主,否则算部分成交，卖完为止
	ExpiredTime  int64   `json:"expiredTime"` //超时时间
	ExpiredDate  string  `json:"expiredDate"`
	PendingTime  int64   `json:"PendingTime"` //挂单时间
	PendingDate  string  `json:"pendingDate"`
	PendedTime   int64   `json:"PendedTime"` //挂单完成时间
	PendedDate   string  `json:"pendedDate"`
	MatchedTime  int64   `json:"matchedTime"` //撮合完成时间
	MatchedDate  string  `json:"matchedDate"`
	FinishedTime int64   `json:"finishedTime"` //交易完成时间
	FinishedDate string  `json:"finishedDate"`
	RawUUID      string  `json:"rawUUID"`   //母单UUID
	Metadata     string  `json:"metadata"`  //存放其他数据，如挂单锁定失败信息
	FinalCost    float64 `json:"finalCost"` //源币的最终消耗数量，主要用于买完（IsBuyAll=true）的最后一笔交易计算结余，此时SrcCount有可能大于FinalCost
	Status       int     `json:"status"`    //状态 0：待交易，1：完成，2：过期，3：撤单
}

func GetMyTxs(user string) ([]Order, error) {
	urlstr := getHTTPURL("user/tx/" + user)

	response, err := performHTTPGet(urlstr)
	if err != nil {
		logger.Errorf("GetMyTxs failed: %v", err)
		return nil, err
	}

	logger.Debugf("GetMyTxs: url=%v request=%v response=%v", urlstr, "nil", string(response))

	var result struct {
		OK  []Order
		Err string
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		logger.Errorf("GetMyTxs failed: %v", err)
		return nil, err
	}

	if len(result.Err) != 0 {
		logger.Errorf("GetMyTxs failed: %v", result.Err)
		return nil, fmt.Errorf(result.Err)
	}

	return result.OK, nil
}

// TxExchange TxExchange
func TxExchange(order *Order) (string, error) {
	// Create New Fund
	urlstr := getHTTPURL("tx/exchange")

	reqBody, err := json.Marshal(order)
	if err != nil {
		return "", err
	}
	response, err := performHTTPPost(urlstr, reqBody)
	if err != nil {
		logger.Errorf("CreateCurrency failed: %v", err)
		return "", err
	}

	logger.Debugf("CreateCurrency: url=%v request=%v response=%v", urlstr, order, string(response))

	var result AppResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		logger.Errorf("CreateCurrency failed: %v", err)
		return "", err
	}

	if len(result.OK) == 0 {
		logger.Errorf("CreateCurrency failed: %v", result.Err)
		return "", fmt.Errorf(result.Err)
	}

	return result.OK, nil
}

// CheckExchange CheckExchange
func CheckExchange(uuid string) (string, error) {
	// Create New Fund
	urlstr := getHTTPURL("tx/exchange/check/" + uuid)

	response, err := performHTTPGet(urlstr)
	if err != nil {
		logger.Errorf("CheckAssign failed: %v", err)
		return "", err
	}

	logger.Debugf("CheckAssign: url=%v request=%v response=%v", urlstr, uuid, string(response))

	var result AppResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		logger.Errorf("CheckAssign failed: %v", err)
		return "", err
	}

	if len(result.Err) != 0 {
		logger.Errorf("CheckAssign failed: %v", result.Err)
		return "", fmt.Errorf(result.Err)
	}

	return result.OK, nil
}

// TxCancel TxCancel
func TxCancel(id string) (string, error) {
	// Create New Fund
	urlstr := getHTTPURL("tx/cancel")

	response, err := performHTTPPost(urlstr, []byte(id))
	if err != nil {
		logger.Errorf("CreateCurrency failed: %v", err)
		return "", err
	}

	logger.Debugf("CreateCurrency: url=%v request=%v response=%v", urlstr, id, string(response))

	var result AppResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		logger.Errorf("CreateCurrency failed: %v", err)
		return "", err
	}

	if len(result.OK) == 0 {
		logger.Errorf("CreateCurrency failed: %v", result.Err)
		return "", fmt.Errorf(result.Err)
	}

	return result.OK, nil
}

// CheckCancel CheckCancel
func CheckCancel(uuid string) (string, error) {
	// Create New Fund
	urlstr := getHTTPURL("tx/cancel/check/" + uuid)

	response, err := performHTTPGet(urlstr)
	if err != nil {
		logger.Errorf("CheckAssign failed: %v", err)
		return "", err
	}

	logger.Debugf("CheckAssign: url=%v request=%v response=%v", urlstr, uuid, string(response))

	var result AppResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		logger.Errorf("CheckAssign failed: %v", err)
		return "", err
	}

	if len(result.Err) != 0 {
		logger.Errorf("CheckAssign failed: %v", result.Err)
		return "", fmt.Errorf(result.Err)
	}

	return result.OK, nil
}
