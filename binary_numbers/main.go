package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var n int
	fmt.Scan(&n)

	_, o, err := toBinaryAndOneCount(n)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(os.Stdout, o)
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
