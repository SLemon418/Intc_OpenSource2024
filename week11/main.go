package main

import (
	"fmt"
	"week11/keyboard"
)

func isprime(num int) bool {
	if num < 2 {
		return false
	} else if num == 2 {
		return true
	} else if num%2 == 0 {
		return false
	} else {
		i := 3
		for i <= num/i {
			if num%i == 0 {
				return false
			}
			i += 2
		}
	}
	return true
}

func main() {
	fmt.Print("첫번쨰 정수 입력 : ")

	first := keyboard.Getinteger()
	fmt.Print("두번쨰 정수 입력 : ")
	last := keyboard.Getinteger()

	for i := first; i <= last; i++ {
		if isprime(i) {
			fmt.Printf("%d ", i)
		}
	}
}
