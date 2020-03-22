// ToDo Make other headers.
// Need to uniq check.

package ledger

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type accClass int

const (
	currentAsset accClass = iota + 1
	fixedAsset
	currentLiability
	fixedLiability
	equity
	sales
	cogs
	sga
	nopinc
	nopexp
	specinc
	specexp
	tax
)

func (c accClass) String() string {
	switch c {
	case currentAsset:
		return "currentAsset"
	case fixedAsset:
		return "fixedAsset"
	case currentLiability:
		return "currentLiability"
	case fixedLiability:
		return "fixedLiability"
	case equity:
		return "equity"
	case sales:
		return "sales"
	case cogs:
		return "cogs"
	case sga:
		return "sga"
	case nopinc:
		return "nopinc"
	case nopexp:
		return "nopexp"
	case specinc:
		return "specinc"
	case specexp:
		return "specexp"
	case tax:
		return "tax"
	default:
		return "Unknown"
	}
}

type ColumnName string

const (
	accMaster ColumnName = "acc"
)

// AccMaster contains names and classes of accs.
type AccMaster struct {
	columnName ColumnName
	accName    []string
	accClass   []accClass
}

// LoadAccMaster loads data from a csv file.
func LoadAccMaster(n string) *AccMaster {
	a := AccMaster{
		columnName: accMaster,
		accName:    []string{},
		accClass:   []accClass{},
	}
	f, err := os.Open(n)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// Make AccMaster from rec([]string).
		a.accName = append(a.accName, rec[0])
		c, err := strconv.Atoi(rec[1])
		if err != nil {
			log.Fatal(err)
		}
		ac := accClass(c)
		a.accClass = append(a.accClass, ac)
	}
	return &a
}

// WriteAccMaster writes current AccMaster to csv a file.
func (a *AccMaster) WriteAccMaster(n string) {
	f, err := os.Create(n)
	if err != nil {
		log.Fatal(err)
	}
	w := csv.NewWriter(f)
	prep := a.transformDim()
	for _, record := range prep {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv", err)
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

func (a *AccMaster) transformDim() [][]string {
	tfm := make([][]string, 0)
	//temp := make([]string, 0)
	for i, n := range a.accName {
		temp := []string{}
		c := strconv.Itoa(int(a.accClass[i]))
		temp = append(temp, n, c)
		tfm = append(tfm, temp)
	}
	return tfm
}

// CreateAccMaster returns accMaster.
func CreateAccMaster() *AccMaster {
	return &AccMaster{
		columnName: accMaster,
		accName:    []string{},
		accClass:   []accClass{},
	}
}

// AddAccMaster adds name and accClass as master data.
func (a *AccMaster) AddAccMaster(n string, c accClass) {
	uniq := true
	for _, v := range a.accName {
		if n == v {
			uniq = false
		}
	}
	if uniq {
		a.accName = append(a.accName, n)
		a.accClass = append(a.accClass, c)
	} else {
		fmt.Printf("%v has already existed in the master.\n", n)
	}
}

// CheckShape returns whether the shape is ok or not.
func (a *AccMaster) CheckShape() bool {
	return len(a.accName) == len(a.accClass)
}

// ColumnName returns the name of the column.
func (a *AccMaster) ColumnName() ColumnName {
	return a.columnName
}

// List returns the list of the item.
func (a *AccMaster) List() []string {
	list := make([]string, 0)
	for _, v := range a.accName {
		list = append(list, v)
	}
	return list
}

//type rec struct {
//date     string
//no       uint
//acc      acc
//sub      sub
//div      div
//taxin    bool
//taxclass taxclass
//taxrate  taxrate
//amt      float64
//note     string
//}
