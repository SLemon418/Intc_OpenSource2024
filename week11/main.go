package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getinteger() int {
	r := bufio.NewReader(os.Stdin)
	num, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	num = strings.TrimSpace(num)
	realNum, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal(err)
	}

	return realNum
}

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

	first := getinteger()
	fmt.Print("두번쨰 정수 입력 : ")
	last := getinteger()

	for i := first; i <= last; i++ {
		if isprime(i) {
			fmt.Printf("%d ", i)
		}
	}
}
