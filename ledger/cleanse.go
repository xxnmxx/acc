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
	col := []int{0, 1, 3}
	revs := csv.cut(col)
	ced := csv.cutOne(1)
	//fmt.Println(csv.getColumns())
	//fmt.Printf("%T\t%v", r, r)
	fmt.Println(ced.Records[:3])
	fmt.Println(revs.Records[:3])

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
	revs := new(Csv)
	for _, slice := range c.Records {
		tmp := make([]string, 0)
		for j, w := range col {
			tmp = append(slice[:col[w]-j], slice[col[w]-j+1:]...)
		}
		revs.Records = append(revs.Records, tmp)
	}
	return revs
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
