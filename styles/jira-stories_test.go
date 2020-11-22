package styles

import (
	"github.com/fr3dch3n/commit/input"
	"testing"
)

func TestBuildStoryStyleCommitMsg(t *testing.T) {
	type args struct {
		jiraStoryInformation JIRAStoryInformation
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
				jiraStoryInformation: JIRAStoryInformation{
					Story: "FTX-123",
				},
				pair: []input.TeamMember{{
					GithubUserName: "pair",
					Email:          "pair@mail.com",
					Abbreviation:   "un",
				}},
				summary:     "i commit things",
				explanation: "Because I can\nAnd I Like it",
			},
			want: "[FTX-123] i commit things\n\nBecause I can\nAnd I Like it\n\n\nCo-authored-by: pair <pair@mail.com>\n",
		},
		{
			name: "simple commit without pair",
			args: args{
				jiraStoryInformation: JIRAStoryInformation{
					Story: "ABC-001",
				},
				pair:        []input.TeamMember{},
				summary:     "i commit things",
				explanation: "Because I can\nAnd I Like it",
			},
			want: "[ABC-001] i commit things\n\nBecause I can\nAnd I Like it\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildStoryStyleCommitMsg(tt.args.jiraStoryInformation, tt.args.pair, tt.args.summary, tt.args.explanation); got != tt.want {
				t.Errorf("BuildStoryStyleCommitMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}
