package main

import (
	"errors"
	"strconv"

	"github.com/hyperledger/fabric/core/util"
	pb "github.com/hyperledger/fabric/protos"
	"google.golang.org/grpc"
)

var (
	peerClientConn *grpc.ClientConn
	serverClient   pb.PeerClient
	chaincodePath  string
	chaincodeName  string
)

func deploy() (err error) {
	myLogger.Debug("------------- deploy")

	resp, err := deployInternal()
	if err != nil {
		myLogger.Errorf("Failed deploying [%s]", err)
		return
	}
	myLogger.Debugf("Resp [%s]", resp.String())
	myLogger.Debugf("Chaincode NAME: [%s]-[%s]", chaincodeName, string(resp.Msg))

	if resp.Status != pb.Response_SUCCESS {
		return errors.New(string(resp.Msg))
	}

	myLogger.Debug("------------- Done!")
	return
}

func createCurrency(currency string, count int64, user string) (txid string, err error) {
	invoker, err := setCryptoClient(user, "")
	if err != nil {
		myLogger.Errorf("Failed getting invoker [%s]", err)
		return
	}
	// invokerCert, err := invoker.GetTCertificateHandlerNext()
	// if err != nil {
	// 	myLogger.Errorf("Failed getting TCert [%s]", err)
	// 	return
	// }
	myLogger.Debugf("Chaincode [createCurrency] args:[%s]-[%s],[%s]-[%s]", "currency", currency, "count", count)

	// chaincodeInput := &pb.ChaincodeInput{Args: util.ToChaincodeArgs("createCurrency", currency, strconv.FormatInt(count, 10), base64.StdEncoding.EncodeToString(invokerCert.GetCertificate()))}
	chaincodeInput := &pb.ChaincodeInput{Args: util.ToChaincodeArgs("createCurrency", currency, strconv.FormatInt(count, 10), user)}

	return invokeChaincode(invoker, chaincodeInput)
}

func releaseCurrency(currency string, count int64, user string) (txid string, err error) {
	invoker, err := setCryptoClient(user, "")
	if err != nil {
		myLogger.Errorf("Failed getting invoker [%s]", err)
		return
	}
	// invokerCert, err := invoker.GetTCertificateHandlerNext()
	// if err != nil {
	// 	myLogger.Errorf("Failed getting TCert [%s]", err)
	// 	return
	// }
	myLogger.Debugf("Chaincode [releaseCurrency] args:[%s]-[%s],[%s]-[%s]", "currency", currency, "count", count)

	chaincodeInput := &pb.ChaincodeInput{Args: util.ToChaincodeArgs("releaseCurrency", currency, strconv.FormatInt(count, 10))}

	// return invokeChaincodeSigma(invoker, invokerCert, chaincodeInput)
	return invokeChaincode(invoker, chaincodeInput)
}

func assignCurrency(assigns string, user string) (txid string, err error) {
	invoker, err := setCryptoClient(user, "")
	if err != nil {
		myLogger.Errorf("Failed getting invoker [%s]", err)
		return
	}
	// invokerCert, err := invoker.GetTCertificateHandlerNext()
	// if err != nil {
	// 	myLogger.Errorf("Failed getting TCert [%s]", err)
	// 	return
	// }
	myLogger.Debugf("Chaincode [assignCurrency] args:[%s]-[%s]", "assigns", assigns)

	chaincodeInput := &pb.ChaincodeInput{Args: util.ToChaincodeArgs("assignCurrency", assigns)}

	// return invokeChaincodeSigma(invoker, invokerCert, chaincodeInput)
	return invokeChaincode(invoker, chaincodeInput)
}

func exchange(exchanges string) (err error) {
	myLogger.Debugf("Chaincode [exchange] args:[%s]-[%s]", "exchanges", exchanges)

	chaincodeInput := &pb.ChaincodeInput{Args: util.ToChaincodeArgs("exchange", exchanges)}

	_, err = invokeChaincode(adminInvoker, chaincodeInput)
	return
}

func lock(orders string, islock bool, srcMethod string) (txid string, err error) {
	myLogger.Debugf("Chaincode [lock] args:[%s]-[%s],[%s]-[%s],[%s]-[%s]", "orders", orders, "islock", islock, "srcMethod", srcMethod)

	chaincodeInput := &pb.ChaincodeInput{Args: util.ToChaincodeArgs("lock", orders, strconv.FormatBool(islock), srcMethod)}

	return invokeChaincode(adminInvoker, chaincodeInput)
}

func getCurrencys() (currencys string, err error) {
	chaincodeInput := &pb.ChaincodeInput{Args: util.ToChaincodeArgs("queryAllCurrency")}

	return queryChaincode(chaincodeInput)
}

func getCurrency(id string) (currency string, err error) {
	myLogger.Debugf("Chaincode [queryCurrencyByID] args:[%s]-[%s]", "id", id)

	chaincodeInput := &pb.ChaincodeInput{Args: util.ToChaincodeArgs("queryCurrencyByID", id)}

	return queryChaincode(chaincodeInput)
}

func getCurrencysByUser(user string) (currencys string, err error) {
	// invoker, err := setCryptoClient(user, "")
	// if err != nil {
	// 	myLogger.Errorf("Failed getting invoker [%s]", err)
	// 	return
	// }
	// invokerCert, err := invoker.GetTCertificateHandlerNext()
	// if err != nil {
	// 	myLogger.Errorf("Failed getting TCert [%s]", err)
	// 	return
	// }

	// cert := base64.StdEncoding.EncodeToString(invokerCert.GetCertificate())
	myLogger.Debugf("Chaincode [getCurrencysByUser] args:[%s]-[%s]", "user", user)

	chaincodeInput := &pb.ChaincodeInput{Args: util.ToChaincodeArgs("queryMyCurrency", user)}

	return queryChaincode(chaincodeInput)
}

func getAsset(owner string) (asset string, err error) {
	myLogger.Debugf("Chaincode [queryAssetByOwner] args:[%s]-[%s]", "owner", owner)
	chaincodeInput := &pb.ChaincodeInput{Args: util.ToChaincodeArgs("queryAssetByOwner", owner)}

	return queryChaincode(chaincodeInput)
}

func getTxLogs() (txLogs string, err error) {
	chaincodeInput := &pb.ChaincodeInput{Args: util.ToChaincodeArgs("queryTxLogs")}
	return queryChaincode(chaincodeInput)
}

func initAccount(user string) (result string, err error) {
	myLogger.Debugf("Chaincode [initAccount] args:[%s]-[%s]", "initAccount", user)

	chaincodeInput := &pb.ChaincodeInput{Args: util.ToChaincodeArgs("initAccount", user)}

	_, err = invokeChaincode(adminInvoker, chaincodeInput)
	return
}

func getMyReleaseLog(user string) (log string, err error) {
	myLogger.Debugf("Chaincode [getMyReleaseLog] args:[%s]-[%s]", "user", user)

	chaincodeInput := &pb.ChaincodeInput{Args: util.ToChaincodeArgs("queryMyReleaseLog", user)}

	return queryChaincode(chaincodeInput)
}

func getMyAssignLog(user string) (log string, err error) {
	myLogger.Debugf("Chaincode [getMyAssignLog] args:[%s]-[%s]", "user", user)

	chaincodeInput := &pb.ChaincodeInput{Args: util.ToChaincodeArgs("queryMyAssignLog", user)}

	return queryChaincode(chaincodeInput)
}
