package main

import (
	"fmt"

	"github.com/xxnmxx/acc/txval"
)

func main() {
	b := []float64{100,0}
	c := []float64{1000,-1000,100,-20,0}
	d := []float64{1000,3000}
	idx := []float64{364,343,321,287,366,6.8,51,293}
	s := txval.NewS4()
	s.SetCap(1000)
	s.SetStk(20)
	s.SetB(0,b)
	s.SetB(1,b)
	s.SetB(2,b)
	s.SetC(0,c)
	s.SetC(1,c)
	s.SetC(2,c)
	s.SetD(0,d)
	s.SetD(1,d)
	s.SetIdx(0,idx)
	s.SetIdx(1,idx)
	fmt.Println(s.Inspect())
}
