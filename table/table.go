package table

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"text/tabwriter"
)

// Table is the 2D data structure with header.
type Table struct {
	header  []string
	Records [][]string
}

// CsvToTable makes table data from csv.
func CsvToTable(i io.Reader) (table *Table, err error) {
	r := csv.NewReader(i)
	t := new(Table)
	r.LazyQuotes = true
	header, err := r.Read()
	if err != nil {
		return nil, err
	}
	t.header = header
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		t.Records = append(t.Records, rec)
	}
	if len(t.header) != len(t.Records[0]) {
		fmt.Println("Unmutched Shape.")
		os.Exit(1)
	}
	return t, nil
}

//Calculation methods.

// Sum returns sum of columns.
func (t *Table) Sum(c string) int {
	sum := 0
	idx := 0
	for i, ele := range t.header {
		if c == ele {
			idx = i
		}
	}
	for _, slice := range t.Records {
		// If an error is occured, Atoi returns 0.
		// So no error check here.
		val, _ := strconv.Atoi(slice[idx])
		sum += val
	}
	return sum
}

// Information methods.

// Header method returns header of the table.
func (t *Table) Header() []string {
	return t.header
}

// SetHeader sets the header of the table.
func (t *Table) SetHeader(header []string) {
	t.header = header
}

// LenOfCols returns number of the columns of the table.
func (t *Table) LenOfCols() int {
	return len(t.header)
}

// LenOfRecs returns number of the records of the table.
func (t *Table) LenOfRecs() int {
	return len(t.Records)
}

// Info returns information of the table.
func (t *Table) Info() {
	f := os.Stdout
	fmt.Fprintf(f, "Cols: %v\nLenOfCols: %v\tLenOfRecs: %v\t\n",
		t.Header(), t.LenOfCols(), t.LenOfRecs())
}

//Display displays alighned table data.
func (t *Table) Display() {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	for _, header := range t.header {
		fmt.Fprintf(w, "%v\t", header)
	}
	fmt.Fprint(w, "\n")
	for _, slice := range t.Records {
		for _, v := range slice {
			fmt.Fprintf(w, "%v\t", v)
		}
		fmt.Fprint(w, "\n")
	}
	w.Flush()
}
