package git

import (
	"fmt"
	"github.com/fr3dch3n/commit/input"
	"os"
	"os/exec"
	"strings"
)

func BuildCommitMsg(story string, pair input.TeamMember, summary string, explanation string, short string, skipShort bool) string {
	var output string

	if story != "" {
		output += fmt.Sprintf("[%s] ", story)
	}

	if !skipShort {
		if (input.TeamMember{}) == pair || pair.Short == "none" {
			output += fmt.Sprintf("%s ", short)
		} else {
			output += fmt.Sprintf("%s|%s ", short, pair.Short)
		}
	}

	output += fmt.Sprintf("%s\n", summary)

	if strings.TrimSpace(explanation) != "" {
		output += fmt.Sprintf("\n%s\n", explanation)
	}

	if (input.TeamMember{}) != pair && pair.Short != "none" {
		coAuthoredBy := fmt.Sprintf("Co-authored-by: %s <%s>\n", pair.GithubUserName, pair.Email)
		output += fmt.Sprintf("\n\n%s", coAuthoredBy)
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
	output := string(out[:])
	fmt.Print(output)
}

func AddP() {
	cmd := exec.Command("git", "add", "-p")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}

func AreThereChanges() bool {
	out, err := exec.Command("git", "status", "--porcelain").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	output := string(out[:])
	return output != ""
}
