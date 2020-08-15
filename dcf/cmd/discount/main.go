package main

import (
	"fmt"

	"github.com/xxnmxx/acc/dcf"
)

func main() {
	tr := 0.3
	d := 500.0
	e := 1000.0
	ulb := dcf.UnleveredBeta(0.5,250,100,tr)
	lb := dcf.LeveredBeta(ulb,d,e,tr)
	ke := dcf.Ke(0.05,lb,0)
	kd := dcf.Kd(0.05,tr)
	wacc := dcf.SimpleWacc(kd,ke,d,e)
	cf := []float64{100, 100, 100, 200, 200}
	val := dcf.Discount(cf, wacc)
	fmt.Println(ulb,lb,ke,kd,e,d,wacc)
	fmt.Println(val)
}
