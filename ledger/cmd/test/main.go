package main

import (
	"fmt"

	"github.com/xxnmxx/acc/ledger"
)

func main() {
	m := ledger.CreateAccMaster()
	m.AddAccMaster("cash", 1)
	m.AddAccMaster("land", 2)
	m.AddAccMaster("ap", 3)
	m.AddAccMaster("land", 2)
	fmt.Printf("%+v", m)
	m.WriteAccMaster("AccMaster")
	//m2 := ledger.LoadAccMaster("AccMaster")
	//fmt.Printf("%+v", m2)
}
