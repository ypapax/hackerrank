package main

import (
	"log"
	"math"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

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
