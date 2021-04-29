package notify

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/kirinlabs/HttpRequest"
)

var request = HttpRequest.NewRequest()

func hmacSha256(stringToSign string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func HTTPRequest(url string, token string) *simplejson.Json {
	headers := map[string]string{
		"Authorization": fmt.Sprintf("token %s", token),
	}
	request.SetHeaders(headers)
	res, err := request.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, _ := res.Body()
	js, err := simplejson.NewJson([]byte(body))
	if err != nil || js == nil {
		log.Fatal("something wrong when call NewFromReader")
	}
	return js
}

func SendDingDingMessage(content string, dingdingURL string, secretKey string) (bool, error) {
	headers := map[string]string{
		"Content-Type": "application/json;charset=utf-8",
	}
	timestamp := time.Now().UnixNano() / 1e6
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, secretKey)
	sign := hmacSha256(stringToSign, secretKey)

	postURL := fmt.Sprintf("%s&timestamp=%d&sign=%s", dingdingURL, timestamp, sign)
	request.SetHeaders(headers)
	res, err := request.Post(postURL, bytes.NewBuffer([]byte(content)))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

	body, _ := res.Body()
	if res.StatusCode() == 200 {
		return true, nil
	}
	return false, errors.New(string(body))
}
