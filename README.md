[![ci](https://github.com/woodpecker-kit/woodpecker-transfer-data/actions/workflows/ci.yml/badge.svg)](https://github.com/woodpecker-kit/woodpecker-transfer-data/actions/workflows/ci.yml)

[![go mod version](https://img.shields.io/github/go-mod/go-version/woodpecker-kit/woodpecker-transfer-data?label=go.mod)](https://github.com/woodpecker-kit/woodpecker-transfer-data)
[![GoDoc](https://godoc.org/github.com/woodpecker-kit/woodpecker-transfer-data?status.png)](https://godoc.org/github.com/woodpecker-kit/woodpecker-transfer-data)
[![goreportcard](https://goreportcard.com/badge/github.com/woodpecker-kit/woodpecker-transfer-data)](https://goreportcard.com/report/github.com/woodpecker-kit/woodpecker-transfer-data)

[![GitHub license](https://img.shields.io/github/license/woodpecker-kit/woodpecker-transfer-data)](https://github.com/woodpecker-kit/woodpecker-transfer-data)
[![codecov](https://codecov.io/gh/woodpecker-kit/woodpecker-transfer-data/branch/main/graph/badge.svg)](https://codecov.io/gh/woodpecker-kit/woodpecker-transfer-data)
[![GitHub latest SemVer tag)](https://img.shields.io/github/v/tag/woodpecker-kit/woodpecker-transfer-data)](https://github.com/woodpecker-kit/woodpecker-transfer-data/tags)
[![GitHub release)](https://img.shields.io/github/v/release/woodpecker-kit/woodpecker-transfer-data)](https://github.com/woodpecker-kit/woodpecker-transfer-data/releases)

## for what

- this project used to woodpecker-ci data transfer define

## Contributing

[![Contributor Covenant](https://img.shields.io/badge/contributor%20covenant-v1.4-ff69b4.svg)](.github/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md)
[![GitHub contributors](https://img.shields.io/github/contributors/woodpecker-kit/woodpecker-transfer-data)](https://github.com/woodpecker-kit/woodpecker-transfer-data/graphs/contributors)

We welcome community contributions to this project.

Please read [Contributor Guide](.github/CONTRIBUTING_DOC/CONTRIBUTING.md) for more information on how to get started.

请阅读有关 [贡献者指南](.github/CONTRIBUTING_DOC/zh-CN/CONTRIBUTING.md) 以获取更多如何入门的信息

## depends

in go mod project

```bash
# warning use private git host must set
# global set for once
# add private git host like github.com to evn GOPRIVATE
$ go env -w GOPRIVATE='github.com'
# use ssh proxy
# set ssh-key to use ssh as http
$ git config --global url."git@github.com:".insteadOf "https://github.com/"
# or use PRIVATE-TOKEN
# set PRIVATE-TOKEN as gitlab or gitea
$ git config --global http.extraheader "PRIVATE-TOKEN: {PRIVATE-TOKEN}"
# set this rep to download ssh as https use PRIVATE-TOKEN
$ git config --global url."ssh://github.com/".insteadOf "https://github.com/"

# before above global settings
# test version info
$ git ls-remote -q https://github.com/woodpecker-kit/woodpecker-transfer-data.git

# test depends see full version
$ go list -mod readonly -v -m -versions github.com/woodpecker-kit/woodpecker-transfer-data
# or use last version add go.mod by script
$ echo "go mod edit -require=$(go list -mod=readonly -m -versions github.com/woodpecker-kit/woodpecker-transfer-data | awk '{print $1 "@" $NF}')"
$ echo "go mod vendor"
```

## Features

- [ ] more perfect test case coverage
- [ ] more perfect benchmark case

## usage

- use this template, replace list below
    - `github.com/woodpecker-kit/woodpecker-transfer-data` to your package name
    - `woodpecker-kit` to your owner name
    - `woodpecker-transfer-data` to your project name

# dev

```bash
# It needs to be executed after the first use or update of dependencies.
$ make init dep
```

- test code

```bash
$ make test testBenchmark
```

add main.go file and run

```bash
# run at env dev use cmd/main.go
$ make dev
```

- ci to fast check

```bash
# check style at local
$ make style

# run ci at local
$ make ci
```

## docker

```bash
# then test build as test/Dockerfile
$ make dockerTestRestartLatest
# clean test build
$ make dockerTestPruneLatest

# more info see
$ make helpDocker
```
