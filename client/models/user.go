package models

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

// following defines structs used for communicate with clients

// restResult defines the response payload for a general REST interface request.
type restResult struct {
	OK  string
	Err string
}

// loginRequest is a object to establish security between client and app.
type loginRequest struct {
	EnrollID     string `protobuf:"bytes,1,opt,name=enrollId" json:"enrollId,omitempty"`
	EnrollSecret string `protobuf:"bytes,2,opt,name=enrollSecret" json:"enrollSecret,omitempty"`
}

// signRequest is a object to sign a file
type signRequest struct {
	EnrollID    string `protobuf:"bytes,1,opt,name=enrollId" json:"enrollId,omitempty"`
	EnrollToken string `protobuf:"bytes,2,opt,name=enrollToken" json:"enrollToken,omitempty"`
	FileName    string `protobuf:"bytes,2,opt,name=fileName" json:"fileName,omitempty"`
	FileContent string `protobuf:"bytes,3,opt,name=fileContent" json:"fileContent,omitempty"`
	FileHash    string `protobuf:"bytes,4,opt,name=fileHash" json:"fileHash,omitempty"`
}

// verifyRequest is a object to verify a signature
type verifyRequest struct {
	EnrollID    string `protobuf:"bytes,1,opt,name=enrollId" json:"enrollId,omitempty"`
	EnrollToken string `protobuf:"bytes,2,opt,name=enrollToken" json:"enrollToken,omitempty"`
	FileContent string `protobuf:"bytes,3,opt,name=fileContent" json:"fileContent,omitempty"`
	FileHash    string `protobuf:"bytes,4,opt,name=fileHash" json:"fileHash,omitempty"`
	Signature   string `protobuf:"bytes,5,opt,name=signature" json:"signature,omitempty"`
}

// signatureRequest is a object to signatures
type signatureRequest struct {
	EnrollID    string `protobuf:"bytes,1,opt,name=enrollId" json:"enrollId,omitempty"`
	EnrollToken string `protobuf:"bytes,2,opt,name=enrollToken" json:"enrollToken,omitempty"`
}

// signatureResponse
type signatureEntity struct {
	FileHash      string `json:"fileHash,omitempty"`
	FileName      string `json:"fileName,omitempty"`
	FileSignature string `json:"fileSignature,omitempty"`
	Timestamp     string `json:"timestamp,omitempty"`
}

// SignatureResponse response of signatures
type SignatureResponse struct {
	OK         string            `json:"OK,omitempty"`
	Error      string            `json:"Error,omitempty"`
	Signatures []signatureEntity `json:"signatures,omitempty"`
}

// Login login
// enrollID: enrollID
// enrollSecret: enrollSecret
func Login(enrollID, enrollSecret string) bool {
	var loginRequest loginRequest
	loginRequest.EnrollID = enrollID
	loginRequest.EnrollSecret = enrollSecret

	reqBody, err := json.Marshal(loginRequest)
	if err != nil {
		return false
	}

	urlstr := getHTTPURL("user/login")
	response, err := performHTTPPost(urlstr, reqBody)
	if err != nil {
		logger.Errorf("Login failed: %v", err)
		return false
	}

	logger.Debugf("Login: url=%v request=%v response=%v", urlstr, string(reqBody), string(response))

	var result restResult
	err = json.Unmarshal(response, &result)
	if err != nil {
		logger.Errorf("Login failed: %v", err)
		return false
	}

	if len(result.OK) == 0 {
		logger.Errorf("Login failed: %v", result.Err)
		return false
	}

	return true
}

// GetAvatarUserid getavatar
func GetAvatarUserid(enrollID string) string {
	var avatar string

	avatar = fmt.Sprintf("/static/img/avatar/%v.jpg", enrollID)

	if _, err := os.Stat(avatar); os.IsNotExist(err) {
		return fmt.Sprintf("/static/img/avatar/%d.jpg", rand.Intn(5))
	}

	return avatar
}
