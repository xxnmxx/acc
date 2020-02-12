package main

import (
	"os"

	"github.com/xxnmxx/acc/table"
)

func main() {
	f := os.Stdin
	t, _ := table.CsvToTable(f)
	t.Display()
	//t.Info()
	//d := t.Sum("借方金額")
	//c := t.Sum("貸方金額")
	//fmt.Println(d, c, d-c)
	//g := t.ToGl()
	//fmt.Println(t.Shape())
	//fmt.Println(g.Records[1000:])
	//g.Display()
}
