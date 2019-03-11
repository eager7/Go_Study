package dapp_sheild

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"time"
)

//const URI = "/v1/empowerings/list?net=kylin"
const URI = "/v1/empowerings/list?net=eos"
const TimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"

type Dapp struct {
	Net       string   `json:"net"`
	Name      string   `json:"name"`
	Contracts []string `json:"contracts"`
}

type Result struct {
	Dapps  []Dapp `json:"dapps"`
	Offset int    `json:"offset"`
}

type RespContracts struct {
	Result Result `json:"result"`
}

func RequestContractList() (*RespContracts, error) {
	var date = time.Now().UTC().Format(TimeFormat)
	auth, err := Authorization("eospark", "GET", URI, "", "", date)
	if err != nil {
		return nil, err
	}
	result := RespContracts{}
	res := gorequest.New()
	ret, resp, errs := res.Timeout(time.Second*2).Get("https://api.dappshield.io"+URI).Set("Date", date).Set("Authorization", auth).EndStruct(&result)
	if errs != nil || ret.StatusCode != http.StatusOK {
		e := fmt.Sprintf("request error:%s", errs)
		req, err := res.MakeRequest()
		if err == nil {
			fmt.Printf("request body:%+v\n", req)
		}
		fmt.Println("resp body:", string(resp))
		return nil, errors.New(e)
	}
	return &result, nil
}

func Authorization(operator, method, uri, contentMd5, password, date string) (string, error) {
	h := md5.Sum([]byte(password))
	hMac := hmac.New(sha1.New, []byte(hex.EncodeToString(h[:])))
	signMsg := method + "&" + uri + "&" + date + "&" + contentMd5
	fmt.Println("sign message:", signMsg)
	if _, err := hMac.Write([]byte(signMsg)); err != nil {
		return "", err
	}
	Signature := base64.StdEncoding.EncodeToString(hMac.Sum(nil))

	sign := "DAppShield " + operator + ":" + Signature
	fmt.Println("date:", date)
	fmt.Println("auth:", sign)
	return sign, nil
}
