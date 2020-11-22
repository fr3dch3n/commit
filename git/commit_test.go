package git

import (
	"testing"
)

func TestReviewSummary(t *testing.T) {
	type args struct {
		summary string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "remove whitespace",
			args: args{
				summary: " word ",
			},
			want: "word",
		},
		{
			name: "multiple words are fine",
			args: args{
				summary: " a few words ",
			},
			want: "a few words",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReviewSummary(tt.args.summary); got != tt.want {
				t.Errorf("ReviewSummary() = %v, want %v", got, tt.want)
			}
		})
	}
}
