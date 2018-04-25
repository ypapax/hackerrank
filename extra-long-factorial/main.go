package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	i, err := strconv.Atoi(strings.TrimSpace(text))
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	r := factorial(int64(i))
	fmt.Println(r.String())
}

func factorial(n int64) big.Int {
	var f big.Int
	f.MulRange(1, n)
	return f
}
