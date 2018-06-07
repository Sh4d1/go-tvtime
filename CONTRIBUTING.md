# Contributing

By participating to this project, you agree to abide our [code of
conduct](/CODE_OF_CONDUCT.md).

## Setup your machine

`go-tvtime` is written in [Go](https://golang.org/).

Prerequisites are:

* Build:
  * `make`
  * [Go 1.10+](http://golang.org/doc/install)

Clone `go-tvtime` from source into `$GOPATH`:

```sh
$ go get github.com/Sh4d1/go-tvtime
$ cd $GOPATH/src/github.com/Sh4d1/go-tvtime
```

Install the build and lint dependencies:

``` sh
$ make setup
```

## Test your change

You can create a branch for your changes and try to build from the source as you go:

``` sh
$ go build
```

When you are satisfied with the changes, we suggest you run:

``` sh
$ make ci
```

Which runs all the linters and tests.

## Submit a pull request

Push your branch to your `go-tvtime` fork and open a pull request against the
master branch.
