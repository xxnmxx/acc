package main

import (
	"fmt"

	"github.com/xxnmxx/acc/dcf"
)

func main() {
	e := dcf.NewEntity()
	cf := []float64{1000,2000,1500,2500,3000}
	fmt.Printf("prime:%+v\n\n",e)
	e.SetCf(cf)
	e.SetTaxRate(0.3)
	e.SetDebt(5000)
	e.SetRiskPremium(0.1)
	e.SetRiskFree(0.001)
	e.SetKd(0.01)
	e.SetCompEnt(0.5,500,1000,0.3)
	e.Calc()
	//e.Eval()
	//fmt.Printf("%+v\n",e)
	for i:=0;i<5;i++{
		e.Calc()
		fmt.Printf("%v. %+v\n\n",i,e)
	}
	e.Eval()
	fmt.Printf("eval: %+v\n",e)
}
