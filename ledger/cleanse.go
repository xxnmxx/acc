package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	home, _ := os.UserHomeDir()
	var p string = filepath.Join(home, "Downloads", "acc.txt")
	csv := readCsv(p)
	col := []int{0}
	r := csv.cut(col)
	fmt.Println(*r)

}

type Csv struct {
	Records [][]string
}

func readCsv(n string) *Csv {
	c := new(Csv)
	f, err := os.Open(n)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	r := csv.NewReader(f)
	r.LazyQuotes = true
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		c.Records = append(c.Records, record)
	}
	return c
}

func (c *Csv) getColumns() []string {
	columns := make([]string, 0)
	for i, column := range c.Records[0] {
		ele := fmt.Sprint(i, ": ", column)
		columns = append(columns, ele)
	}
	return columns
}

func (c *Csv) cut(col []int) *Csv {
	revCsv := new(Csv)
	for pos, slice := range c.Records {
		for i, v := range slice {
			for _, w := range col {
				if i != w {
					revCsv.Records[pos] = append(revCsv.Records[pos], v)
				}
			}
		}
	}
	return revCsv
}
