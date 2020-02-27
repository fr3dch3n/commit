package input

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type TeamMember struct {
	GithubUserName string `json:"username"`
	Email          string `json:"mail"`
	Abbreviation   string `json:"abbreviation"`
}

func (tm TeamMember) String() string {
	return fmt.Sprintf("GithubUsername:%s, Email:%s, Abbreviation: %s", tm.GithubUserName, tm.Email, tm.Abbreviation)
}

func ReadTeamMembersConfig(path string) ([]TeamMember, error) {
	file, err := ioutil.ReadFile(path)
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

func WriteTeamMembersConfig(path string, tms []TeamMember) error {
	b, err := json.MarshalIndent(tms, "", "	")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, b, 0644)
	return err
}

func InitTeamMembersConfig(path string) ([]TeamMember, error) {
	var tms []TeamMember
	tms, err  := ReadTeamMembersConfig(path)
	if err != nil {
		tms = []TeamMember{}
		err = WriteTeamMembersConfig(path, tms)
		if err != nil {
			return tms, err
		}
	}
	return tms, nil
}
