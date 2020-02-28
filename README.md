# commit

> Easily build up a commit-message that conforms your team-conventions.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Table of Contents
  * [Table of Contents](#table-of-contents)
  * [Example](#example)
     * [Special usecases](#special-usecases)
  * [Commands and Flags](#commands-and-flags)
  * [Features](#features)
  * [Installation](#installation)
  * [Running the tests](#running-the-tests)
  * [Built With](#built-with)
  * [Further Todos](#further-todos)
  * [License](#license)

## Example

After a successful installation usage of `commit` looks like this:
```bash
$ commit
Pairing with [abc]: def
Story [ABC-001]: ABC-002
Summary of your commit: we do work
Why did you choose to do that? We need to make sure that we 
keep up the good work.
```

This results in a commit-message like this: 
```
[ABC-002] def|me We do work

We need to make sure that we 
keep up the good work.

Co-authored-by: def-user <def@sample.com>
```

### Special usecases

Command: `commit -neps -m "Some commit-message"`

Commit-Message: 
```
Some commit-message
```
----

Command: `commit -ey -m "yes we did it"` 

Commit-Message: 
```
[last-story] pair|me Yes we did it

Co-authored-by: pair <pair@company.com>
```

## Commands and Flags
```bash
Easily build up a commit-message that conforms your team-conventions.

Usage:
  commit [flags]
  commit [command]

Available Commands:
  help        Help about any command
  version     Print the version number of Commit

Flags:
  -y, --add-all-with-defaults   git add all and use defaults from state
  -a, --git-add                 run git add -p beforehand
  -h, --help                    help for commit
  -m, --message string          provide the commit-message
  -n, --skip-abbreviations      skip listing abbreviations
  -e, --skip-explanation        skip long explanation
  -p, --skip-pair               skip pair integration
  -s, --skip-story              skip story integration
  -v, --verbose                 verbose output

Use "commit [command] --help" for more information about a command.
```

## Features
`commit...`
* uses a state file to save last pair and last story you commited to
* fixes some things in the summary-message (e.g. start with capital letter)
* adds new team-members on the fly
* adds needed configuration in an initial setup
* provides a god-mode (-y) to git-add-all and use pair/story from state without asking

## Installation

To use `commit` you simply have to build it and put the binary on your path:

```bash
git clone https://github.com/fr3dch3n/commit
cd commit
make build
ln -s $(pwd)/commit $HOME/bin/commit  # e.g.
```

`commit` will ask you for all configuration it needs.

## Running the tests

To execute all tests, simply run: `make test`.

## Built With

* [logrus](github.com/sirupsen/logrus) - The golang logging framework
* [cobra](github.com/spf13/cobra) - The cmd-line library
* [testify](github.com/stretchr/testify) - Make testing a blessing

## Further Todos
- [x] save "none" as pair
- [x] summary is mandatory
- [x] make `commit` a cmd-line-tool
- [x] add flag to run git add -p beforehand
- [x] add flag to skip story
- [x] add flag to skip pair
- [x] add flag to skip long explanation
- [x] check if there is anything to commit
- [x] separate state from config
- [x] use abbreviation
- [x] personal infos from team-config
- [x] -y for yes to all git diffs and take default parameter
- [x] initial config setup
- [x] -m Flag to provide a message from cmd-line
- [x] exit if there are no staged files after git-add-phase

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
