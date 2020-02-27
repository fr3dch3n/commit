# commit

> Easily build up a commit-message that conforms your team-conventions.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Table of Contents
* [Table of Contents](#table-of-contents)
* [Example](#example)
* [Features](#features)
* [Installation](#installation)
    * [Binary](#binary)
    * [Configuration files](#configuration-files)
* [Running the tests](#running-the-tests)
* [Built With](#built-with)
* [Further Todos](#further-todos)
* [License](#license)

## Example

After a successful installation usage of `commit` looks like this:
```bash
$ commit
Pairing with (abc): def
Story (ABC-001): ABC-002
Summary of your commit: we do work
Why did you choose to do that? We need to make sure that we 
keep up the good work.
```

This results in a commit-message like this: 
```
[ABC-002] me|def We do work

We need to make sure that we 
keep up the good work.

Co-authored-by: def-user <def@sample.com>
```

## Features
* save last story you commited to
* save last pair you worked with
* minimal fixed in summary-message (e.g. start with capital letter)
* add new team-members on the fly

## Installation

### Binary

To use `commit` you simply have to build it and put the binary on your path:

```bash
git clone https://github.com/fr3dch3n/commit
cd commit
make build
ln -s $(pwd)/commit $HOME/bin/commit  # e.g.
```

### Configuration files

You need a basic configuration file in `~/.commit-config`:
```json
{
    "username": "my_github_username",
    "story": "",
    "pair": "",
    "short": "my short name",
    "teamMembersConfigPath": "~/my-team/team-members.json"
}
```

And a team-members configuration file in the path specified above:
```json
[
    {
        "username": "another_username",
        "mail": "user1@company.com",
        "short": "user1"
    },
    {
        "username": "another_username2",
        "mail": "user2@company.com",
        "short": "user2"
    }
]
```

## Running the tests

To execute all tests, simply run: `make test`.

## Built With

* [logrus](github.com/sirupsen/logrus) - The golang logging framework

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
- [ ] use abbreviation
- [ ] personal infos from team-config
- [ ] -y for yes to all git diffs and take default parameter

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
