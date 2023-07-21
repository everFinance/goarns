# Go ArNS

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/wujunze/goarns?style=flat-square)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/wujunze/goarns)](https://github.com/wujunze/goarns)
[![Go Report Card](https://goreportcard.com/badge/github.com/wujunze/goarns)](https://goreportcard.com/report/github.com/wujunze/goarns)
[![Unit-Tests](https://github.com/wujunze/goarns/workflows/Unit-Tests/badge.svg)](https://github.com/wujunze/goarns/actions)
[![Coverage Status](https://coveralls.io/repos/github/wujunze/goarns/badge.svg?branch=main)](https://coveralls.io/github/wujunze/goarns?branch=main)
[![Go Reference](https://pkg.go.dev/badge/github.com/wujunze/goarns.svg)](https://pkg.go.dev/github.com/wujunze/goarns)

> **[中文说明](README.zh-CN.md)**

## What is ArNS
ArNS: https://ar.io/arns/

### Installation

```shell
go get github.com/everFinance/goarns
```

## How to Use
```go
package example

import (
	"github.com/everFinance/goarns"
	"time"
)

func QueryArNS() {

	dreUrl := "https://dre-3.warp.cc"
	arNSAddress := "bLAgYxAdX2Ry-nt6aH2ixgvJXbpsEYm28NgJgyqfs-U"
	timeout := 10 * time.Second

	domain := "arseeding"

	a := goarns.NewArNS(dreUrl, arNSAddress, timeout)

	txId, err := a.QueryLatestRecord(domain)

	if err != nil {
		print(err)
	}
	print(txId)

}
```


