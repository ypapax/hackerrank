package main

import (
	"log"
	"os"
	"fmt"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	mm, err := readMatrices(os.Stdin)
	if err != nil {
		log.Println(err)
	}
	for i, m := range mm {
		log.Println("matrix", i)
		printMatrix(m)
	}
}

func readMatrices(f *os.File) ([][][]int, error) {
	var matrixCount int
	if _, err := fmt.Scan(&matrixCount); err != nil {
		return nil, err
	}
	var mm [][][]int
	for i := 0; i<matrixCount; i++{
		var dim int
		_, err := fmt.Scan(&dim)
		if err != nil {
			return nil, err
		}
		m, err := scanMatrix(dim, f)
		if err != nil {
			return nil, err
		}
		mm = append(mm, m)
	}
	return mm, nil
}

func scanMatrix(size int, f *os.File) ([][]int, error) {
	var m [][]int
	for i := 0; i < size; i++ {
		a, err := scanSlice(size, f)
		if err != nil {
			log.Println(err)
			return m, err
		}
		m = append(m, a)
	}
	return m, nil
}

func scanSlice(size int, f *os.File) ([]int, error) {
	in := make([]int, size)
	for i := range in {
		_, err := fmt.Fscan(f, &in[i])
		if err != nil {
			return in[:i], err
		}
	}
	return in, nil
}

func printMatrix(m [][]int, highlight ...[]int) {
	for i, r := range m {
		for j, v := range r {
			_ = i
			_ = j
			var color, noColor string
			for _, h := range highlight {
				if i == h[0] && j == h[1] {
					color = "\033[0;31m" // red
					noColor = "\033[0m"
					break
				}

			}
			fmt.Fprint(os.Stderr, color, v, noColor, " ")
		}
		fmt.Fprintln(os.Stderr)
	}
}