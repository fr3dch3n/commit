package git

import (
	"fmt"
	"github.com/fr3dch3n/commit/input"
	"os/exec"
	"strings"
)

func BuildCommitMsg(story string, pair input.TeamMember, summary string, explanation string, short string) string {
	var output string

	output += fmt.Sprintf("[%s] ", story)

	if (input.TeamMember{}) == pair {
		output += fmt.Sprintf("%s ", short)
	} else {
		output += fmt.Sprintf("%s|%s ", short, pair.Short)
	}

	output += fmt.Sprintf("%s", summary)

	if strings.TrimSpace(explanation) != "" {
		output += fmt.Sprintf("\n\n%s", explanation)
	}

	if (input.TeamMember{}) != pair {
		coAuthoredBy := fmt.Sprintf("Co-authored-by: %s <%s>\n", pair.GithubUserName, pair.Email)
		output += fmt.Sprintf("\n\n\n%s", coAuthoredBy)
	}

	return output
}

func ReviewSummary(summary string) string {
	cleanedOfWhitespace := strings.TrimSpace(summary)
	firstChar := string(cleanedOfWhitespace[0])
	return strings.ToUpper(firstChar) + strings.TrimPrefix(cleanedOfWhitespace, firstChar)
}

func Commit(commitMsg string) {
	out, err := exec.Command("git", "commit", "-m", commitMsg).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Command Successfully Executed")
	output := string(out[:])
	fmt.Println(output)
}
