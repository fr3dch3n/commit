package input

import (
	"io"
	"strings"
	"testing"
)

func Test_GetInputOrElse(t *testing.T) {
	type args struct {
		ioreader io.Reader
		msg      string
		current  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "simple input",
			args:    args{ioreader: strings.NewReader("override"), msg: "", current: "fallback",},
			want:    "override",
			wantErr: false,
		},
		{
			name:    "no input",
			args:    args{ioreader: strings.NewReader(""), msg: "", current: "fallback",},
			want:    "fallback",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetInputOrElse(tt.args.ioreader, tt.args.msg, tt.args.current)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInputOrElse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetInputOrElse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
