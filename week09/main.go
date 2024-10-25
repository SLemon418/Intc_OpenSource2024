package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	answer := rand.Intn(6) + 1
	fmt.Println("주사위 값", answer)

	for guesses := 0; guesses < 3; guesses++ {
		fmt.Print("주사위 값을 추측해보세요: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("숫자만 입력하세요.")
		} else {
			if guess == answer {
				fmt.Println("맞았습니다!")
				break
			} else if guess < answer {
				fmt.Println("틀렸습니다. 추측한 값이 작습니다.")
			} else {
				fmt.Println("틀렸습니다. 추측한 값이 큽니다.")
			}
		}
	}
}
