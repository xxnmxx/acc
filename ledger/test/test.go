package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	home, _ := os.UserHomeDir()
	p := filepath.Join(home, "Downloads", "acc.text")
	fmt.Println(p)
	fmt.Printf("%T\t%v", p, p)
}
