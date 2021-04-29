package notify

import "testing"

func TestHTTPRequest(t *testing.T) {
	type args struct {
		url   string
		token string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "yhkl",
			args: args{url: "https://api.github.com/user/repos?page=1&per_page=1000", token: "419ed57e343e23a1dae83799e497fe0392c0ef22"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HTTPRequest(tt.args.url, tt.args.token)
		})
	}
}
