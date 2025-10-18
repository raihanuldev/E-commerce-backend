package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}
type Payload struct {
	Sub         string `json:"sub"`
	FristName   string `json:"frist_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner string `json:"is_shop_owner"`
}

func CreateJwt(secret string,data Payload) (string,error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	headerByteArr, err := json.Marshal(header)
	if(err!=nil){
		return "",err
	}
	base64Header:= Base64URLEncode(headerByteArr)
	byteArrData,err := json.Marshal(data)
	if(err!=nil){
		return "",err
	}
	base64Payload:=Base64URLEncode(byteArrData)
	message:= base64Header+"."+base64Payload

	secretByteArr:= []byte(secret)
	messagebyteArr:= []byte(message)

	h:=hmac.New(sha256.New,secretByteArr)
	h.Write(messagebyteArr)

	signure:=h.Sum(nil)

	signureBase64 := Base64URLEncode(signure)

	jwt:= base64Header+"."+base64Payload+"."+signureBase64
	return jwt,nil;

}

func Base64URLEncode(data []byte)string{
	 return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}