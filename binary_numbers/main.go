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
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	decimal, err := readIntFromStdin()
	if err != nil {
		log.Println(err)
		return
	}
	_, o, err := toBinaryAndOneCount(decimal)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(os.Stdout, o)
}

func readIntFromStdin() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
		return 0, err
	}
	i, err := strconv.Atoi(strings.TrimSpace(text))
	if err != nil {
		log.Println(err)
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
			if cur > max {
				max = cur
			}
		case '0':
			cur = 0
		default:
			return r, 0, fmt.Errorf("not support character %+v, expected: 0 or 1", c)
		}
	}
	return r, max, nil
}