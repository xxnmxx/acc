package txval

import (
	"fmt"
	"strings"

	"github.com/xxnmxx/acc"
)

// wip implement index calc. A B C D.
type S4 struct {
	Cap float64
	Stk float64
	B   [][]float64 // shape is 3,2
	C   [][]float64 // shape is 3,5
	D   [][]float64 // shape is 2,2
	Idx [][]float64 // shape is 2,8. A:0-4, B:5, C:6, D:7
}

func NewS4() *S4 {
	return &S4{
		B:   newMat(3, 2),
		C:   newMat(3, 5),
		D:   newMat(2, 2),
		Idx: newMat(2, 8),
	}
}

func newMat(r, c int) [][]float64 {
	mat := make([][]float64, r)
	for i := range mat {
		mat[i] = make([]float64, c)
	}
	return mat
}

// Methods for print.
// Inspect returns string for printout.
func (s4 *S4) Inspect() string {
	var b strings.Builder
	hb := "[+]div\t[-]ex\tttl"
	hc := "[+]inc\t[-]ex\t[+]ntx\t[-]wht\t[+]nol\tttl"
	hd := "[+]re\t[+]cap\tttl"
	bs := s4.calcB()
	cs := s4.calcC()
	ds := s4.calcD()
	vs := s4.valsPar50()
	val := s4.Val()
	b.WriteString("val: " + fmt.Sprint(val) + ", " + "v1: " + fmt.Sprint(vs[0]) + ", " + "v2: " + fmt.Sprint(vs[1]) + "\n")
	b.WriteString("b1: " + fmt.Sprint(bs[0]) + ", " + "b2: " + fmt.Sprint(bs[1]) + "\n")
	b.WriteString(printEle(hb, s4.B))
	b.WriteString("c1: " + fmt.Sprint(cs[0]) + ", " + "c1av: " + fmt.Sprint(cs[1]) + ", " + "c2: " + fmt.Sprint(cs[2]) + ", " + "c2av: " + fmt.Sprint(cs[3]) + "\n")
	b.WriteString(printEle(hc, s4.C))
	b.WriteString("d1: " + fmt.Sprint(ds[0]) + ", " + "d2: " + fmt.Sprint(ds[1]) + "\n")
	b.WriteString(printEle(hd, s4.D))
	return b.String()
}

func printEle(h string, d [][]float64) string {
	var b strings.Builder
	b.WriteString(h + "\n")
	for _, v := range d {
		for _, k := range v {
			strK := fmt.Sprint(k)
			b.WriteString(strK + "\t")
		}
		strSum := fmt.Sprint(sumSlice(v))
		b.WriteString(strSum)
		b.WriteString("\n")
	}
	return b.String()
}

func sumSlice(s []float64) float64 {
	var sum float64
	for _, v := range s {
		sum += v
	}
	return sum
}

// Methods for caluclation.
// calcB returns slice of b1:0, b2:1.
func (s4 *S4) calcB() []float64 {
	bs := make([]float64, 2)
	bs[0] = acc.RoundDown((sumSlice(s4.B[0])+sumSlice(s4.B[1]))/2, 0)
	bs[0] = acc.RoundDown(bs[0]/s4.stkPar50()*1000, 1)
	bs[1] = acc.RoundDown((sumSlice(s4.B[1])+sumSlice(s4.B[2]))/2, 0)
	bs[1] = acc.RoundDown(bs[1]/s4.stkPar50()*1000, 1)
	return bs
}

// calcC returns slice of c1:0, c1av:1, c2:2, c2av:3.
func (s4 *S4) calcC() []float64 {
	cs := make([]float64, 4)
	cs[0] = acc.RoundDown(sumSlice(s4.C[0])/s4.stkPar50()*1000, 0)
	cs[1] = acc.RoundDown((sumSlice(s4.C[0])+sumSlice(s4.C[1]))/2/s4.stkPar50()*1000, 0)
	cs[2] = acc.RoundDown(sumSlice(s4.C[1])/s4.stkPar50()*1000, 0)
	cs[3] = acc.RoundDown((sumSlice(s4.C[1])+sumSlice(s4.C[2]))/2/s4.stkPar50()*1000, 0)
	return cs
}

// calcD returns slice of d1:0, d2:1.
func (s4 *S4) calcD() []float64 {
	ds := make([]float64, 2)
	ds[0] = acc.RoundDown(sumSlice(s4.D[0])/s4.stkPar50()*1000, 0)
	ds[1] = acc.RoundDown(sumSlice(s4.D[1])/s4.stkPar50()*1000, 0)
	return ds
}

func (s4 *S4) stkPar50() float64 { return acc.Round(s4.Cap/50, 0) }

func (s4 *S4) Val() float64 {
	v := min(s4.valsPar50())
	return acc.RoundDown(v*s4.Cap/50, 0)
}

// Breakdown of the 2nd dim of the s4.Idx
// A:0-4, B:5, C:6, D:7.
// wip
func (s4 *S4) valsPar50() []float64 {
	vals := make([]float64, 2)
	for i := range vals {
		vals[i] = s4.valPar50(i)
	}
	return vals
}

func (s4 *S4) valPar50(ix int) float64 {
	ixA := min(s4.Idx[ix][0:5])
	ixB := s4.Idx[ix][5]
	ixC := s4.Idx[ix][6]
	ixD := s4.Idx[ix][7]
	b := s4.calcB()[0]
	c := min(s4.calcC()[0:2])
	d := s4.calcD()[0]
	div := acc.RoundDown(b/ixB, 2)
	inc := acc.RoundDown(c/ixC, 2)
	eq := acc.RoundDown(d/ixD, 2)
	ratio := acc.RoundDown((div+inc+eq)/3, 2)
	val := acc.RoundDown(ixA*ratio*0.7, 1) // 0.7 is temp.
	return val
}

func min(s []float64) float64 {
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

// Methods for input.
// SetCap sets Cap. Cap is float.
func (s4 *S4) SetCap(f float64) { s4.Cap = f }
// SetStk sets Stk. Stk is float.
func (s4 *S4) SetStk(f float64) { s4.Stk = f }
// Set B sets B. B is slice shape of 3,2.
func (s4 *S4) SetB(i int, v []float64) { s4.B[i] = v }
// Set C sets C. C is slice shape of 3,5.
func (s4 *S4) SetC(i int, v []float64) { s4.C[i] = v }
// Set D sets D. D is slice shape of 2,2.
func (s4 *S4) SetD(i int, v []float64) { s4.D[i] = v }
// Set Idx sets Idx. Idx is slice shape of 2,8.
// A:0-4, B:5, C:6, D:7.
func (s4 *S4) SetIdx(i int, v []float64) { s4.Idx[i] = v }
