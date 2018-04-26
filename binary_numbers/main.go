package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
	"log"
)

func main() {
	decimal, err := readInt()
	if err != nil {
		log.Println(err)
		return
	}
	_, o, err := toBinaryAndOneCount(decimal)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(o)
}

func readInt() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	i, err := strconv.Atoi(strings.TrimSpace(text))
	if err != nil {
		return 0, err
	}
	return i, nil
}

func toBinaryAndOneCount(n int) (string, int, error) {
	r := strconv.FormatInt(int64(n), 2)
	var max, cur int
	for _, c := range r {
		switch c {
		case '1':
			cur++
		case '0':
			if cur > max {
				max = cur
			}
			cur = 0
		default:
			return r, 0, fmt.Errorf("not support character %+v, expected: 0 or 1", c)
		}
	}
	return r, max, nil
}