package input

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// CommitConfig contains two configuration-parameters.
type CommitConfig struct {
	// Story mode defines if you will be asked for the current story.
	CommitStyle string `json:"commitStyle"`

	// TeamMembersConfigPath specifies the path to the team-members-configuration-file.
	TeamMembersConfigPath string `json:"teamMembersConfigPath"`
}

func (c *CommitConfig) String() string {
	return fmt.Sprintf("CommitStyle: %s, TeamMembersConfigPath: %s", c.CommitStyle, c.TeamMembersConfigPath)
}

func (c *CommitConfig) containsMinimalSet() error {
	if c.TeamMembersConfigPath == "" {
		return errors.New("the teamMembersConfigPath is not specified")
	}
	if c.CommitStyle == "" {
		return errors.New("the commitStyle is not specified")
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
		CommitStyle:           oldConfig.CommitStyle,
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
	teamMembersConfigPath, err := GetNonEmpty("Enter the teamMembersConfigPath")
	if err != nil {
		return CommitConfig{}, err
	}
	storyMode, err := GetNonEmpty("Do you want to work with conventional-commits or jira-style? [conventional|jira]")
	if err != nil {
		return CommitConfig{}, err
	}
	var storyModeBool = ""
	if storyMode == "conventional" {
		storyModeBool = "conventional"
	} else {
		storyModeBool = "jira"
	}

	return CommitConfig{
		CommitStyle:           storyModeBool,
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
