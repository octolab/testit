> # 👨‍🍳 testit
>
> Extended `go test` for better experience.

[![Build][build.icon]][build.page]
[![Documentation][docs.icon]][docs.page]
[![Quality][quality.icon]][quality.page]
[![Template][template.icon]][template.page]
[![Coverage][coverage.icon]][coverage.page]
[![Mirror][mirror.icon]][mirror.page]

## 💡 Idea

```bash
$ testit go compile
# -> go test -run ^Fake$$ ./... | { grep -v 'no test files' | grep -v 'no tests to run' || true }
$ testit go --colored --panicparse [testflag]
# -> gotest [testflag] | panicparse // with color preserve
```

A full description of the idea is available [here][design.page].

## 🏆 Motivation

...

## 🤼‍♂️ How to

...

## 🧩 Installation

### Homebrew

```bash
$ brew install :owner/tap/:binary
```

### Binary

```bash
$ curl -sSfL https://raw.githubusercontent.com/:owner/:repository/master/bin/install | sh
# or
$ wget -qO-  https://raw.githubusercontent.com/:owner/:repository/master/bin/install | sh
```

> Don't forget about [security](https://www.idontplaydarts.com/2016/04/detecting-curl-pipe-bash-server-side/).

### Source

```bash
# use standard go tools
$ go get github.com/:owner/:repository@:version
# or use egg tool
$ egg tools add github.com/:owner/:repository@:version
```

> [egg][] is an `extended go get`.

### Shell completions

```bash
$ :binary completion bash|fish|powershell|zsh > /path/to/completions/...
# or
$ source <(:binary completion bash|fish|powershell|zsh)
```

<p align="right">made with ❤️ for everyone</p>

[awesome.icon]:     https://awesome.re/mentioned-badge.svg
[build.page]:       https://travis-ci.com/octolab/testit
[build.icon]:       https://travis-ci.com/octolab/testit.svg?branch=master
[coverage.page]:    https://codeclimate.com/github/octolab/testit/test_coverage
[coverage.icon]:    https://api.codeclimate.com/v1/badges/c570179a9335c747e77c/test_coverage
[design.page]:      https://www.notion.so/33715348cc114ea79dd350a25d16e0b0?r=0b753cbf767346f5a6fd51194829a2f3
[docs.page]:        https://pkg.go.dev/go.octolab.org/toolset/testit
[docs.icon]:        https://img.shields.io/badge/docs-pkg.go.dev-blue
[mirror.page]:      https://bitbucket.org/kamilsk/testit
[mirror.icon]:      https://img.shields.io/badge/mirror-bitbucket-blue
[promo.page]:       https://github.com/octolab/testit
[quality.page]:     https://goreportcard.com/report/go.octolab.org/toolset/testit
[quality.icon]:     https://goreportcard.com/badge/go.octolab.org/toolset/testit
[template.page]:    https://github.com/octomation/go-tool
[template.icon]:    https://img.shields.io/badge/template-go--tool-blue

[_]:                https://img.shields.io/sourcegraph/rrc/github.com/octolab/testit
[egg]:              https://github.com/kamilsk/egg
