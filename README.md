# orders-tracker-cli

[![CI](https://github.com/cristovaoolegario/orders-tracker-cli/actions/workflows/main.yml/badge.svg)](https://github.com/cristovaoolegario/orders-tracker-cli/actions/workflows/main.yml)
[![codecov](https://codecov.io/gh/cristovaoolegario/orders-tracker-cli/branch/main/graph/badge.svg?token=o5n6lISvdW)](https://codecov.io/gh/cristovaoolegario/orders-tracker-cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/cristovaoolegario/orders-tracker-cli)](https://goreportcard.com/report/github.com/cristovaoolegario/orders-tracker-cli)

[![Go Reference](https://pkg.go.dev/badge/github.com/cristovaoolegario/orders-tracker-cli.svg)](https://pkg.go.dev/github.com/cristovaoolegario/orders-tracker-cli)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/cristovaoolegario/orders-tracker-cli)](https://github.com/cristovaoolegario/orders-tracker-cli/releases)
[![GitHub license](https://img.shields.io/github/license/cristovaoolegario/orders-tracker-cli)](https://github.com/cristovaoolegario/orders-tracker-cli/blob/main/LICENSE)

> This project still under development.
>
> This project was inspired by [track-correios](https://github.com/mauriciomutte/track-correios).

A CLI tool written in [Go](https://golang.org/) to track your orders.

It supports the following APIs:

- Correios API

## Installing

```shell
go get github.com/cristovaoolegario/orders-tracker-cli

go install github.com/cristovaoolegario/orders-tracker-cli
```

## Usage

```shell
orders-tracker-cli correios YOUR_ORDER_CODE
```

### Example

![How to track a correios order](https://github.com/cristovaoolegario/orders-tracker-cli/blob/main/static/usage_example.gif?raw=true)

## Contributing

If you want to contribute please fork the repository and make the changes as you'd like and submit a new PR.
