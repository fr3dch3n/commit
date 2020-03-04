package git

import (
	"errors"
	"fmt"

	"github.com/fr3dch3n/commit/input"
	log "github.com/sirupsen/logrus"
)

// GetPair tries to get the current pairing-partner from the user-input.
func GetPair(commitConfig input.CommitConfig, currentPair string, teamMembers []input.TeamMember) (input.TeamMember, error) {
	var pair input.TeamMember
	var pairAbbreviation string
	var err error
	if currentPair == "none" {
		pairAbbreviation, err = input.Get("Current pairing partner")
	} else {
		pairAbbreviation, err = input.GetWithDefault("Current pairing partner", currentPair)
	}
	if err != nil {
		return input.TeamMember{}, err
	}
	log.Debug("PairAbbreviation: " + pairAbbreviation)
	if pairAbbreviation == "" {
		return input.TeamMember{Abbreviation: "none"}, nil
	}

	pair, err = GetTeamMemberByAbbreviation(teamMembers, pairAbbreviation)
	if err != nil && err.Error() == "not-found" {
		newMember, err := GetAndSaveNewTeamMember(commitConfig.TeamMembersConfigPath, pairAbbreviation, teamMembers)
		if err != nil {
			return input.TeamMember{}, err
		}
		pair = newMember
	} else if err != nil {
		return input.TeamMember{}, err
	}

	return pair, nil
}

// GetTeamMemberByAbbreviation makes a lookup for a team-member by the abbreviation.
func GetTeamMemberByAbbreviation(tms []input.TeamMember, abbreviation string) (input.TeamMember, error) {
	var found input.TeamMember
	for _, tm := range tms {
		if tm.Abbreviation == abbreviation {
			found = tm
		}
	}
	if (input.TeamMember{}) == found {
		return input.TeamMember{}, errors.New("not-found")
	}
	return found, nil
}

// GetAndSaveNewTeamMember runs GetNewTeamMemberFromInput and then saves the result in the team-members file.
func GetAndSaveNewTeamMember(path string, abbreviation string, tms []input.TeamMember) (input.TeamMember, error) {
	newMember, err := getNewTeamMemberFromInput(abbreviation)
	if err != nil {
		return input.TeamMember{}, err
	}
	err = input.WriteTeamMembersConfig(path, append(tms, newMember))
	if err != nil {
		return input.TeamMember{}, err
	}
	return newMember, nil
}

func getNewTeamMemberFromInput(abbreviation string) (input.TeamMember, error) {
	fmt.Println("Creating team-member with abbreviation " + abbreviation)
	username, err := input.GetNonEmpty("Enter username")
	if err != nil {
		return input.TeamMember{}, nil
	}
	mail, err := input.GetNonEmpty("Enter mail")
	if err != nil {
		return input.TeamMember{}, nil
	}

	return input.TeamMember{
		GithubUserName: username,
		Email:          mail,
		Abbreviation:   abbreviation,
	}, nil
}
