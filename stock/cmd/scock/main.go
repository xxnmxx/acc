package main

import (
	"fmt"

	"github.com/xxnmxx/acc/stock"
)

func main() {
	//d1 := new(stock.Data)
	//d1.InputSheet("test")
	d := stock.DataImport("ipt2")
	fmt.Println(d)
	b0, b1 := d.Div()
	fmt.Println(b0, b1)
	c0, c1, c2 := d.Income()
	fmt.Println(c0, c1, c2)
	d0, d1 := d.Equity()
	fmt.Println(d0, d1)
	r1 := d.Ratio(d.Index1)
	r2 := d.Ratio(d.Index2)
	fmt.Println(r1, r2)
	vp1 := d.ValueParJPY50(d.Index1)
	vp2 := d.ValueParJPY50(d.Index2)
	fmt.Println(vp1, vp2)
	fmt.Println(d.ValueParStock())

}
