package input

import (
	"encoding/json"
	"io/ioutil"
)

// State contains two state-parameters.
// CurrentStory specifies the last saved story.
// CurrentPair specifies the last saved pairing-partner.
type State struct {
	CurrentStory string   `json:"story"`
	CurrentPair  []string `json:"pair"`
}

// ReadState reads and parses the state from the path.
func ReadState(path string) (State, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return State{}, nil
	}
	config := State{}
	err = json.Unmarshal([]byte(file), &config)
	if err != nil {
		return State{}, err
	}
	return config, nil
}

// WriteState marshalls and saves the new pair and story to filesystem.
func WriteState(path string, pair []TeamMember, story string) error {
	var pairAbbrev []string
	for _, p := range pair {
		pairAbbrev = append(pairAbbrev, p.Abbreviation)
	}
	newState := State{
		CurrentStory: story,
		CurrentPair:  pairAbbrev,
	}
	b, err := json.MarshalIndent(newState, "", "	")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, b, 0644)
	return err
}
