package git

import (
	"errors"
	"github.com/fr3dch3n/commit/input"
	log "github.com/sirupsen/logrus"
	"os"
)

func GetPair(commitConfig input.CommitConfig, currentPair string, teamMembers []input.TeamMember) (input.TeamMember, error) {
	var pair input.TeamMember
	pairAbbreviation, err := input.GetInputOrElse(os.Stdin, "Pairing with", currentPair)
	if err != nil {
		return input.TeamMember{}, err
	}
	log.Debug("PairAbbreviation: " + pairAbbreviation)
	if pairAbbreviation == "none" {
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

func GetAndSaveNewTeamMember(path string, abbreviation string, tms []input.TeamMember) (input.TeamMember, error) {
	newMember, err := input.GetNewTeamMemberFromInput(os.Stdin, abbreviation)
	if err != nil {
		return input.TeamMember{}, err
	}
	err = input.WriteTeamMembersConfig(path, append(tms, newMember))
	if err != nil {
		return input.TeamMember{}, err
	}
	return newMember, nil
}
