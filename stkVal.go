package acc

type Prof struct {
	Stock  Stock
	Div    [2]Div
	Rev    [3]Rev
	Equity [2]Equity
}

type Stock struct {
	Issued   int
	Treasury int
}

type Div struct {
	DivAmount int
	ExtraDiv  int
}

type Rev struct {
	TaxableIncome int
	ExtraIncome   int
	DivIncome     int
	WhtOnDiv      int
	Nol           int
}

type Equity struct {
	TaxCapital int
	TaxRE      int
}

func (p *Prof) ICalv() []float64 {
	divs := []int{}
	revs := []int{}
	eqs := []int{}

	for _, div := range p.Div {
		d := div.vCalc()
		divs = append(divs, d)
	}
	
	for _, rev := range p.Rev {
		r := rev.vCalc()
		revs = append(revs, r)
	}
	vrevs := revSelect(revs)

	for _, eq := range p.Equity {
		e := eq.vCalc()
		eqs = append(eqs, e)
	}

	shares := float64(p.Equity[0].TaxCapital * 1000 / 50)

	idiv := RoundDown(float64(divs[0]) / shares, 2.0)
	irev := RoundDown(float64(vrevs[0]) / shares, 2.0)
	ieq := RoundDown(float64(eqs[0])/ shares, 2.0)

	is := []float64{idiv,irev,ieq}

	return is

}

func (r *Rev) vCalc() int {
	v := r.TaxableIncome - r.ExtraIncome + r.DivIncome - r.WhtOnDiv + r.Nol
	return v
}

func (d *Div) vCalc() int {
	v := d.DivAmount - d.ExtraDiv
	return v
}

func (e *Equity) vCalc() int {
	v := e.TaxCapital + e.TaxRE
	return v
}

func revSelect(r []int )[]int{
	r00 := r[0]
	r05 := int(RoundDown(float64((r[0] + r[1]) / 2), 2.0))
	r10 := r[1]
	r15 := int(RoundDown(float64((r[1] + r[2]) / 2), 2.0))

	vRevs := []int{}
	if r00 < r05 {
		vRevs = append(vRevs, r00)
	}else{
		vRevs = append(vRevs, r05)
	}
	// For One-element Company test.Get max value.
	if r10 < r15 {
		vRevs = append(vRevs, r15)
	}else{
		vRevs = append(vRevs, r10)
	}
	return vRevs
}

type values interface {
	vCalc() int
}
