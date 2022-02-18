package main

import "fmt"

func main() {

	var n int32 = 42
	fmt.Printf("%v, %T\n", n, n)

	var m uint32 = 42
	fmt.Printf("%v, %T", m, m)
}
