package input

import (
	"reflect"
	"testing"
)

func Test_ReadCommitConfig(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    CommitConfig
		wantErr bool
	}{
		{
			name: "read simple test-config",
			args: args{
				path: "../test-resources/simple-commit-config.json",
			},
			want: CommitConfig{
				GithubUsername:        "my_username",
				CurrentStory:          "ABC-001",
				CurrentPair:           "abc",
				TeamMembersConfigPath: "path",
				Short:                 "abc",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadCommitConfig(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadCommitConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadCommitConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
