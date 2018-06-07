[![Build Status](https://travis-ci.com/Sh4d1/go-tvtime.svg?branch=master)](https://travis-ci.com/Sh4d1/go-tvtime)
[![Release](https://img.shields.io/github/release/Sh4d1/go-tvtime.svg?style=flat-square)](https://github.com/Sh4d1/go-tvtime/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE.md)
[![Coverage Status](https://img.shields.io/codecov/c/github/Sh4d1/go-tvtime/master.svg?style=flat-square)](https://codecov.io/gh/Sh4d1/go-tvtime)
[![Go Report Card](https://goreportcard.com/badge/github.com/Sh4d1/go-tvtime?style=flat-square)](https://goreportcard.com/report/github.com/Sh4d1/go-tvtime)
[![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser)

# Go TvTime

Go TvTime is a CLI application to manage TV Time account and episodes.

## Install

Install latest version using Golang (recommended)
```
$ go get github.com/Sh4d1/go-tvtime/cmd/tvtime
```

Or build from source
```
$ go get -u github.com/Sh4d1/go-tvtime
$ cd $GOPATH/src/github.com/Sh4d1/go-tvtime
$ make build
```

## Usage

```
$ tvtime -h
NAME:
   tvtime - tvtime

USAGE:
    tvtime [global options] command [command options] [arguments...]

VERSION:
    master

AUTHOR:
    Patrik Cyvoct (patrik@ptrk.io)

COMMANDS:
    login      Login to TV Time
    user       Display current user
    upcoming   Display upcoming episodes
    watchlist  Display watchlist
    help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
    --help, -h     show help
    --version, -v  print the version
```

## Contributions

This project adheres to the Contributor Covenant [code of conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.
We appreciate your contribution. Please refer to our [contributing guidelines](CONTRIBUTING.md).

## License

MIT

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FSh4d1%2Fgo-tvtime.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2FSh4d1%2Fgo-tvtime?ref=badge_large)
