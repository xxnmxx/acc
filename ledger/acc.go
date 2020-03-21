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

// AccMaster contains names and classes of accs.
type AccMaster struct {
	accName  []string
	accClass []accClass
}

// LoadAccMaster loads data from a csv file.
func LoadAccMaster(n string) *AccMaster {
	m := new(AccMaster)
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
		m.accName = append(m.accName, rec[0])
		c, err := strconv.Atoi(rec[1])
		if err != nil {
			log.Fatal(err)
		}
		ac := accClass(c)
		m.accClass = append(m.accClass, ac)
	}
	return m
}

// WriteAccMaster writes current AccMaster to csv a file.
func (m *AccMaster) WriteAccMaster(n string) {
	f, err := os.Create(n)
	if err != nil {
		log.Fatal(err)
	}
	w := csv.NewWriter(f)
	prep := m.transformDim()
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

func (m *AccMaster) transformDim() [][]string {
	tfm := make([][]string, 0)
	//temp := make([]string, 0)
	for i, n := range m.accName {
		temp := []string{}
		c := strconv.Itoa(int(m.accClass[i]))
		temp = append(temp, n, c)
		tfm = append(tfm, temp)
	}
	return tfm
}

// CreateAccMaster returns accMaster.
func CreateAccMaster() *AccMaster {
	return &AccMaster{
		accName:  []string{},
		accClass: []accClass{},
	}
}

// AddAccMaster adds name and accClass as master data.
func (m *AccMaster) AddAccMaster(n string, c accClass) {
	uniq := true
	for _, v := range m.accName {
		if n == v {
			uniq = false
		}
	}
	if uniq {
		m.accName = append(m.accName, n)
		m.accClass = append(m.accClass, c)
	} else {
		fmt.Printf("%v has already existed in the master.\n", n)
	}
}

// CheckShape returns whether the shape is ok or not.
func (m *AccMaster) CheckShape() bool {
	return len(m.accName) == len(m.accClass)
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
