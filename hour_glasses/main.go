package main

import (
	"fmt"
	"log"
	"os"
)

const matrix_size = 6

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	largestHourGlassSum(os.Stdin)
}

func largestHourGlassSum(f *os.File) (int, error) {
	//var i int

	m, err := scanMatrix(matrix_size, f)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	fmt.Println(m)
	return 0, fmt.Errorf("not implemented")
}

func scanMatrix(size int, f *os.File) ([][]int, error) {
	var m [][]int
	for i:=0; i<matrix_size; i++ {
		a, err := scanSlice(matrix_size, f)
		if err != nil {
			log.Println(err)
			return m, err
		}
		m = append(m, a)
	}
	return m, nil
}

func scanSlice(size int, f *os.File) ([]int, error) { // https://stackoverflow.com/a/33811723/1024794
	in := make([]int, size)
	for i := range in {
		_, err := fmt.Fscan(f, &in[i])
		if err != nil {
			return in[:i], err
		}
	}
	return in, nil
}
