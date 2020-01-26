package main

import (
	"bufio"
	"fmt"
	"os"
)

func Input() {
	fmt.Print(">> ")
	r := bufio.NewScanner(os.Stdin)
	for r.Scan() {
		fmt.Print(">> ")
		cmd := r.Text()
		switch cmd {
		case "exit":
			break
		default:
			fmt.Println(cmd)
			fmt.Print(">> ")
		}
	}
}

func main() {
	Input()
}
