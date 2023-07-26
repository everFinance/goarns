# Go ArNS

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/everFinance/goarns?style=flat-square)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/everFinance/goarns)](https://github.com/everFinance/goarns)
[![Go Report Card](https://goreportcard.com/badge/github.com/everFinance/goarns)](https://goreportcard.com/report/github.com/everFinance/goarns)
[![Unit-Tests](https://github.com/everFinance/goarns/workflows/Unit-Tests/badge.svg)](https://github.com/everFinance/goarns/actions)
[![Coverage Status](https://coveralls.io/repos/github/everFinance/goarns/badge.svg?branch=main)](https://coveralls.io/github/everFinance/goarns?branch=main)
[![Go Reference](https://pkg.go.dev/badge/github.com/everFinance/goarns.svg)](https://pkg.go.dev/github.com/everFinance/goarns)


## 什么是 ArNS  
ArNS: https://ar.io/arns/

### 安装

```shell
go get github.com/everFinance/goarns
```

## 如何使用
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



