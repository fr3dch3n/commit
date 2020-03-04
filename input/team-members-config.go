package input

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// TeamMember contains all information needed to identify a single team-member.
type TeamMember struct {
	GithubUserName string `json:"username"`
	Email          string `json:"mail"`
	Abbreviation   string `json:"abbreviation"`
}

func (tm TeamMember) String() string {
	return fmt.Sprintf("GithubUsername:%s, Email:%s, Abbreviation: %s", tm.GithubUserName, tm.Email, tm.Abbreviation)
}

func readTeamMembersConfig(path string) ([]TeamMember, error) {
	file, err := ioutil.ReadFile(os.ExpandEnv(path))
	if err != nil {
		return []TeamMember{}, err
	}
	var config []TeamMember
	err = json.Unmarshal([]byte(file), &config)
	if err != nil {
		return []TeamMember{}, err
	}
	return config, nil
}

// WriteTeamMembersConfig writes a list of team-members to the filesystem.
func WriteTeamMembersConfig(path string, tms []TeamMember) error {
	b, err := json.MarshalIndent(tms, "", "	")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(os.ExpandEnv(path), b, 0644)
	return err
}

// InitTeamMembersConfig reads and parses a TeamMembersConfig.
// If not found, it creates on.
func InitTeamMembersConfig(path string) ([]TeamMember, error) {
	var tms []TeamMember
	tms, err := readTeamMembersConfig(path)
	if err != nil {
		tms = []TeamMember{}
		err = WriteTeamMembersConfig(path, tms)
		if err != nil {
			return tms, err
		}
	}
	return tms, nil
}
