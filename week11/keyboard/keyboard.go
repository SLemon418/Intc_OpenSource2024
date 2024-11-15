package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Getinteger() (int, error) {
	r := bufio.NewReader(os.Stdin)
	num, err := r.ReadString('\n')
	if err != nil {
		return 0, err
	}

	num = strings.TrimSpace(num)
	realNum, err := strconv.Atoi(num)
	if err != nil {
		return 0, err
	}

	return realNum, nil
}

func Getfloat() (float64, error) {
	r := bufio.NewReader(os.Stdin)
	num, err := r.ReadString('\n')
	if err != nil {
		return 0, err
	}

	num = strings.TrimSpace(num)
	realNum, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return 0, err
	}

	return realNum, nil
}
