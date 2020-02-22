package git

import (
	"github.com/fr3dch3n/commit/input"
	log "github.com/sirupsen/logrus"
	"os"
)

func GetPair(commitConfig input.CommitConfig, teamMembers []input.TeamMember) (input.TeamMember, error) {
	var pair input.TeamMember
	pairShort, err := input.GetInputOrElse(os.Stdin, "Pairing with", commitConfig.CurrentPair)
	if err != nil {
		return input.TeamMember{}, err
	}
	log.Debug("PairShort: " + pairShort)
	if pairShort == "none" {
		return input.TeamMember{}, nil
	}
	for _, tm := range teamMembers {
		if tm.Short == pairShort {
			pair = tm
		}
	}
	if (input.TeamMember{}) == pair {
		newMember, err := input.GetNewTeamMemberFromInput(os.Stdin)
		if err != nil {
			return input.TeamMember{}, err
		}
		err = input.WriteTeamMembersConfig(commitConfig.TeamMembersConfigPath, append(teamMembers, newMember))
		if err != nil {
			return input.TeamMember{}, err
		}
		pair = newMember
	}

	return pair, nil
}
