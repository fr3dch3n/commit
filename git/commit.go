package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)


// ReviewSummary fixes the commit-summary by some simple rules.
func ReviewSummary(summary string) string {
	cleanedOfWhitespace := strings.TrimSpace(summary)
	return cleanedOfWhitespace
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

// Make an commit with no content
func EmptyCommit(commitMsg string) {
	out, err := exec.Command("git", "commit", "--allow-empty", "-m", commitMsg).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	output := string(out[:])
	fmt.Print(output)
}
