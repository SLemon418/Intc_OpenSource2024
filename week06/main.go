package main

import (
	"fmt"
)

func main() {
	i := 13
	f := 12.9
	c1 := 'Z'
	c2 := '김'

	fmt.Printf("value i: %d, value f : %f\n", i, f)
	fmt.Printf("i * f = %f\n", float64(i)*f)
	fmt.Println("c1:", c1, "c2:", c2)
	fmt.Printf("%x", c2)
}
