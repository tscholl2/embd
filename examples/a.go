package main

import "fmt"

//go:generate embd b.txt

func main() {
	fmt.Println(b)
}
