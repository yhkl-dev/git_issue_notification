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
				dingdingURL: "https://oapi.dingtalk.com/robot/send?access_token=c471ab384eee2fc8f93ff623cb771320fc34761fbff68c19a915c57515b9decd",
				secretKey:   "SEC97a57299528ce2827345d2d20a8180209bbbfc3157eac28a5045c2b000241e38",
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
