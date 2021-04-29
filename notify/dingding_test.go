package notify

import (
	"fmt"
	"testing"
)

func TestSendDingDingMessage(t *testing.T) {
	type args struct {
		content     string
		dingdingURL string
		secretKey   string
	}
	data := `{"msgtype":"markdown","markdown":{"title":"系统异常","text":"%s"},"at":{"atMobiles":[],"isAtAll":false}}`
	fmtDat := fmt.Sprintf(data, "content")
	tests := []struct {
		name string
		args args
	}{
		{name: "yhkl",
			args: args{
				content:     fmtDat,
				dingdingURL: "https://oapi.dingtalk.com/robot/send?access_token=ac48ccf5d20ba41b11bf6c335a0afb2567df312eac50a34fb0d31b643d3c6c89",
				secretKey:   "SEC873c3a1ea4f0e26c391f4e294cb1a63893c47cb6023ec6381d3306f7b979e0f5",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			respFlag, err := SendDingDingMessage(tt.args.content, tt.args.dingdingURL, tt.args.secretKey)
			if respFlag == false || err != nil {
				fmt.Println("error send message")
				fmt.Println(err)
			}
			fmt.Println(respFlag)
		})
	}
}
