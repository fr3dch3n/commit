package git

import (
	"fmt"
	"github.com/fr3dch3n/commit/input"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"strings"
)

func BuildCommitMsg(story string, pair input.TeamMember, summary string, explanation string, me input.TeamMember, skipAbbreviation bool) string {
	log.Debug("story: " + story)
	log.Debug("pair: " + pair.Abbreviation)
	var output string

	if story != "" {
		output += fmt.Sprintf("[%s] ", story)
	}

	if !skipAbbreviation {
		if (input.TeamMember{}) == pair || pair.Abbreviation == "none" {
			output += fmt.Sprintf("%s ", me.Abbreviation)
		} else {
			output += fmt.Sprintf("%s|%s ", pair.Abbreviation, me.Abbreviation)
		}
	}

	output += fmt.Sprintf("%s\n", summary)

	if strings.TrimSpace(explanation) != "" {
		output += fmt.Sprintf("\n%s\n", explanation)
	}

	if (input.TeamMember{}) != pair && pair.Abbreviation != "none" {
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

func Add(mode string) {
	cmd := exec.Command("git", "add", mode)
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
