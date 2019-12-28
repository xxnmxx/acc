package acc

import(
	"fmt"
)

type Prof struct {
	Stock Stock
	Div [2]Div
	Rev [3]Rev
	Equity [2]Equity
}

type Stock struct {
	Issued int
	Treasury int
}

type Div struct {
	DivAmount int
	ExtraDiv int
}

type Rev struct{
	TaxableIncome int
	ExtraIncome int
	DivIncome int
	WhtOnDiv int
	Nol int
}

type Equity struct {
	TaxCapital int
	TaxRE int
}

//func (p *prof) amountCalc() int {
	//n := p.Equity[0].TaxCapital / 50 * 1000
	//return n
//}

func(p *Prof) Valuate()  {
	divs := []int{}
	//revs := []float64
	//eqs := []float64
	
	for _, div := range p.Div {
		d := div.vCalc()	
		divs = append(divs,d)
	} 
	fmt.Println(divs)
}
func (r *Rev) vCalc() int {
	v := r.TaxableIncome - r.ExtraIncome + r.DivIncome - r.WhtOnDiv + r.Nol
	return v
}

func (d Div) vCalc() int {
	v := d.DivAmount - d.ExtraDiv
	return v
}

func (e *Equity) vCalc() int {
	v := e.TaxCapital + e.TaxRE
	return v
}

type values interface{
	vCalc() int
}
