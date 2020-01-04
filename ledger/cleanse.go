package ledger

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

// Csv is the struct for table data.
type Csv struct {
	Records [][]string
}

//ReadCsv reads csv data from file.
func ReadCsv(n string) *Csv {
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

//GetColumns gets header of the table.
func (c *Csv) GetColumns() []string {
	columns := make([]string, 0)
	for i, column := range c.Records[0] {
		ele := fmt.Sprint(i, ": ", column)
		columns = append(columns, ele)
	}
	return columns
}

//Cut cuts the specified columns.
func (c *Csv) Cut(col []int) {
	for _, slice := range c.Records {
		for j, w := range col {
			if w < len(slice)-1 {
				slice = append(slice[:w-j], slice[w-j+1:]...)
				slice[len(slice)-1] = ""
				slice = slice[:len(slice)-1]
			}
		}
	}
}

//CutOne cuts the specified column.
func (c *Csv) CutOne(dlcol int) *Csv {
	ced := new(Csv)
	for _, slice := range c.Records {
		tmp := make([]string, len(c.Records)-1)
		tmp = append(slice[:dlcol], slice[dlcol+1:]...)
		ced.Records = append(ced.Records, tmp)
	}
	return ced
}
