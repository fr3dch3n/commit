package input

import (
	"encoding/json"
	"io/ioutil"
)

type State struct {
	CurrentStory string `json:"story"`
	CurrentPair  string `json:"pair"`
}


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

func WriteState(path string, pair TeamMember, story string) error {
	newState := State{
		CurrentStory: story,
		CurrentPair:  pair.Short,
	}
	b, err := json.MarshalIndent(newState, "", "	")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, b, 0644)
	return err
}

