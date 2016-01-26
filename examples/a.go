package foo

import "fmt"

//go:generate embd b.txt

func main() {
	fmt.Println(b)
}
