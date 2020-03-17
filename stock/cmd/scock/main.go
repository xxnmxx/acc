package main

import (
	"fmt"

	"github.com/xxnmxx/acc/stock"
)

func main() {
	d := stock.DataImport("ipt")
	fmt.Println(d)
}
