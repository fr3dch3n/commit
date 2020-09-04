package input

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// CommitConfig contains two configuration-parameters.
type CommitConfig struct {
	// Abbreviation specifies the personal abbreviation.
	Abbreviation string `json:"abbreviation"`

	// Story mode defines if you will be asked for the current story.
	StoryMode string `json:"storyMode"`

	// TeamMembersConfigPath specifies the path to the team-members-configuration-file.
	TeamMembersConfigPath string `json:"teamMembersConfigPath"`
}

func (c *CommitConfig) String() string {
	return fmt.Sprintf("Abbreviation: %s, StoryMode: %s, TeamMembersConfigPath: %s", c.Abbreviation, c.StoryMode, c.TeamMembersConfigPath)
}

func (c *CommitConfig) containsMinimalSet() error {
	if c.Abbreviation == "" {
		return errors.New("your abbreviation is not specified")
	}
	if c.TeamMembersConfigPath == "" {
		return errors.New("the teamMembersConfigPath is not specified")
	}
	if c.StoryMode == "" {
		return errors.New("the storyMode is not specified")
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
		StoryMode:             oldConfig.StoryMode,
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
	storyMode, err := GetNonEmpty("Do you work with stories? [yes|no]")
	if err != nil {
		return CommitConfig{}, err
	}
	var storyModeBool = ""
	if storyMode == "yes" {
		storyModeBool = "true"
	} else {
		storyModeBool = "false"
	}

	return CommitConfig{
		Abbreviation:          abbreviation,
		StoryMode:             storyModeBool,
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
