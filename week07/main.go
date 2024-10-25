package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var now = time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	fmt.Println(int(month))
	fmt.Println("오늘은 ", year, "년", int(month), "월", day, "일 입니다.")
	fmt.Println("지금 시간은 ", now.Hour(), "시", now.Minute(), "분", now.Second(), "초 입니다.")

	var army = "우리는 !가와 !민에 충성을 다하는 대한민! 육군이다"
	armyFixed := strings.NewReplacer("!", "국")
	fmt.Println(army)
	fmt.Println(armyFixed.Replace(army))

	fmt.Println("이름 입력:")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(i)
	}

	score, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	score = strings.TrimSpace(score)
	realScore, err := strconv.ParseInt(score, 16, 32)
	var grade string
	if realScore >= 90 {
		grade = "A"
	} else {
		grade = "BCDF"
	}
	fmt.Printf("%d점은 %s등급 입니다\n", realScore, grade)
}
