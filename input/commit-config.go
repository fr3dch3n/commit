package input

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// CommitConfig contains two configuration-parameters.
// Abbreviation specifies the personal abbreviation.
// TeamMembersConfigPath specifies the path to the team-members-configuration-file.
type CommitConfig struct {
	Abbreviation          string `json:"abbreviation"`
	TeamMembersConfigPath string `json:"teamMembersConfigPath"`
}

func (c *CommitConfig) String() string {
	return fmt.Sprintf("Abbreviation: %s, TeamMembersConfigPath: %s", c.Abbreviation, c.TeamMembersConfigPath)
}

func (c *CommitConfig) containsMinimalSet() error {
	if c.Abbreviation == "" {
		return errors.New("your abbreviation is not specified")
	}
	if c.TeamMembersConfigPath == "" {
		return errors.New("the teamMembersConfigPath is not specified")
	}
	return nil
}

func readCommitConfig(path string) (CommitConfig, error) {
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

func writeCommitConfig(path string, oldConfig CommitConfig) error {
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

func getFromInput() (CommitConfig, error) {
	abbreviation, err := GetNonEmpty("Enter your abbreviation")
	if err != nil {
		return CommitConfig{}, err
	}
	teamMembersConfigPath, err := GetNonEmpty("Enter the teamMembersConfigPath")
	if err != nil {
		return CommitConfig{}, err
	}
	return CommitConfig{
		Abbreviation:          abbreviation,
		TeamMembersConfigPath: teamMembersConfigPath,
	}, nil
}

// InitCommitConfig reads and validates the commit-configuration and returns the corresponding object or error.
func InitCommitConfig(path string) (CommitConfig, error) {
	var config CommitConfig
	fromFS, err := readCommitConfig(path)
	if err != nil {
		config, err = getFromInput()
		if err != nil {
			return config, err
		}
		err = writeCommitConfig(path, config)
		if err != nil {
			return config, err
		}
		return config, nil
	}
	if err = fromFS.containsMinimalSet(); err != nil {
		config, err = getFromInput()
		if err != nil {
			return config, err
		}
		err = writeCommitConfig(path, config)
		if err != nil {
			return config, err
		}
		return config, nil

	}
	return fromFS, nil
}
