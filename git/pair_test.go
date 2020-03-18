package git

import (
	"reflect"
	"testing"
)

func Test_separateAbbreviation(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "single pair",
			args: args{
				input: "mem1",
			},
			want: []string{"mem1"},
		},
		{
			name: "pair by comma",
			args: args{
				input: "mem1,mem2",
			},
			want: []string{"mem1", "mem2"},
		},
		{
			name: "pair by whitespace",
			args: args{
				input: "mem1 mem2",
			},
			want: []string{"mem1", "mem2"},
		},
		{
			name: "pair by comma and whitespace",
			args: args{
				input: " mem1 , mem2 ",
			},
			want: []string{"mem1", "mem2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SeparateAbbreviation(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SeparateAbbreviation() = %v, want %v", got, tt.want)
			}
		})
	}
}
