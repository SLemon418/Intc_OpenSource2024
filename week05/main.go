package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	var i int = 55
	var f float64 = 1.92

	fmt.Printf("%f %f\n", f, math.Floor(f))
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(i))
	fmt.Println("f is", f)
	fmt.Println("i is", i)
	fmt.Print("i is ", i, "\n")
	fmt.Println("i is", i)
	fmt.Printf("i is %d\n", i)
}
