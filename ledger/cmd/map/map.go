package main

import "fmt"

func main() {
	r := makeTemp(header)
	for k, v := range r {
		fmt.Println(k, v)
	}
}

var header = []string{"acc", "div", "amt"}

type rec struct {
	acc string
	div string
	amt int
}
