package acc

type Prof struct {
	Stock     Stock
	Div       [2]Div
	Rev       [3]Rev
	Equity    [2]Equity
	Industory [2]Industory
	L         float64
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

type Industory struct {
	Code int
	A    [5]int
	B    float64
	C    float64
	D    float64
}

func (p *Prof) Eval() []float64 {
	idx := p.ICalc()
	adj := float64(p.Equity[0].TaxCapital * 1000 / (p.Stock.Issued - p.Stock.Treasury))
	vs := []float64{}

	for i := range p.Industory {
		a := float64(p.Industory[i].selector())
		b := RoundDown(idx[0]/p.Industory[i].B, 1.0)
		c := RoundDown(idx[1]/p.Industory[i].C, 1.0)
		d := RoundDown(idx[2]/p.Industory[i].D, 1.0)
		prepv := RoundDown(a*(b+c+d)/3.0, 1.0)
		prepv2 := 0.0
		if p.L == 0.5 {
			prepv2 = RoundDown(prepv*0.5, 1.0)
		} else if p.L == 1 {
			prepv2 = RoundDown(prepv*0.7, 1.0)
		} else {
			prepv2 = RoundDown(prepv*0.6, 1.0)
		}
		pv := prepv2 * adj / 50.0
		vs = append(vs, pv)
	}
	return vs
}

func (p *Prof) ICalc() []float64 {
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
	vrevs := revSelector(revs)

	for _, eq := range p.Equity {
		e := eq.vCalc()
		eqs = append(eqs, e)
	}

	shares := float64(p.Equity[0].TaxCapital * 1000 / 50)
	preidiv := RoundDown(float64((divs[0]+divs[1])/2), 0.0)

	idiv := RoundDown(preidiv*1000/shares, 1.0)
	irev := RoundDown(float64(vrevs[0]*1000)/shares, 0.0)
	ieq := RoundDown(float64(eqs[0]*1000)/shares, 0.0)

	is := []float64{idiv, irev, ieq}
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

func (ind *Industory) selector() int {
	a := ind.A[0]
	for _, v := range ind.A {
		if a > v {
			a = v
		}
	}
	return a
}

func revSelector(r []int) []int {
	r00 := r[0]
	r05 := int(RoundDown(float64((r[0]+r[1])/2), 2.0))
	r10 := r[1]
	r15 := int(RoundDown(float64((r[1]+r[2])/2), 2.0))

	vRevs := []int{}
	if r00 < r05 {
		vRevs = append(vRevs, r00)
	} else {
		vRevs = append(vRevs, r05)
	}
	// For One-element Company test.Get max value.
	if r10 < r15 {
		vRevs = append(vRevs, r15)
	} else {
		vRevs = append(vRevs, r10)
	}
	return vRevs
}

type values interface {
	vCalc() int
}
