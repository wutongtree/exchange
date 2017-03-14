package models

import (
	"encoding/json"
	"fmt"
)

// AppResponse AppResponse
type AppResponse struct {
	OK  string
	Err string
}

type Currency struct {
	ID         string  `json:"id"`
	Count      float64 `json:"count"`
	LeftCount  float64 `json:"leftCount"`
	User       string  `json:"user"`
	CreateTime int64   `json:"createTime"`
	CreateDate string  `json:"createDate"`
}

type Assign struct {
	Owner string `json:"owner"`
	Count int64  `json:"count"`
}

type Assigns struct {
	User     string   `json:"user"`
	Currency string   `json:"currency"`
	Assigns  []Assign `json:"assigns"`
}

type Asset struct {
	Owner     string  `json:"owner"`
	Currency  string  `json:"currency"`
	Count     float64 `json:"count"`
	LockCount float64 `json:"lockCount"`
}

// MyCurrencys MyCurrencys
func MyCurrencys(user string) ([]Currency, error) {
	urlstr := getHTTPURL("user/currency/" + user)

	response, err := performHTTPGet(urlstr)
	if err != nil {
		logger.Errorf("MyCurrencys failed: %v", err)
		return nil, err
	}

	logger.Debugf("MyCurrencys: url=%v request=%v response=%v", urlstr, user, string(response))

	var result struct {
		OK  []Currency
		Err string
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		logger.Errorf("MyCurrencys failed: %v", err)
		return nil, err
	}

	if len(result.Err) != 0 {
		logger.Errorf("MyCurrencys failed: %v", result.Err)
		return nil, fmt.Errorf(result.Err)
	}

	return result.OK, nil
}

// AllCurrencys AllCurrencys
func AllCurrencys() ([]Currency, error) {
	urlstr := getHTTPURL("currency")

	response, err := performHTTPGet(urlstr)
	if err != nil {
		logger.Errorf("AllCurrencys failed: %v", err)
		return nil, err
	}

	logger.Debugf("AllCurrencys: url=%v request=%v response=%v", urlstr, "nil", string(response))

	var result struct {
		OK  []Currency
		Err string
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		logger.Errorf("AllCurrencys failed: %v", err)
		return nil, err
	}

	if len(result.Err) != 0 {
		logger.Errorf("AllCurrencys failed: %v", result.Err)
		return nil, fmt.Errorf(result.Err)
	}

	return result.OK, nil
}

// MyAssets MyAssets
func MyAssets(user string) ([]Asset, error) {
	urlstr := getHTTPURL("user/asset/" + user)

	response, err := performHTTPGet(urlstr)
	if err != nil {
		logger.Errorf("MyAssets failed: %v", err)
		return nil, err
	}

	logger.Debugf("MyAssets: url=%v request=%v response=%v", urlstr, user, string(response))

	var result struct {
		OK  []Asset
		Err string
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		logger.Errorf("MyAssets failed: %v", err)
		return nil, err
	}

	if len(result.Err) != 0 {
		logger.Errorf("MyAssets failed: %v", result.Err)
		return nil, fmt.Errorf(result.Err)
	}

	return result.OK, nil
}

// MyCurrencys MyCurrencys
func CurrencyId(id string) (*Currency, error) {
	urlstr := getHTTPURL("currency/" + id)

	response, err := performHTTPGet(urlstr)
	if err != nil {
		logger.Errorf("CurrencyId failed: %v", err)
		return nil, err
	}

	logger.Debugf("CurrencyId: url=%v request=%v response=%v", urlstr, id, string(response))

	var result struct {
		OK  Currency
		Err string
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		logger.Errorf("CurrencyId failed: %v", err)
		return nil, err
	}

	if len(result.Err) != 0 {
		logger.Errorf("CurrencyId failed: %v", result.Err)
		return nil, fmt.Errorf(result.Err)
	}

	return &result.OK, nil
}

