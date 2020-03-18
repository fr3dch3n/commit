# commit

> Easily build up a commit-message that conforms your team-conventions.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Table of Contents
  * [Table of Contents](#table-of-contents)
  * [Example](#example)
     * [Special usecases](#special-usecases)
        * [Blank commit-message](#blank-commit-message)
        * [Use defaults](#use-defaults)
  * [Commands and Flags](#commands-and-flags)
  * [Features](#features)
  * [Installation](#installation)
  * [Running the tests](#running-the-tests)
  * [Built With](#built-with)
  * [License](#license)

## Example

After a successful installation usage of `commit` looks like this:
```bash
$ commit
Current pairing partner (separate by [,| ])
» pair
Creating team-member with abbreviation pair
Enter username
» pair1234
Enter mail
» pair@mycompany.com
Current story
» #999
Summary of your commit
» fix build-step
Why did you choose to do that?
» Building the package failed because ...
»
»
```

This results in a commit-message like this: 
```
[#999] pair|fma Fix build-step

Building the package failed because ...

Co-authored-by: pair1234 <pair@mycompany.com>
```

### Special usecases

Some useful oneliners.

#### Blank commit-message

Write a blank commit-message but keep co-authored-by.

Command: `commit -b -m "Some commit-message"`

Commit-Message: 
```
Some commit-message
```
----
#### Use defaults

Use pair and story from state without review.

Command: `commit -y -m "yes we did it"` 

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
  -b, --blank            blank
  -y, --god-mode         git add all and use defaults from state
  -h, --help             help for commit
  -m, --message string   provide the commit-message
  -a, --no-git-add       do not run git add -p beforehand
  -p, --skip-pair        skip pair integration
  -v, --verbose          verbose output

Use "commit [command] --help" for more information about a command.
```

## Features
`commit...`
* uses a state file to save last pair and last story you commited to
* fixes some things in the summary-message (e.g. start with capital letter)
* adds new team-members on the fly
* adds needed configuration in an initial setup
* provides a god-mode (-y) to git-add-all and use pair/story from state without asking
* add multiple pairing partners

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

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
