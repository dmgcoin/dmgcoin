
Dmgcoin
====

[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://choosealicense.com/licenses/isc/)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/dmgcoin/dmgcoin)

Dmgcoin is the reference full node implementation written in Go (golang).

## What is dmg

Dmgcoin is an attempt at a proof-of-work cryptocurrency with instant confirmations and sub-second block times. It is based on [the PHANTOM protocol](https://eprint.iacr.org/2018/104.pdf), a generalization of Nakamoto consensus.

## Requirements

Go 1.18 or later.

## Installation

#### Build from Source

- Install Go according to the installation instructions here:
  http://golang.org/doc/install

- Ensure Go was installed properly and is a supported version:

```bash
$ go version
```

- Run the following commands to obtain and install dmgcoin including all dependencies:

```bash
$ git clone https://github.com/dmgcoin/dmgcoin
$ cd dmgcoin
$ go install . ./cmd/...
```

- Dmgcoin (and utilities) should now be installed in `$(go env GOPATH)/bin`. If you did
  not already add the bin directory to your system path during Go installation,
  you are encouraged to do so now.


## Getting Started

Dmgcoin has several configuration options available to tweak how it runs, but all
of the basic operations work with zero configuration.

```bash
$ dmgcoin
```

## Discord
Join our discord server using the following link: https://discord.gg/xh7d5tVy

## Issue Tracker

The [integrated github issue tracker](https://github.com/dmgcoin/dmgcoin/issues)
is used for this project.

Issue priorities may be seen at https://github.com/orgs/dmgcoin/projects/4

## Documentation

The [documentation](https://github.com/dmgcoin/docs) is a work-in-progress

## License

Dmgcoin is licensed under the copyfree [ISC License](https://choosealicense.com/licenses/isc/).
