package notify

import (
	"fmt"
	"log"

	"github.com/bitly/go-simplejson"
	"github.com/kirinlabs/HttpRequest"
)

var request = HttpRequest.NewRequest()

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
