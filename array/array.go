package main

import "fmt"

type GenInfo struct {
	Rows, Cols int
	Stride     int
	Data       []float64
}

type Matrix interface {
	Shape() (r, c int)
	At(r, c int) float64
	//T() Matrix
}

type Dense struct {
	mat              GenInfo
	capRows, capCols int
}

func (d *Dense) Shape() (r, c int)   { return d.mat.Rows, d.mat.Cols }
func (d *Dense) At(r, c int) float64 { return d.mat.Data[d.mat.Stride*(r-1)+(c-1)] }

func main() {
	d := []float64{1, 2, 3, 4, 5, 6}
	m := Dense{
		mat: GenInfo{
			Rows:   2,
			Cols:   3,
			Stride: 3,
			Data:   d,
		},
		capRows: 2,
		capCols: 3,
	}
	r, c := m.Shape()
	fmt.Println(r, c, m.At(1, 3))
}
