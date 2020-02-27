package input

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type CommitConfig struct {
	Short                 string `json:"short"`
	TeamMembersConfigPath string `json:"teamMembersConfigPath"`
}

func (c CommitConfig) String() string {
	return fmt.Sprintf("Short: %s, TeamMembersConfigPath: %s", c.Short, c.TeamMembersConfigPath)
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
	if c.Short == "" {
		return errors.New("your short-name is not specified")
	} else if c.TeamMembersConfigPath == "" {
		return errors.New("the teamMembersConfigPath is not specified")
	}
	return nil
}

func WriteCommitConfig(path string, oldConfig CommitConfig) error {
	newConfig := CommitConfig{
		Short:                 oldConfig.Short,
		TeamMembersConfigPath: oldConfig.TeamMembersConfigPath,
	}
	b, err := json.MarshalIndent(newConfig, "", "	")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, b, 0644)
	return err
}
