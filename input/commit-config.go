package input

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type CommitConfig struct {
	Abbreviation          string `json:"abbreviation"`
	TeamMembersConfigPath string `json:"teamMembersConfigPath"`
}

func (c CommitConfig) String() string {
	return fmt.Sprintf("Abbreviation: %s, TeamMembersConfigPath: %s", c.Abbreviation, c.TeamMembersConfigPath)
}

func ReadCommitConfig(path string) (CommitConfig, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return CommitConfig{}, err
	}
	config := CommitConfig{}
	err = json.Unmarshal([]byte(file), &config)
	if err != nil {
		return CommitConfig{}, err
	}
	return config, nil
}

func ContainsMinimalSet(c CommitConfig) error {
	if c.Abbreviation == "" {
		return errors.New("your abbreviation-name is not specified")
	} else if c.TeamMembersConfigPath == "" {
		return errors.New("the teamMembersConfigPath is not specified")
	}
	return nil
}

func WriteCommitConfig(path string, oldConfig CommitConfig) error {
	newConfig := CommitConfig{
		Abbreviation:          oldConfig.Abbreviation,
		TeamMembersConfigPath: oldConfig.TeamMembersConfigPath,
	}
	b, err := json.MarshalIndent(newConfig, "", "	")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, b, 0644)
	return err
}
