package stock

import (
	"math"

	"github.com/xxnmxx/acc"
)

func (d *Data) EqParShares() (eps int) {
	eq := d.TaxReport0.Equity.Capital
	s := d.Base.IssuedStock - d.Base.TreasuryStock
	eps = int(acc.RoundDown(float64(eq)/float64(s), 0.0))
	return eps
}

func (d *Data) Shares() int {
	s := (d.TaxReport0.Equity.Capital * 1000) / 50
	return s
}

func (d *Data) Div() (b0, b1 float64) {
	div0 := d.TaxReport0.Div.Normal - d.TaxReport0.Div.Extra
	div1 := d.TaxReport1.Div.Normal - d.TaxReport1.Div.Extra
	div2 := d.TaxReport2.Div.Normal - d.TaxReport2.Div.Extra
	tempb0 := acc.RoundDown((float64(div0)+float64(div1))/2, 0.0)
	tempb1 := acc.RoundDown((float64(div1)+float64(div2))/2, 0.0)
	b0 = acc.RoundDown(tempb0*1000/float64(d.Shares()), 1.0)
	b1 = acc.RoundDown(tempb1*1000/float64(d.Shares()), 1.0)
	return b0, b1
}

func (d *Data) Income() (c0, c1, c2 float64) {
	inc0 := d.TaxReport0.Income.Income - d.TaxReport0.Income.Extra + d.TaxReport0.Income.ExemptDiv - d.TaxReport0.Income.Wht + d.TaxReport0.Income.Nol
	inc1 := d.TaxReport1.Income.Income - d.TaxReport1.Income.Extra + d.TaxReport1.Income.ExemptDiv - d.TaxReport1.Income.Wht + d.TaxReport1.Income.Nol
	inc2 := d.TaxReport2.Income.Income - d.TaxReport2.Income.Extra + d.TaxReport2.Income.ExemptDiv - d.TaxReport2.Income.Wht + d.TaxReport2.Income.Nol
	tempc0 := acc.RoundDown(float64(inc0*1000)/float64(d.Shares()), 0.0)
	tempc1 := acc.RoundDown((float64(inc0*1000+inc1*1000)/2)/float64(d.Shares()), 0.0)
	tempc2 := acc.RoundDown(float64(inc1*1000)/float64(d.Shares()), 0.0)
	tempc3 := acc.RoundDown((float64(inc1*1000+inc2*1000)/2)/float64(d.Shares()), 0.0)
	c0 = math.Min(float64(tempc0), float64(tempc1))
	c1 = math.Max(float64(tempc0), float64(tempc1))
	c2 = math.Max(float64(tempc2), float64(tempc3))
	return c0, c1, c2
}

func (d *Data) Equity() (d0, d1 float64) {
	eq0 := d.TaxReport0.Equity.Capital + d.TaxReport0.Equity.Re
	eq1 := d.TaxReport1.Equity.Capital + d.TaxReport1.Equity.Re
	d0 = float64(acc.RoundDown(float64(eq0*1000)/float64(d.Shares()), 0.0))
	d1 = float64(acc.RoundDown(float64(eq1*1000)/float64(d.Shares()), 0.0))
	return d0, d1
}

func (d *Data) Ratio(i Index) (ratio, rb, rc, rd float64) {
	b0, _ := d.Div()
	c0, _, _ := d.Income()
	d0, _ := d.Equity()
	rb = acc.RoundDown(b0/float64(i.B), 2.0)
	rc = acc.RoundDown(c0/float64(i.C), 2.0)
	rd = acc.RoundDown(d0/float64(i.D), 2.0)
	ratio = acc.RoundDown((rb+rc+rd)/3.0, 1.0)
	return ratio, rb, rc, rd
}

func (d *Data) ValueParJPY50(i Index) (vp50 float64) {
	listOfa := []int{i.A.A1, i.A.A2, i.A.A3, i.A.A4, i.A.A5}
	a := listOfa[0]
	for _, v := range listOfa {
		if a > v {
			a = v
		}
	}
	ratio, _, _, _ := d.Ratio(i)
	vp50 = acc.RoundDown(float64(a)*ratio*d.Base.Size, 1.0)
	return vp50
}

func (d *Data) ValueParStock() (vps int) {
	vps1 := acc.RoundDown(d.ValueParJPY50(d.Index1)*(float64(d.EqParShares())/50.0), 0.0)
	vps2 := acc.RoundDown(d.ValueParJPY50(d.Index2)*(float64(d.EqParShares())/50.0), 0.0)
	vps = int(math.Min(vps1, vps2))
	return vps
}
