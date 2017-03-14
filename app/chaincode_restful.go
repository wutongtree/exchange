package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric/core/util"
	pb "github.com/hyperledger/fabric/protos"
	"github.com/spf13/viper"
)

type loginResponse struct {
	OK    string `json:"OK,omitempty"`
	Error string `json:"Error,omitempty"`
}

// rpcRequest defines the JSON RPC 2.0 request payload for the /chaincode endpoint.
type rpcRequest struct {
	Jsonrpc string            `json:"jsonrpc,omitempty"`
	Method  string            `json:"method,omitempty"`
	Params  *pb.ChaincodeSpec `json:"params,omitempty"`
	ID      int64             `json:"id,omitempty"`
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

func deployChaincodeRest() (err error) {
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
	err = loginRest(loginReqBody)
	if err != nil {
		myLogger.Errorf("Failed login [%s]", err)
		return
	}

	request := &rpcRequest{
		Jsonrpc: "2.0",
		Method:  "deploy",
		Params: &pb.ChaincodeSpec{
			Type: pb.ChaincodeSpec_GOLANG,
			ChaincodeID: &pb.ChaincodeID{
				Path: chaincodePath,
			},
			CtorMsg:              &pb.ChaincodeInput{Args: util.ToChaincodeArgs("init")},
			SecureContext:        name,
			ConfidentialityLevel: confidentialityLevel,
		},
		ID: time.Now().Unix(),
	}

	reqBody, err := json.Marshal(request)
	if err != nil {
		myLogger.Errorf("Failed deploying [%s]", err)
		return
	}

	respBody, err := doHTTPPost(restURL+"/chaincode", reqBody)
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
	myLogger.Debugf("ChaincodeName [%s]", chaincodeName)

	myLogger.Debug("------------- deploy Done! -------------")

	return
}

func invokeChaincodeRest(secureContext string, chaincodeInput *pb.ChaincodeInput) (ret string, err error) {
	myLogger.Debug("------------- invoke chainde -------------")

	request := &rpcRequest{
		Jsonrpc: "2.0",
		Method:  "invoke",
		Params: &pb.ChaincodeSpec{
			Type: pb.ChaincodeSpec_GOLANG,
			ChaincodeID: &pb.ChaincodeID{
				Name: chaincodeName,
			},
			CtorMsg:              chaincodeInput,
			SecureContext:        secureContext,
			ConfidentialityLevel: confidentialityLevel,
		},
		ID: time.Now().Unix(),
	}

	reqBody, err := json.Marshal(request)
	if err != nil {
		myLogger.Errorf("Failed invoke [%s]", err)
		return
	}

	respBody, err := doHTTPPost(restURL+"/chaincode", reqBody)
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

func queryChaincodeRest(secureContext string, chaincodeInput *pb.ChaincodeInput) (ret string, err error) {
	myLogger.Debug("------------- invoke chainde -------------")

	request := &rpcRequest{
		Jsonrpc: "2.0",
		Method:  "query",
		Params: &pb.ChaincodeSpec{
			Type: pb.ChaincodeSpec_GOLANG,
			ChaincodeID: &pb.ChaincodeID{
				Name: chaincodeName,
			},
			CtorMsg: chaincodeInput,
			// SecureContext:        secureContext,
			ConfidentialityLevel: confidentialityLevel,
		},
		ID: time.Now().Unix(),
	}

	reqBody, err := json.Marshal(request)
	if err != nil {
		myLogger.Errorf("Failed invoke [%s]", err)
		return
	}

	respBody, err := doHTTPPost(restURL+"/chaincode", reqBody)
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

func loginRest(reqBody []byte) (err error) {
	myLogger.Debug("------------- login -------------")

	respBody, err := doHTTPPost(restURL+"/registrar", reqBody)
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
