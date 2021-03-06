package input

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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
				path: "../test-resources/commit-config/simple-config.json",
			},
			want: CommitConfig{
				TeamMembersConfigPath: "path",
				CommitStyle:           "conventional",
			},
			wantErr: false,
		},
		{
			name: "read empty test-config",
			args: args{
				path: "../test-resources/commit-config/empty-config.json",
			},
			want:    CommitConfig{},
			wantErr: false,
		},
		{
			name: "path does not exist",
			args: args{
				path: "../test-resources/commit-config/nonexistent-config.json",
			},
			want:    CommitConfig{},
			wantErr: true,
		},
		{
			name: "broken config file",
			args: args{
				path: "../test-resources/commit-config/broken-config.json",
			},
			want:    CommitConfig{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readCommitConfig(tt.args.path)
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

func Test_ContainsMinimalSet(t *testing.T) {
	type args struct {
		c CommitConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "empty config",
			args: args{
				c: CommitConfig{
					TeamMembersConfigPath: "",
				},
			},
			wantErr: true,
		},
		{
			name: "only abbreviation is missing",
			args: args{
				c: CommitConfig{
					TeamMembersConfigPath: "/some/path",
				},
			},
			wantErr: true,
		},
		{
			name: "only config-path is missing",
			args: args{
				c: CommitConfig{
					TeamMembersConfigPath: "",
				},
			},
			wantErr: true,
		},
		{
			name: "full config",
			args: args{
				c: CommitConfig{
					CommitStyle:           "conventional",
					TeamMembersConfigPath: "/some/path",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.args.c.containsMinimalSet(); (err != nil) != tt.wantErr {
				t.Errorf("ContainsMinimalSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_WriteCommitConfig(t *testing.T) {
	type args struct {
		path      string
		pair      TeamMember
		story     string
		oldConfig CommitConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "write full config",
			args: args{
				path: "../test-resources/test-output/out-commit-config.json",
				pair: TeamMember{
					GithubUserName: "member1",
					Email:          "member1@company.com",
					Abbreviation:   "m1",
				},
				story: "TR-410",
				oldConfig: CommitConfig{
					TeamMembersConfigPath: "test-resources/commit-config/no-longer-existent-config.json",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := ioutil.TempFile("", "tmp-commit-config-")
			if err := writeCommitConfig(file.Name(), tt.args.oldConfig); (err != nil) != tt.wantErr {
				t.Errorf("WriteCommitConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
			config, err := readCommitConfig(file.Name())
			if err != nil {
				t.Errorf("ReadCommitConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, tt.args.oldConfig.TeamMembersConfigPath, config.TeamMembersConfigPath)
			_ = os.Remove(file.Name())
		})
	}
}
