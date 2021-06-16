package lib

import "testing"

func TestWechatClient_SendMessage(t *testing.T) {
	type fields struct {
		CorpID        string
		CorpSecret    string
		AgentID       int
		tokenExpireAt int64
		accessToken   string
	}
	type args struct {
		users     []string
		robotUrls []string
		msg       string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "default",
			fields: fields{
				CorpID:     "wx380fabtest5f",
				CorpSecret: "inxHHpD46tVIb_teste9msSqIA5SqWRs",
				AgentID:    1000007,
			},
			args: args{
				users: []string{"test"},
				msg:   "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &WechatChannel{
				CorpID:        tt.fields.CorpID,
				CorpSecret:    tt.fields.CorpSecret,
				AgentID:       tt.fields.AgentID,
				tokenExpireAt: tt.fields.tokenExpireAt,
				accessToken:   tt.fields.accessToken,
			}
			if err := client.SendMessage(tt.args.users, tt.args.robotUrls, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("WechatChannel.SendMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
