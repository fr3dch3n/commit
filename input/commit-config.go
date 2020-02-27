package input

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type CommitConfig struct {
	Abbreviation          string `json:"abbreviation"`
	TeamMembersConfigPath string `json:"teamMembersConfigPath"`
}

func (c *CommitConfig) String() string {
	return fmt.Sprintf("Abbreviation: %s, TeamMembersConfigPath: %s", c.Abbreviation, c.TeamMembersConfigPath)
}

func (c *CommitConfig) ContainsMinimalSet() error {
	if c.Abbreviation == "" {
		return errors.New("your abbreviation is not specified")
	}
	if c.TeamMembersConfigPath == "" {
		return errors.New("the teamMembersConfigPath is not specified")
	}
	return nil
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

func GetFromInput(ioreader io.Reader) (CommitConfig, error) {
	abbreviation, err := GetNonEmptyInput(ioreader, "Enter your abbreviation")
	if err != nil {
		return CommitConfig{}, err
	}
	teamMembersConfigPath, err := GetNonEmptyInput(ioreader, "Enter the teamMembersConfigPath")
	if err != nil {
		return CommitConfig{}, err
	}
	return CommitConfig{
		Abbreviation:          abbreviation,
		TeamMembersConfigPath: teamMembersConfigPath,
	}, nil
}

func InitCommitConfig(path string) (CommitConfig, error) {
	var config CommitConfig
	fromFS, err  := ReadCommitConfig(path)
	if err != nil {
		config, err = GetFromInput(os.Stdin)
		if err != nil {
			return config, err
		}
		err = WriteCommitConfig(path, config)
		if err != nil {
			return config, err
		}
		return config, nil
	}
	if err = fromFS.ContainsMinimalSet(); err != nil {
		config, err = GetFromInput(os.Stdin)
		if err != nil {
			return config, err
		}
		err = WriteCommitConfig(path, config)
		if err != nil {
			return config, err
		}
		return config, nil

	}
	return config, nil
}
