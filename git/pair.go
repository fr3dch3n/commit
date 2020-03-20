package git

import (
	"errors"
	"fmt"
	"strings"

	"github.com/fr3dch3n/commit/input"
	log "github.com/sirupsen/logrus"
)

// GetPair tries to get the current pairing-partner from the user-input.
func GetPair(commitConfig input.CommitConfig, currentPair []string, teamMembers []input.TeamMember) ([]input.TeamMember, error) {
	var pairAbbreviation string
	var err error
	if len(currentPair) == 0 {
		pairAbbreviation, err = input.Get("Current pairing partner (separate by [,| ])")
	} else {
		pairAbbreviation, err = input.GetWithDefault("Current pairing partner (separate by [,| ])", strings.Join(currentPair, ","))
	}
	if err != nil {
		return []input.TeamMember{}, err
	}
	log.Debug("PairAbbreviation: " + pairAbbreviation)

	cutAbbrevs := SeparateAbbreviation(pairAbbreviation)
	log.Debugf("cutAbbrevs: %v", cutAbbrevs)

	if len(cutAbbrevs) == 0 {
		return []input.TeamMember{}, nil
	}

	var pair []input.TeamMember
	for _, abbrev := range cutAbbrevs {
		member, err := GetTeamMemberByAbbreviation(teamMembers, abbrev)
		if err != nil && err.Error() == "not-found" {
			newMember, err := GetAndSaveNewTeamMember(commitConfig.TeamMembersConfigPath, abbrev, teamMembers)
			if err != nil {
				log.Debug("GetAndSaveNewTeamMember: ", err)
				continue
			}
			pair = append(pair, newMember)
		} else if err != nil {
			log.Debug("GetTeamMemberByAbbreviation: ", err)
			continue
		} else {
			pair = append(pair, member)
		}
	}
	return pair, nil
}

// SeparateAbbreviation takes a string and splits it by [,| ] and returns a list of existing strings
func SeparateAbbreviation(input string) []string {
	var result []string
	splitByComma := strings.Split(input, ",")
	for _, x := range splitByComma {
		splitByWhiteSpace := strings.Split(x, " ")
		for _, y := range splitByWhiteSpace {
			splitByPipe := strings.Split(y, "|")
			for _, z := range splitByPipe {
				if z != "" && len(z) > 1 {
					result = append(result, z)
				}
			}
		}
	}
	return result
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
