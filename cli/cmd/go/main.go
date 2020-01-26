package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arr := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter text:")
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			fmt.Println(text)
			arr = append(arr, text)
		} else {
			break
		}
	}
	sum := 0
	for _, v := range arr {
		intV, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		sum += intV
	}
	fmt.Printf("arr: %v\tsum: %v\n", arr, sum) //Why sum is 0 value???
}
