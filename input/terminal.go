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
