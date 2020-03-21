package main

import "fmt"

func main() {
	r := makeRec()
	fmt.Printf("%+v", *r)
}

var re = []interface{}{acc, div, amt}
var acc = []string{"cash", "land", "ar"}
var div = []string{"hq", "shop", "hotel"}
var amt = []int{}
var header = []string{"acc", "div", "amt"}

type rec struct {
	acc string
	div string
	amt int
}

func inputElem(l []string) string {
	for i, v := range l {
		fmt.Printf("%v. %v\t", i, v)
	}
	fmt.Print("\nSelect: ")
	var i int
	fmt.Scan(&i)
	return l[i]
}

func makeRec() *rec {
	r := new(rec)
	r.acc = inputElem(re[0].([]string))
	r.div = inputElem(re[1].([]string))
	var i int
	fmt.Print("Input Degit: ")
	fmt.Scan(&i)
	r.amt = i
	return r
}
