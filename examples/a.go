package foo

import "fmt"

//go:generate embd b.txt
//go:generate embd -n c dir/hello.txt

func main() {
	fmt.Println(b)
	fmt.Println(c)
}
