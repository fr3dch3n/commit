package input

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func GetInputOrElse(ioreader io.Reader, msg string, current string) (string, error) {
	var input string
	if current != "" {
		reader := bufio.NewReader(ioreader)
		fmt.Print(msg + " (" + current + "): ")
		input, _ = reader.ReadString('\n')
	} else {
		reader := bufio.NewReader(ioreader)
		fmt.Print(msg + ":")
		input, _ = reader.ReadString('\n')
	}
	cleanInput := strings.TrimSpace(input)
	if cleanInput != "" {
		return cleanInput, nil
	} else {
		return current, nil
	}
}

func GetInput(ioreader io.Reader, msg string) (string, error) {
	var err error
	var input string
	reader := bufio.NewReader(ioreader)
	fmt.Print(msg + ": ")
	input, err = reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func GetMultiLineInput(ioreader io.Reader, msg string) (string, error) {
	scn := bufio.NewScanner(ioreader)
	fmt.Print(msg)
	var lines []string
	var nrOfEnters = 0
	for scn.Scan() {
		line := scn.Text()
		if len(line) == 0 {
			nrOfEnters += 1
			if nrOfEnters == 2 {
				break
			}
		} else {
			nrOfEnters = 0
			lines = append(lines, line)
		}
	}

	return strings.Join(lines, "\n"), nil
}

func GetNewTeamMemberFromInput(ioreader io.Reader) (TeamMember, error) {
	fmt.Println("Specifying new team member")
	short, err := GetInput(ioreader, "Enter short")
	if err != nil {
		return TeamMember{}, nil
	}
	username, err := GetInput(ioreader, "Enter username")
	if err != nil {
		return TeamMember{}, nil
	}
	mail, err := GetInput(ioreader, "Enter mail")
	if err != nil {
		return TeamMember{}, nil
	}

	return TeamMember{
		GithubUserName: username,
		Email:          mail,
		Short:          short,
	}, nil
}
