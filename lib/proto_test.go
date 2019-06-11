package lib

import (
	"reflect"
	"testing"
    "fmt"
)

func TestDecodeErrorMessage(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    ErrorMessage
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "default",
			args: args{
				data: []byte("{\"host\":\"gateway\",\"server\":\"profile\",\"users\":[\"lvfei\",\"kkfnui\"],\"error_id\":\"sha1\",\"message\":\"test is test\",\"detail\":\"no detail\"}"),
			},
			want: ErrorMessage{
				Host:    "gateway",
				Server:  "profile",
				Users:   []string{"lvfei", "kkfnui"},
				ErrorID: "sha1",
				Message: "test is test",
				Detail:  "no detail",
			},
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				data: []byte("{\"host1\":gateway\",\"server\":\"profile\",\"users\":[\"lvfei\",\"kkfnui\"],\"error_id\":\"sha1\",\"message\":\"test is test\",\"detail\":\"no detail\"}"),
			},
			want:    ErrorMessage{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
            fmt.Printf("test messsage:%s\n", tt.args.data)
			got, err := DecodeErrorMessage(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeErrorMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeErrorMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeErrorMessage(t *testing.T) {
	type args struct {
		message ErrorMessage
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "default",
			args: args{
				message: ErrorMessage{
					Host:    "gateway",
					Server:  "profile",
					Users:   []string{"lvfei", "kkfnui"},
					ErrorID: "sha1",
					Message: "test is test",
					Detail:  "no detail\ntest",
				},
			},
			want:    "{\"host\":\"gateway\",\"server\":\"profile\",\"users\":[\"lvfei\",\"kkfnui\"],\"error_id\":\"sha1\",\"message\":\"test is test\",\"detail\":\"no detail\\ntest\"}",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncodeErrorMessage(tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeErrorMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EncodeErrorMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
