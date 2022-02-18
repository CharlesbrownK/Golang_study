package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Hello World!")
	// fmt.Println(5*2)

	var i int
	i = 42

	var j float32 = 27
	k := 99

	fmt.Println(i)
	fmt.Printf("%v, %T", k, k)
	fmt.Println("")
	fmt.Printf("%f, %T", j, j)
}
