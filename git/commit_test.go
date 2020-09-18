package git

import (
	"github.com/fr3dch3n/commit/input"
	"testing"
)

func TestBuildCommitMsg(t *testing.T) {
	type args struct {
		ctype        string
		scope        string
		pair         []input.TeamMember
		summary      string
		explanation  string
		abbreviation input.TeamMember
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first simple commit",
			args: args{
				ctype: "fix",
				scope: "tf",
				pair: []input.TeamMember{{
					GithubUserName: "pair",
					Email:          "pair@mail.com",
					Abbreviation:   "un",
				}},
				summary:     "i commit things",
				explanation: "Because I can\nAnd I Like it",
				abbreviation: input.TeamMember{
					GithubUserName: "myself",
					Email:          "me@company.com",
					Abbreviation:   "me",
				},
			},
			want: "fix(tf): i commit things\n\nBecause I can\nAnd I Like it\n\n\nCo-authored-by: pair <pair@mail.com>\n",
		},
		{
			name: "simple commit without pair",
			args: args{
				ctype:       "feat",
				scope:       "ABC-001",
				pair:        []input.TeamMember{},
				summary:     "i commit things",
				explanation: "Because I can\nAnd I Like it",
				abbreviation: input.TeamMember{
					GithubUserName: "myself",
					Email:          "me@company.com",
					Abbreviation:   "me",
				},
			},
			want: "feat(ABC-001): i commit things\n\nBecause I can\nAnd I Like it\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildCommitMsg(tt.args.ctype, tt.args.scope, tt.args.pair, tt.args.summary, tt.args.explanation, tt.args.abbreviation); got != tt.want {
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
