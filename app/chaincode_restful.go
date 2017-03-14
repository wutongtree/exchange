package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	pb "github.com/hyperledger/fabric/protos"
	"github.com/spf13/viper"
)

type loginResponse struct {
	OK    string `json:"OK,omitempty"`
	Error string `json:"Error,omitempty"`
}

// Carries the chaincode function and its arguments.
// UnmarshalJSON in transaction.go converts the string-based REST/JSON input to
// the []byte-based current ChaincodeInput structure.
type ChaincodeInput struct {
	Function string   `protobuf:"bytes,1,rep,name=function,proto3" json:"function,omitempty"`
	Args     []string `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
}

// Carries the chaincode specification. This is the actual metadata required for
// defining a chaincode.
type ChaincodeSpec struct {
	Type                 pb.ChaincodeSpec_Type   `protobuf:"varint,1,opt,name=type,enum=protos.ChaincodeSpec_Type" json:"type,omitempty"`
	ChaincodeID          *pb.ChaincodeID         `protobuf:"bytes,2,opt,name=chaincodeID" json:"chaincodeID,omitempty"`
	CtorMsg              *ChaincodeInput         `protobuf:"bytes,3,opt,name=ctorMsg" json:"ctorMsg,omitempty"`
	Timeout              int32                   `protobuf:"varint,4,opt,name=timeout" json:"timeout,omitempty"`
	SecureContext        string                  `protobuf:"bytes,5,opt,name=secureContext" json:"secureContext,omitempty"`
	ConfidentialityLevel pb.ConfidentialityLevel `protobuf:"varint,6,opt,name=confidentialityLevel,enum=protos.ConfidentialityLevel" json:"confidentialityLevel,omitempty"`
	Metadata             []byte                  `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Attributes           []string                `protobuf:"bytes,8,rep,name=attributes" json:"attributes,omitempty"`
}

// rpcRequest defines the JSON RPC 2.0 request payload for the /chaincode endpoint.
type rpcRequest struct {
	Jsonrpc string         `json:"jsonrpc,omitempty"`
	Method  string         `json:"method,omitempty"`
	Params  *ChaincodeSpec `json:"params,omitempty"`
	ID      int64          `json:"id,omitempty"`
}

type rpcID struct {
	StringValue string
	IntValue    int64
}

// rpcResponse defines the JSON RPC 2.0 response payload for the /chaincode endpoint.
type rpcResponse struct {
	Jsonrpc string     `json:"jsonrpc,omitempty"`
	Result  *rpcResult `json:"result,omitempty"`
	Error   *rpcError  `json:"error,omitempty"`
	ID      int64      `json:"id"`
}

// rpcResult defines the structure for an rpc sucess/error result message.
type rpcResult struct {
	Status  string    `json:"status,omitempty"`
	Message string    `json:"message,omitempty"`
	Error   *rpcError `json:"error,omitempty"`
}

// rpcError defines the structure for an rpc error.
type rpcError struct {
	// A Number that indicates the error type that occurred. This MUST be an integer.
	Code int64 `json:"code,omitempty"`
	// A String providing a short description of the error. The message SHOULD be
	// limited to a concise single sentence.
	Message string `json:"message,omitempty"`
	// A Primitive or Structured value that contains additional information about
	// the error. This may be omitted. The value of this member is defined by the
	// Server (e.g. detailed error information, nested errors etc.).
	Data string `json:"data,omitempty"`
}

