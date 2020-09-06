package dcf

type Entity struct {
	cf            []float64
	taxRate       float64
	debt          float64
	bvEquity      float64
	fvEquity      float64
	ev            float64
	kd            float64
	ke            float64
	riskPremium   float64
	riskFree      float64
	compEnt       [][]float64 // lb, Debt, Equity(fmv), taxRate
	unleveredBeta float64
	leveredBeta   float64
	wacc          float64
}

func NewEntity() *Entity {
	return &Entity{
		fvEquity: 10.0,
	}
}

// Set methods.
func (e *Entity) SetCf(cf []float64)        { e.cf = cf }
func (e *Entity) SetTaxRate(tr float64)     { e.taxRate = tr }
func (e *Entity) SetDebt(d float64)         { e.debt = d }
func (e *Entity) SetKd(i float64)           { e.kd = Kd(i, e.taxRate) }
func (e *Entity) SetRiskPremium(rp float64) { e.riskPremium = rp }
func (e *Entity) SetRiskFree(rf float64)    { e.riskFree = rf }
func (e *Entity) SetCompEnt(lb, debt, fvEq, taxRate float64) {
	ce := []float64{lb, debt, fvEq, taxRate}
	e.compEnt = append(e.compEnt, ce)
}

func (e *Entity) SetKe() {
	e.ke = Ke(e.riskPremium, e.leveredBeta, e.riskFree)
}

func (e *Entity) SetUnleveredBeta() {
	ttl := 0.0
	for _, v := range e.compEnt {
		ttl += UnleveredBeta(v[0], v[1], v[2], v[3])
	}
	e.unleveredBeta = ttl / float64(len(e.compEnt))
}

func (e *Entity) SetLeveredBeta() {
	e.leveredBeta = LeveredBeta(e.unleveredBeta, e.debt, e.fvEquity, e.taxRate)
}

func (e *Entity) SetWacc() {
	e.wacc = SimpleWacc(e.kd, e.ke, e.debt, e.fvEquity)
}

func (e *Entity) SetEv() {
	e.ev = Discount(e.cf, e.wacc)
}

func (e *Entity) SetFvEquity() {
	e.fvEquity = e.ev - e.debt
}

// Calculation methods.
func (e *Entity) Calc() {
	e.SetUnleveredBeta()
	e.SetLeveredBeta()
	e.SetKe()
	e.SetWacc()
	e.SetEv()
	e.SetFvEquity()
}

func (e *Entity) Eval() {
	for e.ev != Discount(e.cf, e.wacc) {
		e.Calc()
	}
}

// Get methods.
//func (e *Entity) CashFlow() float64             { return e.cf }
func (e *Entity) TaxRate() float64              { return e.taxRate }
func (e *Entity) Equity() float64               { return e.fvEquity }
func (e *Entity) EntityValue() float64          { return e.ev }
func (e *Entity) Kd() float64                   { return e.kd }
func (e *Entity) Ke() float64                   { return e.ke }
func (e *Entity) RiskPremium() float64          { return e.riskPremium }
func (e *Entity) RiskFree() float64             { return e.riskFree }
func (e *Entity) ComparableEntity() [][]float64 { return e.compEnt }
func (e *Entity) UnleveredBeta() float64        { return e.unleveredBeta }
func (e *Entity) LeveredBeta() float64          { return e.leveredBeta }
func (e *Entity) Wacc() float64                 { return e.wacc }
