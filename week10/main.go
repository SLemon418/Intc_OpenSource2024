package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("정수 입력")
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

	counts := 0
	i := 1
	for i <= realNum/i {
		if realNum%i == 0 {
			counts++
			break
		}
		i++
	}
	if counts == 0 && realNum < 2 {
		fmt.Println(realNum, "은 소수 입니다.")
	} else {
		fmt.Println(realNum, "은 소수가 아닙니다.")
	}
}
