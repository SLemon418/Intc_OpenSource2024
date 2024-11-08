package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	fmt.Print("정수 입력 : ")
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

	if isprime(realNum) {
		fmt.Println(realNum, "은 소수 입니다.")
	} else {
		fmt.Println(realNum, "은 소수가 아닙니다.")
	}
}
