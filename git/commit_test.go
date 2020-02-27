package git

import (
	"github.com/fr3dch3n/commit/input"
	"testing"
)

func TestBuildCommitMsg(t *testing.T) {
	type args struct {
		story       string
		pair        input.TeamMember
		summary     string
		explanation string
		short       input.TeamMember
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first simple commit",
			args: args{
				story: "ABC-001",
				pair: input.TeamMember{
					GithubUserName: "pair",
					Email:          "pair@mail.com",
					Short:          "un",
				},
				summary:     "I commit things",
				explanation: "Because I can\nAnd I Like it",
				short:       input.TeamMember{
					GithubUserName: "myself",
					Email:          "me@company.com",
					Short:          "me",
				},
			},
			want: "[ABC-001] un|me I commit things\n\nBecause I can\nAnd I Like it\n\n\nCo-authored-by: pair <pair@mail.com>\n",
		},
		{
			name: "simple commit without pair",
			args: args{
				story: "ABC-001",
				pair: input.TeamMember{
					Short: "none",
				},
				summary:     "I commit things",
				explanation: "Because I can\nAnd I Like it",
				short:       input.TeamMember{
					GithubUserName: "myself",
					Email:          "me@company.com",
					Short:          "me",
				},
			},
			want: "[ABC-001] me I commit things\n\nBecause I can\nAnd I Like it\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildCommitMsg(tt.args.story, tt.args.pair, tt.args.summary, tt.args.explanation, tt.args.short, false); got != tt.want {
				t.Errorf("BuildCommitMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			name: "make first char capital",
			args: args{
				summary: "word",
			},
			want: "Word",
		},
		{
			name: "remove whitespace",
			args: args{
				summary: " word ",
			},
			want: "Word",
		},
		{
			name: "multiple words are fine",
			args: args{
				summary: " a few words ",
			},
			want: "A few words",
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
