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
