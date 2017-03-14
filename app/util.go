package main

import (
	"encoding/json"
	"math"
	"strconv"
	"strings"
)

// isJSON is a helper function to determine if a given string is proper JSON.
func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

// // formatRPCError formats the ERROR response to aid in JSON RPC 2.0 implementation
// func formatError(code int64, msg string, data string) rpcResult {
// 	err := &rpcError{Code: code, Message: msg, Data: data}
// 	error := rpcResult{Status: "Error", Error: err}

// 	return error
// }

// // formatRPCOK formats the OK response to aid in JSON RPC 2.0 implementation
// func formatOK(msg string) rpcResult {
// 	result := rpcResult{Status: "OK", Message: msg}

// 	return result
// }

// // formatRPCResponse consumes either an RPC ERROR or OK rpcResult and formats it
// // in accordance with the JSON RPC 2.0 specification.
// func formatResponse(res rpcResult, id *rpcID) rpcResponse {
// 	var response rpcResponse

// 	// Format a successful response
// 	if res.Status == "OK" {
// 		response = rpcResponse{Jsonrpc: "2.0", Result: &res, ID: id}
// 	} else {
// 		// Format an error response
// 		response = rpcResponse{Jsonrpc: "2.0", Error: res.Error, ID: id}
// 	}

// 	return response
// }

func convertInteger2Decimal(num int64) float64 {
	nStr := strconv.FormatInt(num, 10)

	return float64(num) / math.Pow10(len(nStr))
}

func getScore(srcCount, desCount float64, time int64) float64 {
	return round(desCount/srcCount, 6)*math.Pow10(6) + convertInteger2Decimal(time)
}

func convertPriceReciprocal(num int64) float64 {
	return math.Pow10(10) / float64(num)
}

func getBSKey(srcCurrency, desCurrency string) string {
	return ExchangeKey + "_" + srcCurrency + "_" + desCurrency
}

func getBSKeyByOne(key string) string {
	splits := strings.Split(key, "_")

	return getBSKey(splits[2], splits[1])
}

// round 四舍五入
func round(val float64, places int) float64 {
	var t float64
	f := math.Pow10(places)
	x := val * f
	if math.IsInf(x, 0) || math.IsNaN(x) {
		return val
	}
	if x >= 0.0 {
		t = math.Ceil(x)
		if (t - x) > 0.50000000001 {
			t -= 1.0
		}
	} else {
		t = math.Ceil(-x)
		if (t + x) > 0.50000000001 {
			t -= 1.0
		}
		t = -t
	}
	x = t / f

	if !math.IsInf(x, 0) {
		return x
	}

	return t
}

// 对挂单按时间排序
type Orders []*Order

func (x Orders) Len() int           { return len(x) }
func (x Orders) Less(i, j int) bool { return x[i].PendedTime > x[j].PendedTime }
func (x Orders) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type Historys []*History

func (x Historys) Len() int           { return len(x) }
func (x Historys) Less(i, j int) bool { return x[i].Time > x[j].Time }
func (x Historys) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
