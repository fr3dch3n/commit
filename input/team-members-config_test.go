package input

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func Test_ReadTeamMembersConfig(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []TeamMember
		wantErr bool
	}{
		{
			name: "read simple test-config",
			args: args{
				path: "../test-resources/team-members-config/simple-config.json",
			},
			want: []TeamMember{
				{
					GithubUserName: "member1",
					Email:          "member1@company.com",
					Abbreviation:   "m1",
				},
				{
					GithubUserName: "member2",
					Email:          "member2@company.com",
					Abbreviation:   "m2",
				},
			},
			wantErr: false,
		},
		{
			name: "read empty test-config",
			args: args{
				path: "../test-resources/team-members-config/empty-config.json",
			},
			want:    []TeamMember{},
			wantErr: false,
		},
		{
			name: "path does not exist",
			args: args{
				path: "../test-resources/team-members-config/nonexistent-config.json",
			},
			want:    []TeamMember{},
			wantErr: true,
		},
		{
			name: "broken config file",
			args: args{
				path: "../test-resources/team-members-config/broken-config.json",
			},
			want:    []TeamMember{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadTeamMembersConfig(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadTeamMembersConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadTeamMembersConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_WriteTeamMembersConfig(t *testing.T) {
	type args struct {
		path string
		tms  []TeamMember
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "write empty list of team members",
			args: args{
				path: "",
				tms:  []TeamMember{},
			},
			wantErr: false,
		},
		{
			name: "write filled list of team members",
			args: args{
				path: "",
				tms: []TeamMember{
					{
						GithubUserName: "member1",
						Email:          "m1@company.com",
						Abbreviation:   "m1",
					},
					{
						GithubUserName: "member2",
						Email:          "m2@company.com",
						Abbreviation:   "m2",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := ioutil.TempFile("", "tmp-team-members-config-")
			if err := WriteTeamMembersConfig(file.Name(), tt.args.tms); (err != nil) != tt.wantErr {
				t.Errorf("WriteTeamMembersConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
			tms, err := ReadTeamMembersConfig(file.Name())
			if err != nil {
				t.Errorf("ReadTeamMembersConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, tt.args.tms, tms)
			_ = os.Remove(file.Name())
		})
	}
}
