package main

import (
	"log"
	"math"
	"fmt"
	"io"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var c int
	_, err := fmt.Scan(&c)
	if err != nil {
		log.Println("error", err)
		return
	}
	log.Println(c, "test cases")
	for {
		var i int
		_, err := fmt.Scan(&i)
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Println("error", err)
		}
		if isPrime(i) {
			fmt.Println("Prime")
		} else {
			fmt.Println("Not prime")
		}
	}
}

func isPrime(n int) bool {
	log.Println("n", n)
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		log.Println("i", i)
		if n%i == 0 {
			return false
		}
	}
	return true
}
