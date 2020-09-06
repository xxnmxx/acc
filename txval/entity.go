package txval

type Entity struct {
	Valuers []Valuer
}

type Valuer interface {
	NameOfMethod() string
	Valuation() float64
}