func doHTTPPost(url string, reqBody []byte) ([]byte, error) {
	resp, err := http.Post(url, "application/json;charset=utf-8", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func loginRestful(reqBody []byte) (err error) {
	myLogger.Debug("------------- login -------------")

	respBody, err := doHTTPPost(restURL+"registrar", reqBody)
	if err != nil {
		myLogger.Errorf("Failed login [%s]", err)
		return
	}

	result := new(loginResponse)
	err = json.Unmarshal(respBody, result)
	if err != nil {
		myLogger.Errorf("Failed login [%s]", err)
		return
	}

	myLogger.Debugf("Resp [%s]", string(respBody))

	if result.Error != "" {
		myLogger.Errorf("Failed login [%s]", result.Error)
		return
	}

	myLogger.Infof("Successful login [%s]", result.OK)
	myLogger.Debug("------------- login! -------------")

	return
}

func deployChaincodeRestful() (err error) {
	myLogger.Debug("------------- deploy chaincode -------------")

	chaincodeName = viper.GetString("chaincode.id.name")
	if chaincodeName != "" {
		myLogger.Infof("Using existing chaincode [%s]", chaincodeName)
		return
	}

	chaincodePath = viper.GetString("chaincode.id.path")
	name := viper.GetString("app.admin.name")
	pwd := viper.GetString("app.admin.pwd")

	if chaincodePath == "" || name == "" || pwd == "" {
		err = fmt.Errorf("config error: check your config.yaml")
		return
	}

	loginRequest := &User{
		EnrollID:     name,
		EnrollSecret: pwd,
	}
	loginReqBody, err := json.Marshal(loginRequest)
	err = loginRestful(loginReqBody)
	if err != nil {
		myLogger.Errorf("Failed login [%s]", err)
		return
	}

	request := &rpcRequest{
		Jsonrpc: "2.0",
		Method:  "deploy",
		Params: &ChaincodeSpec{
			Type: pb.ChaincodeSpec_GOLANG,
			ChaincodeID: &pb.ChaincodeID{
				Path: chaincodePath,
			},
			CtorMsg: &ChaincodeInput{
				Function: "init",
				Args:     []string{},
			},
			//Timeout:1,
			SecureContext:        name,
			ConfidentialityLevel: confidentialityLevel,
			// Metadata:             adminCert.GetCertificate(),
			//Attributes:[]string{},
		},
		ID: time.Now().Unix(),
	}

	reqBody, err := json.Marshal(request)
	if err != nil {
		myLogger.Errorf("Failed deploying [%s]", err)
		return
	}

	respBody, err := doHTTPPost(restURL+"chaincode", reqBody)
	if err != nil {
		myLogger.Errorf("Failed deploying [%s]", err)
		return
	}

	result := new(rpcResponse)
	err = json.Unmarshal(respBody, result)
	if err != nil {
		myLogger.Errorf("Failed deploying [%s]", err)
		return
	}

	myLogger.Debugf("Resp [%s]", string(respBody))

	if result.Error != nil {
		myLogger.Errorf("Failed deploying [%s]", result.Error.Message)
		return
	}
	if result.Result.Status != "OK" {
		myLogger.Errorf("Failed deploying [%s]", result.Result.Message)
		return
	}

	chaincodeName = result.Result.Message

	myLogger.Debug("------------- deploy Done! -------------")

	return
}

func invokeChaincodeRestful(secureContext string, chaincodeInput *ChaincodeInput) (ret string, err error) {
	myLogger.Debug("------------- invoke chainde -------------")

	request := &rpcRequest{
		Jsonrpc: "2.0",
		Method:  "invoke",
		Params: &ChaincodeSpec{
			Type: pb.ChaincodeSpec_GOLANG,
			ChaincodeID: &pb.ChaincodeID{
				Name: chaincodeName,
			},
			CtorMsg: chaincodeInput,
			//Timeout:1,
			SecureContext:        secureContext,
			ConfidentialityLevel: confidentialityLevel,
			// Metadata:             adminCert.GetCertificate(),
			//Attributes:[]string{},
		},
		ID: time.Now().Unix(),
	}

	reqBody, err := json.Marshal(request)
	if err != nil {
		myLogger.Errorf("Failed invoke [%s]", err)
		return
	}

	respBody, err := doHTTPPost(restURL+"chaincode", reqBody)
	if err != nil {
		myLogger.Errorf("Failed invoke [%s]", err)
		return
	}

	result := new(rpcResponse)
	err = json.Unmarshal(respBody, result)
	if err != nil {
		myLogger.Errorf("Failed invoke [%s]", err)
		return
	}

	myLogger.Debugf("Resp [%s]", string(respBody))

	if result.Error != nil {
		myLogger.Errorf("Failed invoke [%s]", result.Error.Message)
		err = fmt.Errorf("result.Error.Message")
		return
	}
	if result.Result.Status != "OK" {
		myLogger.Errorf("Failed invoke [%s]", result.Result.Message)
		err = fmt.Errorf("result.Result.Message")
		return
	}

	myLogger.Debug("------------- invoke chainde Done! -------------")

	ret = result.Result.Message
	return
}

func queryChaincodeRestful(secureContext string, chaincodeInput *ChaincodeInput) (ret string, err error) {
	myLogger.Debug("------------- invoke chainde -------------")

	request := &rpcRequest{
		Jsonrpc: "2.0",
		Method:  "query",
		Params: &ChaincodeSpec{
			Type: pb.ChaincodeSpec_GOLANG,
			ChaincodeID: &pb.ChaincodeID{
				Name: chaincodeName,
			},
			CtorMsg: chaincodeInput,
			//Timeout:1,
			SecureContext:        secureContext,
			ConfidentialityLevel: confidentialityLevel,
			// Metadata:             adminCert.GetCertificate(),
			//Attributes:[]string{},
		},
		ID: time.Now().Unix(),
	}

	reqBody, err := json.Marshal(request)
	if err != nil {
		myLogger.Errorf("Failed invoke [%s]", err)
		return
	}

	respBody, err := doHTTPPost(restURL+"chaincode", reqBody)
	if err != nil {
		myLogger.Errorf("Failed invoke [%s]", err)
		return
	}

	result := new(rpcResponse)
	err = json.Unmarshal(respBody, result)
	if err != nil {
		myLogger.Errorf("Failed invoke [%s]", err)
		return
	}

	myLogger.Debugf("Resp [%s]", string(respBody))

	if result.Error != nil {
		myLogger.Errorf("Failed invoke [%s]", result.Error.Message)
		err = fmt.Errorf("result.Error.Message")
		return
	}
	if result.Result.Status != "OK" {
		myLogger.Errorf("Failed invoke [%s]", result.Result.Message)
		err = fmt.Errorf("result.Result.Message")
		err = fmt.Errorf("")
		return
	}

	myLogger.Debug("------------- invoke chainde Done! -------------")

	if result.Result.Message == "null" {
		return
	}
	ret = result.Result.Message
	return
}

// bz
func createCurrencyRestful(secureContext, currency string, count int64, user string) (txid string, err error) {
	myLogger.Debugf("Chaincode [createCurrency] args:[%s]-[%s],[%s]-[%s]", "currency", currency, "count", count)

	chaincodeInput := &ChaincodeInput{
		Function: "createCurrency",
		Args:     []string{currency, strconv.FormatInt(count, 10), user}}

	return invokeChaincodeRestful(secureContext, chaincodeInput)
}

func releaseCurrencyRestful(secureContext, currency string, count int64, user string) (txid string, err error) {
	myLogger.Debugf("Chaincode [releaseCurrency] args:[%s]-[%s],[%s]-[%s]", "currency", currency, "count", count)

	chaincodeInput := &ChaincodeInput{
		Function: "releaseCurrency",
		Args:     []string{currency, strconv.FormatInt(count, 10)}}

	return invokeChaincodeRestful(secureContext, chaincodeInput)
}

func assignCurrencyRestful(secureContext, assigns string, user string) (txid string, err error) {
	myLogger.Debugf("Chaincode [assignCurrency] args:[%s]-[%s]", "assigns", assigns)

	chaincodeInput := &ChaincodeInput{
		Function: "assignCurrency",
		Args:     []string{assigns}}

	return invokeChaincodeRestful(secureContext, chaincodeInput)
}

func exchangeRestful(secureContext, exchanges string) (err error) {
	myLogger.Debugf("Chaincode [exchange] args:[%s]-[%s]", "exchanges", exchanges)

	chaincodeInput := &ChaincodeInput{
		Function: "exchange",
		Args:     []string{exchanges}}

	_, err = invokeChaincodeRestful(secureContext, chaincodeInput)
	return
}

func lockRestful(secureContext, orders string, islock bool, srcMethod string) (txid string, err error) {
	myLogger.Debugf("Chaincode [lock] args:[%s]-[%s],[%s]-[%s],[%s]-[%s]", "orders", orders, "islock", islock, "srcMethod", srcMethod)

	chaincodeInput := &ChaincodeInput{
		Function: "lock",
		Args:     []string{orders, strconv.FormatBool(islock), srcMethod}}

	return invokeChaincodeRestful(secureContext, chaincodeInput)
}

func getCurrencysRestful(secureContext string) (currencys string, err error) {
	chaincodeInput := &ChaincodeInput{
		Function: "queryAllCurrency",
		Args:     []string{}}

	return queryChaincodeRestful(secureContext, chaincodeInput)
}

func getCurrencyRestful(secureContext, id string) (currency string, err error) {
	myLogger.Debugf("Chaincode [queryCurrencyByID] args:[%s]-[%s]", "id", id)

	chaincodeInput := &ChaincodeInput{
		Function: "queryCurrencyByID",
		Args:     []string{id}}

	return queryChaincodeRestful(secureContext, chaincodeInput)
}

func getCurrencysByUserRestful(secureContext, user string) (currencys string, err error) {
	myLogger.Debugf("Chaincode [getCurrencysByUser] args:[%s]-[%s]", "user", user)

	chaincodeInput := &ChaincodeInput{
		Function: "queryMyCurrency",
		Args:     []string{user}}

	return queryChaincodeRestful(secureContext, chaincodeInput)
}

func getAssetRestful(secureContext, owner string) (asset string, err error) {
	myLogger.Debugf("Chaincode [queryAssetByOwner] args:[%s]-[%s]", "owner", owner)
	chaincodeInput := &ChaincodeInput{
		Function: "queryAssetByOwner",
		Args:     []string{owner}}

	return queryChaincodeRestful(secureContext, chaincodeInput)
}

func getTxLogsRestful(secureContext string) (txLogs string, err error) {
	chaincodeInput := &ChaincodeInput{
		Function: "queryTxLogs",
		Args:     []string{}}

	return queryChaincodeRestful(secureContext, chaincodeInput)
}

func initAccountRestful(secureContext string, user string) (result string, err error) {
	myLogger.Debugf("Chaincode [initAccount] args:[%s]-[%s]", "initAccount", user)

	chaincodeInput := &ChaincodeInput{
		Function: "initAccount",
		Args:     []string{user}}

	_, err = invokeChaincodeRestful(secureContext, chaincodeInput)
	return
}

func getMyReleaseLogRestful(secureContext, user string) (log string, err error) {
	myLogger.Debugf("Chaincode [getMyReleaseLog] args:[%s]-[%s]", "user", user)

	chaincodeInput := &ChaincodeInput{
		Function: "queryMyReleaseLog",
		Args:     []string{user}}

	return queryChaincodeRestful(secureContext, chaincodeInput)
}

func getMyAssignLogRestful(secureContext, user string) (log string, err error) {
	myLogger.Debugf("Chaincode [getMyAssignLog] args:[%s]-[%s]", "user", user)

	chaincodeInput := &ChaincodeInput{
		Function: "queryMyAssignLog",
		Args:     []string{user}}

	return queryChaincodeRestful(secureContext, chaincodeInput)
}
