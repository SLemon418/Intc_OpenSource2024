package keyboard

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Getinteger() int {
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
