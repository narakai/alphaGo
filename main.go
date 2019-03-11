package main

import "fmt"

func main() {
	const s = "Go语言"
	for i, r := range s {
		fmt.Printf("%#U  ： %d\n", r, i)
	}
}
