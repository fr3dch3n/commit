package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fr3dch3n/commit/input"
	log "github.com/sirupsen/logrus"
)

// BuildCommitMsg returns a string which is the whole commit message.
// Parameters are all previously asked for information like pair, story, summary and explanation.
func BuildCommitMsg(story string, pair []input.TeamMember, summary string, explanation string, me input.TeamMember, blank bool) string {
	log.Debug("story: " + story)
	log.Debugf("pair: %v", pair)
	var output string

	if story != "" && !blank {
		output += fmt.Sprintf("[%s] ", story)
	}

	if !blank {
		if len(pair) == 0 {
			output += fmt.Sprintf("%s ", me.Abbreviation)
		} else {
			var pairstring = ""
			for _, p := range pair {
				pairstring += fmt.Sprintf("%s|", p.Abbreviation)
			}
			output += fmt.Sprintf("%s%s ", pairstring, me.Abbreviation)

		}
	}

	output += fmt.Sprintf("%s\n", summary)

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

// ReviewSummary fixes the commit-summary by some simple rules.
func ReviewSummary(summary string) string {
	cleanedOfWhitespace := strings.TrimSpace(summary)
	firstChar := string(cleanedOfWhitespace[0])
	return strings.ToUpper(firstChar) + strings.TrimPrefix(cleanedOfWhitespace, firstChar)
}

// Commit executes git-commit with the build message.
func Commit(commitMsg string) {
	out, err := exec.Command("git", "commit", "-m", commitMsg).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	output := string(out[:])
	fmt.Print(output)
}

// Add executes the git-add command with a mode.
// A mode could be: -p or simply a dot.
func Add(mode, extraArg string) {
	var cmd *exec.Cmd
	if extraArg == "" {
		cmd = exec.Command("git", "add", mode)
	} else {
		cmd = exec.Command("git", "add", mode, extraArg)
	}
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}

// AreThereChanges checks if there are local changes in the current dir.
func AreThereChanges() bool {
	out, err := exec.Command("git", "status", "--porcelain").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	output := string(out[:])
	return output != ""
}

// AnythingStaged checks if there are staged changes in the current dir.
func AnythingStaged() bool {
	out, err := exec.Command("git", "diff", "--name-only", "--cached").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	output := strings.TrimSpace(string(out[:]))
	return output != ""
}
