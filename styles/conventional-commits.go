package styles

import (
	"fmt"
	"github.com/fr3dch3n/commit/input"
	"github.com/fr3dch3n/commit/utils"
	log "github.com/sirupsen/logrus"
	"strings"
)


type ConventionalCommitsInformation struct {
	commitType string
	Scope      string
}

// BuildConventionalCommitMsg returns a string which is the whole commit message.
// Parameters are all previously asked for information like pair, Scope, summary and explanation.
func BuildConventionalCommitMsg(cci ConventionalCommitsInformation, pair []input.TeamMember, summary string, explanation string) string {
	log.Debug("Scope: " + cci.Scope)
	log.Debugf("pair: %v", pair)
	var output string

	if cci.commitType != "" {
		output += fmt.Sprintf("%s", cci.commitType)
	}
	if cci.Scope != "" {
		output += fmt.Sprintf("(%s)", cci.Scope)
	}

	output += fmt.Sprintf(": %s\n", summary)

	if strings.TrimSpace(explanation) != "" {
		output += fmt.Sprintf("\n%s\n", explanation)
	}

	if len(pair) != 0 {
		output += fmt.Sprintf("\n\n")
		for _, p := range pair {
			coAuthoredBy := fmt.Sprintf("Co-authored-by: %s <%s>\n", p.GithubUserName, p.Email)
			output += fmt.Sprintf("%s", coAuthoredBy)
		}
	}

	return output
}


func GatherConventionalCommitInformation(state input.State) ConventionalCommitsInformation {
	ctype, err := input.GetCommitType(false)
	utils.Check(err)

	scope, err := input.GetWithDefault("Current Scope", state.CurrentScope)
	utils.Check(err)

	return ConventionalCommitsInformation{
		commitType: ctype,
		Scope:      scope,
	}
}
