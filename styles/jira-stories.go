package styles

import (
	"fmt"
	"github.com/fr3dch3n/commit/input"
	"github.com/fr3dch3n/commit/utils"
	log "github.com/sirupsen/logrus"
	"strings"
)

type JIRAStoryInformation struct {
	Story string
}

// BuildStoryStyleCommitMsg returns a string which is the whole commit message.
// Parameters are all previously asked for information like pair, Scope, summary and explanation.
func BuildStoryStyleCommitMsg(cci JIRAStoryInformation, pair []input.TeamMember, summary string, explanation string) string {
	log.Debugf("pair: %v", pair)
	var output string

	if cci.Story != "" {
		output += fmt.Sprintf("[%s]", cci.Story)
	}

	output += fmt.Sprintf(" %s\n", summary)

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


func GatherStoryStyleInformation(state input.State) JIRAStoryInformation {
	scope, err := input.GetWithDefault("Current Story", state.CurrentScope)
	utils.Check(err)

	return JIRAStoryInformation{
		Story: scope,
	}
}
