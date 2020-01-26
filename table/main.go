package main

import (
	"fmt"
	"os"

	"github.com/xxnmxx/acc"
)

func main() {
	f := os.Stdin
	t, _ := acc.CsvToTable(f)
	//t.Display()
	t.Info()
	d := t.Sum("借方金額")
	c := t.Sum("貸方金額")
	fmt.Println(d, c, d-c)
}