// CreateCurrency CreateCurrency
func CreateCurrency(currency *Currency) (string, error) {
	// Create New Fund
	urlstr := getHTTPURL("currency/create")

	reqBody, err := json.Marshal(currency)
	if err != nil {
		return "", err
	}
	response, err := performHTTPPost(urlstr, reqBody)
	if err != nil {
		logger.Errorf("CreateCurrency failed: %v", err)
		return "", err
	}

	logger.Debugf("CreateCurrency: url=%v request=%v response=%v", urlstr, currency, string(response))

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

// CheckCreate CheckCreate
func CheckCreate(txid string) (string, error) {
	// Create New Fund
	urlstr := getHTTPURL("currency/create/check/" + txid)

	response, err := performHTTPGet(urlstr)
	if err != nil {
		logger.Errorf("CheckCreate failed: %v", err)
		return "", err
	}

	logger.Debugf("CheckCreate: url=%v request=%v response=%v", urlstr, txid, string(response))

	var result AppResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		logger.Errorf("CheckCreate failed: %v", err)
		return "", err
	}

	if len(result.Err) != 0 {
		logger.Errorf("CheckCreate failed: %v", result.Err)
		return "", fmt.Errorf(result.Err)
	}

	return result.OK, nil
}

// ReleaseCurrency ReleaseCurrency
func ReleaseCurrency(currency *Currency) (string, error) {
	// Create New Fund
	urlstr := getHTTPURL("currency/release")

	reqBody, err := json.Marshal(currency)
	if err != nil {
		return "", err
	}
	response, err := performHTTPPost(urlstr, reqBody)
	if err != nil {
		logger.Errorf("ReleaseCurrency failed: %v", err)
		return "", err
	}

	logger.Debugf("ReleaseCurrency: url=%v request=%v response=%v", urlstr, currency, string(response))

	var result AppResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		logger.Errorf("ReleaseCurrency failed: %v", err)
		return "", err
	}

	if len(result.OK) == 0 {
		logger.Errorf("ReleaseCurrency failed: %v", result.Err)
		return "", fmt.Errorf(result.Err)
	}

	return result.OK, nil
}

// CheckRelease CheckRelease
func CheckRelease(txid string) (string, error) {
	// Create New Fund
	urlstr := getHTTPURL("currency/release/check/" + txid)

	response, err := performHTTPGet(urlstr)
	if err != nil {
		logger.Errorf("CheckRelease failed: %v", err)
		return "", err
	}

	logger.Debugf("CheckRelease: url=%v request=%v response=%v", urlstr, txid, string(response))

	var result AppResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		logger.Errorf("CheckRelease failed: %v", err)
		return "", err
	}

	if len(result.Err) != 0 {
		logger.Errorf("CheckRelease failed: %v", result.Err)
		return "", fmt.Errorf(result.Err)
	}

	return result.OK, nil
}

// AssignCurrency AssignCurrency
func AssignCurrency(assigns *Assigns) (string, error) {
	// Create New Fund
	urlstr := getHTTPURL("currency/assign")

	reqBody, err := json.Marshal(assigns)
	if err != nil {
		return "", err
	}
	response, err := performHTTPPost(urlstr, reqBody)
	if err != nil {
		logger.Errorf("ReleaseCurrency failed: %v", err)
		return "", err
	}

	logger.Debugf("AssignCurrency: url=%v request=%v response=%v", urlstr, assigns, string(response))

	var result AppResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		logger.Errorf("AssignCurrency failed: %v", err)
		return "", err
	}

	if len(result.OK) == 0 {
		logger.Errorf("AssignCurrency failed: %v", result.Err)
		return "", fmt.Errorf(result.Err)
	}

	return result.OK, nil
}

// CheckAssign CheckAssign
func CheckAssign(txid string) (string, error) {
	// Create New Fund
	urlstr := getHTTPURL("currency/assign/check/" + txid)

	response, err := performHTTPGet(urlstr)
	if err != nil {
		logger.Errorf("CheckAssign failed: %v", err)
		return "", err
	}

	logger.Debugf("CheckAssign: url=%v request=%v response=%v", urlstr, txid, string(response))

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
