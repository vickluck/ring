package main

import (
	"github.com/tnngo/nog"
	"github.com/tnngo/ring/balance"
)

func main() {
	b := balance.New()
	b.SetMode("111", balance.MODE_RR)
	b.Push("111", "11111")
	b.Push("111", "11111111111")

	b.Push("222", "22222222")

	//b.SetMode("222", balance.MODE_RR)
	//b.Push("22222")

	nog.SL.Debug(b.Pull("111"))
	nog.SL.Debug(b.Pull("111"))
	nog.SL.Debug(b.Pull("111"))
	nog.SL.Debug(b.Pull("222"))
	nog.SL.Debugf("111 length: %d", b.Len("111"))
}
