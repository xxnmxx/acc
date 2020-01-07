package series

type Series struct {
	Name     string
	elements Elements
	t        Type
	Err      error
}

type Elements interface {
	Elem(i int) Element
	Len() int
}

type Element interface {
}
