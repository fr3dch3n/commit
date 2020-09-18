package input

import (
	"fmt"
	"io"
	"strings"

	"github.com/chzyer/readline"
	"github.com/fr3dch3n/commit/utils"
	log "github.com/sirupsen/logrus"
)

var rl *readline.Instance

func init() {
	var err error
	c := &readline.Config{
		Prompt:          "\033[31m»\033[0m ",
		InterruptPrompt: "^C",
	}

	rl, err = readline.NewEx(c)
	if err != nil {
		panic(err)
	}

}

// TODO shutdown close

// GetWithDefault asks for user-input and provides a default.
func GetWithDefault(msg, defa string) (string, error) {
	fmt.Println(msg)
	line, err := rl.ReadlineWithDefault(defa)
	log.Debug("Read: " + line)
	if err == io.EOF {
		return strings.TrimSpace(line), nil
	} else if err != nil && err.Error() == "Interrupt" {
		utils.Abort()
	}
	return strings.TrimSpace(line), nil
}

// Get asks for user-input.
func Get(msg string) (string, error) {
	fmt.Println(msg)
	line, err := rl.Readline()
	log.Debug("Read: " + line)
	if err == io.EOF {
		return strings.TrimSpace(line), nil
	} else if err != nil && err.Error() == "Interrupt" {
		utils.Abort()
	}
	return strings.TrimSpace(line), nil
}

// GetNonEmpty asks for user-input until the input is not empty.
func GetNonEmpty(msg string) (string, error) {
	fmt.Println(msg)
	line, err := rl.Readline()
	log.Debug("Read: " + line)
	if err == io.EOF || (err == nil && line == "") {
		return GetNonEmpty(msg)
	} else if err != nil && err.Error() == "Interrupt" {
		utils.Abort()
	}
	return strings.TrimSpace(line), nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func GetCommitType(printHelp bool) (string, error) {
	help := `Allowed commit types are:
fix:		this commit patches a bug
feat:		this commit introduces a new feature
build:		this commit affects the build system
chore:		this commit changes tools or other things that do not go into production at all (grunt tasks)
ci:		this commit changes the CI configuration files and scripts
docs:		this commit changes documentation
style:		this commit does not affect meaning of code (white-space, formatting, missing semi-colons, etc)
refactor:	this commit neither fixes a bug nor adds a feature
revert:		this commit reverts a preceded commit
perf:		this commit improves performance
test:		this commit adds missing tests or corrects existing tests
Note: For a breaking change append a ! after the type/scope.`

	if printHelp {
		fmt.Println(help)
	}
	fmt.Println("Commit type")
	allowedValues := []string{"fix", "feat", "build", "chore", "ci", "docs", "style", "refactor", "perf", "test", "fix!", "feat!", "build!", "chore!", "ci!", "docs!", "style!", "perf!", "test!"}

	line, err := rl.Readline()
	log.Debug("Read: " + line)
	if err != nil && err.Error() == "Interrupt" {
		utils.Abort()
	} else if err == io.EOF || (err == nil && line == "") || !stringInSlice(line, allowedValues) {
		return GetCommitType(true)
	}
	return strings.TrimSpace(line), nil
}

// GetMultiLineInput lets a user input many lines until two blank lines follow one another.
func GetMultiLineInput(msg string) (string, error) {
	var lines []string
	var emptyLineCounter int = 0
	fmt.Println(msg)
	for {
		line, err := rl.Readline()
		if line == "" {
			if emptyLineCounter == 1 {
				break
			} else {
				emptyLineCounter++
			}
		} else {
			emptyLineCounter = 0
		}
		if err != nil {
			panic(err)
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n"), nil
}
