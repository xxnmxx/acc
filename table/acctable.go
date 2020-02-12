// ToDo
// Make converter
// Fix Dr & Cr 2020/01/27
package table

import "strconv"

import "log"

// AccTable

type AccTable struct {
	Header []string
	Date   Records
	No     Records
	Acc    Records
	Sub    Records
	Dev    Records
	Tax    Records
	Amount Records
	Note   Records
}

type Records interface {
	Len() int
}

type DateRecs struct{ Recs []string }

func (r DateRecs) Len() int { return len(r.Recs) }
func SetDate(recs []string, c int) (r DateRecs) {
	for _, v := range recs {
		r.Recs = append(r.Recs, v)
	}
	return r
}

type NoRecs struct{ Recs []int }

func (r NoRecs) Len() int { return len(r.Recs) }
func SetNo(recs []string) (r NoRecs) {
	for _, v := range recs {
		intv, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		r.Recs = append(r.Recs, intv)
	}
	return r
}

type AccRecs struct{ Recs []string }

func (r AccRecs) Len() int { return len(r.Recs) }
func SetAcc(recs []string) (r AccRecs) {
	for _, v := range recs {
		r.Recs = append(r.Recs, v)
	}
	return r
}

type SubRecs struct{ Recs []string }

func (r SubRecs) Len() int { return len(r.Recs) }
func SetSub(recs []string) (r SubRecs) {
	for _, v := range recs {
		r.Recs = append(r.Recs, v)
	}
	return r
}

type DevRecs struct{ Recs []string }

func (r DevRecs) Len() int { return len(r.Recs) }
func SetDev(recs []string) (r DevRecs) {
	for _, v := range recs {
		r.Recs = append(r.Recs, v)
	}
	return r
}

type AmountRecs struct{ Recs []int }

func (r AmountRecs) Len() int { return len(r.Recs) }
func SetAmount(recs []string) (r AmountRecs) {
	for _, v := range recs {
		intv, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		r.Recs = append(r.Recs, intv)
	}
	return r
}

type TaxRecs struct{ Recs []string }

func (r TaxRecs) Len() int { return len(r.Recs) }
func SetTax(recs []string) (r TaxRecs) {
	for _, v := range recs {
		r.Recs = append(r.Recs, v)
	}
	return r
}

type NoteRecs struct{ Recs []string }

func (r NoteRecs) Len() int { return len(r.Recs) }
func SetNote(recs []string) (r NoteRecs) {
	for _, v := range recs {
		r.Recs = append(r.Recs, v)
	}
	return r
}

// Converter
var yayoiDr = []int{1, 2, 9, 10, 11, 12, 14, 23}  // Fix: add tax
var yayoiCr = []int{1, 2, 16, 17, 18, 19, 21, 23} // Fix: add tax

// Fix: Amount of Cr has some problems.
// Remove empty slices.
// Change cr sign is not working.
func (t *Table) ToGl() (gl Table) {
	// Prep for Conv.
	// Make Dr
	dr := make([][]string, 0)
	for _, slice := range t.Records {
		tempDr := make([]string, 0)
		for _, idx := range yayoiDr {
			tempDr = append(tempDr, slice[idx])
		}
		dr = append(dr, tempDr)
	}
	// Make Cr
	cr := make([][]string, 0)
	for _, slice := range t.Records {
		tempCr := make([]string, 0)
		for _, idx := range yayoiCr {
			tempCr = append(tempCr, slice[idx])
		}
		cr = append(cr, tempCr)
	}
	for i := range cr {
		cr[i][6] = "-" + cr[i][6]
	}
	// Concat Dr & Cr.
	marged := make([][]string, 0)
	for _, slice := range dr {
		tempMrg := make([]string, 0)
		for _, v := range slice {
			if slice[0] != "" && slice[2] != "" {
				tempMrg = append(tempMrg, v)
			}
		}
		marged = append(marged, tempMrg)
	}
	for _, slice := range cr {
		tempMrg := make([]string, 0)
		for _, v := range slice {
			if slice[0] != "" && slice[2] != "" {
				tempMrg = append(tempMrg, v)
			}
		}
	}
	gl.header = []string{"Date", "No", "Acc", "Sub", "dev", "Tax", "Amount", "Note"}
	gl.Records = marged
	return gl
}

// Make AccTable.
//a := new(AccTable)
//a.Date = SetDate(marged[0])
//a.No = SetNo(marged[1])
//a.Acc = SetAcc(marged[2])
//a.Sub = SetSub(marged[3])
//a.Dev = SetDev(marged[4])
//a.Tax = SetTax(marged[5])
//a.Amount = SetAmount(marged[6])
//a.Note = SetNote(marged[7])
