package main

import (
	"fmt"
	"github.com/xxnmxx/acc"
)

func main() {
	s := acc.Stock{100, 0}

	d0 := acc.Div{100, 0}
	d1 := acc.Div{150, 0}

	r0 := acc.Rev{1000, 0, 0, 0, 0}
	r1 := acc.Rev{1000, 0, 0, 0, 0}
	r2 := acc.Rev{1000, 0, 0, 0, 0}

	e0 := acc.Equity{200, 1000}
	e1 := acc.Equity{200, 1000}

	corp := acc.Prof{
		s,[2]acc.Div{d0,d1},[3]acc.Rev{r0,r1,r2},[2]acc.Equity{e0,e1},
	}
	fmt.Println(corp)

	ds := []int{}
	for _, d := range corp.Div {
		ds = append(ds, d.VCalc())
	}
	fmt.Println(ds)
}
