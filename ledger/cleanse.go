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
	//fmt.Println(csv)
	col := []int{0, 1, 2, 3}
	csv.cut(col)
	//ced := csv.cutOne(1)
	//fmt.Println(csv.getColumns())
	//fmt.Printf("%T\t%v", r, r)
	//fmt.Println(ced.Records[:3])
	//fmt.Println(revs.Records[:3])
	fmt.Println(csv.Records[:3])
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

func (c *Csv) cut(col []int) {
	//revs := new(Csv)
	for _, slice := range c.Records {
		//tmp := make([]string, len(c.Records[pos])-len(col))
		for j, w := range col {
			//tmp = append(slice[:w-j], slice[w-j+1:]...)
			if w < len(slice)-1 {
				slice = append(slice[:w-j], slice[w-j+1:]...)
				slice[len(slice)-1] = ""
				slice = slice[:len(slice)-1]
			}
		}
		//revs.Records[pos] = slice
	}
	//return revs
}

func (c *Csv) cutOne(dlcol int) *Csv {
	ced := new(Csv)
	for _, slice := range c.Records {
		tmp := make([]string, len(c.Records)-1)
		tmp = append(slice[:dlcol], slice[dlcol+1:]...)
		ced.Records = append(ced.Records, tmp)
	}
	return ced
}
