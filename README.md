# commit

> Easily build up a commit-message that conforms your team-conventions.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Table of Contents
  * [Table of Contents](#table-of-contents)
  * [Usage](#usage)
     * [Conventional Commits](#conventional-commits)
     * [Story Commits](#story-commits)
     * [Special usecases](#special-usecases)
        * [Make an empty commit (a.k.a. git commit --allow-empty)](#make-an-empty-commit-aka-git-commit---allow-empty)
        * [Skip "git add -p"](#skip-git-add--p)
  * [Commands and Flags](#commands-and-flags)
  * [Features](#features)
  * [Installation](#installation)
     * [use github release](#use-github-release)
     * [build from scratch](#build-from-scratch)
  * [Running the tests](#running-the-tests)
  * [Built With](#built-with)
  * [License](#license)


## Usage

There are two types of supported commit-styles: `conventional-commits` and `story-commits`.

### Conventional Commits

```bash
$ commit
# runs `git add -p`
Current pairing partner (separate by [,| ])
» pair
Creating team-member with abbreviation pair (This part is skipped once the pair is known.)
Enter username
» pair1234
Enter mail
» pair@mycompany.com
Commit type
» fix
Current Scope
» actions
Summary of your commit
» fix build-step by providing ...
Why did you choose to do that?
» Building the package failed because ...
»
»
```

This results in a commit-message like this: 
```
fix(actions): fix build-step by providing ...

Building the package failed because ...

Co-authored-by: pair1234 <pair@mycompany.com>
```

### Story Commits

```bash
$ commit
# runs `git add -p`
Current pairing partner (separate by [,| ])
» pair
Creating team-member with abbreviation pair (This part is skipped once the pair is known.)
Enter username
» pair1234
Enter mail
» pair@mycompany.com
Current Story
» CI-999
Summary of your commit
» fix build-step
Why did you choose to do that?
» Building the package failed because ...
»
»
```

This results in a commit-message like this: 
```
[CI-999] fix build-step

Building the package failed because ...

Co-authored-by: pair1234 <pair@mycompany.com>
```

### Special usecases

Some useful flags.

#### Make an empty commit (a.k.a. git commit --allow-empty)

To make a commit without any changes (an empty commit), use `commit -e`.

#### Skip "git add -p"

To skip the `git add -p` that runs at start, simply use `commit -s`.

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
  -e, --empty-commit   make an empty commit
  -h, --help           help for commit
  -s, --skip-git-add   do not run git add -p beforehand
  -v, --verbose        verbose output

Use "commit [command] --help" for more information about a command.
```

## Features
`commit...`
* uses a state file to save last pair and last story you commited to
* adds new team-members on the fly
* adds the necessary configuration in an initial setup
* add multiple pairing partners

## Installation
### use github release
1. Select the binary from `Releases` according to your platform.
1. Rename it to `commit`.
1. Make it executable.
1. Put it in your path.
### build from scratch
To use `commit` you simply have to build it and put the binary on your path:

```bash
git clone https://github.com/fr3dch3n/commit
cd commit
make debug-build
ln -s $(pwd)/commit $HOME/bin/commit  # e.g.
```

`commit` will ask you for all configuration it needs.

## Running the tests

To execute all tests, simply run: `make test`.

## Built With

* [logrus](github.com/sirupsen/logrus) - The golang logging framework
* [cobra](github.com/spf13/cobra) - The cmd-line library
* [readline](https://github.com/chzyer/readline) - Lib to read from command-line
* [testify](github.com/stretchr/testify) - Make testing a blessing

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
