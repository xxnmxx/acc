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

	i0 := acc.Industory{92, [5]int{395, 369, 332, 319, 308}, 5.5, 45.0, 232.0}
	i1 := acc.Industory{94, [5]int{444, 374, 289, 279, 270}, 5.2, 36.0, 175.0}
	corp := acc.Prof{
		s, [2]acc.Div{d0, d1}, [3]acc.Rev{r0, r1, r2}, [2]acc.Equity{e0, e1},
		[2]acc.Industory{i0, i1}, 1.0,
	}
	fmt.Println(corp)
	fmt.Println(corp.ICalc())
	fmt.Println(corp.Eval())
}
