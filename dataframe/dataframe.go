package dataframe

type Series struct {
	Name     string
	Elements Elements
	t        Type
}

type Type interface{}

type Elements interface{}
