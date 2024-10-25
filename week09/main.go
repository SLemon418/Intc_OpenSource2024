package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	answer := rand.Intn(6) + 1
	fmt.Println("주사위 값", answer)
}
