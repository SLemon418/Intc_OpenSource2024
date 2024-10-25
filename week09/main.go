package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func main() {
	success := false
	rand.Seed(time.Now().UnixNano())
	answer := rand.Intn(6) + 1
	fmt.Println(answer)

	for guesses := 0; guesses < 3; guesses++ {
		fmt.Printf("%d번의 기회가 남았습니다. 주사위 값을 추측해보세요: ", 3-guesses)
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("숫자만 입력하세요.")
			guesses--
		} else {
			if guess == answer {
				fmt.Println("맞았습니다!")
				success = true
				break
			} else if guess < answer {
				fmt.Println("틀렸습니다. 추측한 값이 작습니다.")
			} else {
				fmt.Println("틀렸습니다. 추측한 값이 큽니다.")
			}
		}
	}
	if !success {
		fmt.Printf("실패했습니다. 주사위 값은 %d 입니다.\n", answer)
	}

	r := fmt.Sprintf("%d\n", rand.Intn(6)+1)
	fmt.Println(reflect.TypeOf(r))
	fmt.Printf("%T\n", 2.1)

	i := 1
	for i <= 100 { // while
		fmt.Printf("%3d점\n", i)
		i = i + 1
	}
}
