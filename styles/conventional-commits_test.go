package styles

import (
	"github.com/fr3dch3n/commit/input"
	"testing"
)

func TestBuildConventionalCommitMsg(t *testing.T) {
	type args struct {
		conventionalCommitsInformation ConventionalCommitsInformation
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
				conventionalCommitsInformation: ConventionalCommitsInformation{
					commitType: "fix",
					Scope:      "tf",
				},
				pair: []input.TeamMember{{
					GithubUserName: "pair",
					Email:          "pair@mail.com",
					Abbreviation:   "un",
				}},
				summary:     "i commit things",
				explanation: "Because I can\nAnd I Like it",
			},
			want: "fix(tf): i commit things\n\nBecause I can\nAnd I Like it\n\n\nCo-authored-by: pair <pair@mail.com>\n",
		},
		{
			name: "simple commit without pair",
			args: args{
				conventionalCommitsInformation: ConventionalCommitsInformation{
					commitType: "feat",
					Scope:      "ABC-001",
				},
				pair:        []input.TeamMember{},
				summary:     "i commit things",
				explanation: "Because I can\nAnd I Like it",
			},
			want: "feat(ABC-001): i commit things\n\nBecause I can\nAnd I Like it\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildConventionalCommitMsg(tt.args.conventionalCommitsInformation, tt.args.pair, tt.args.summary, tt.args.explanation); got != tt.want {
				t.Errorf("BuildConventionalCommitMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}
