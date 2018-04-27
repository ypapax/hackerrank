package main

import (
	"log"
	"os"
	"fmt"
)

func init(){
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	mm, err := readMatrices(os.Stdin)
	if err != nil {
		log.Println(err)
	}
	for i, m := range mm {
		log.Println("matrix", i)
		printMatrix(m)
	}
}

func arrangeMatrix(m [][]int) bool {
	if isArranged(m) {
		return true
	}
	for boxNumber, box := range m {
		for ballTypeNumber := range box {
			if boxNumber == ballTypeNumber {
				continue
			}
			targetBox := ballTypeNumber
			m = swap(m, [2]int{boxNumber, ballTypeNumber}, [2]int{targetBox, ballTypeNumber})
		}
		if isArranged(m) {
			return true
		}
	}
	return isArranged(m)
}

func isArranged(m [][]int) bool {
	for _, container := range m {
		typeDetected := false
		for _, amount := range container {
			if amount == 0 {
				continue
			}
			if typeDetected {
				log.Printf("container %+v has different types of balls\n", container)
				return false
			}
			typeDetected = true
		}
	}
	return true
}

func swap(m [][]int, from, to [2]int) [][]int {
	m[from[0]][from[1]]--
	m[to[0]][to[1]]++
	return m
}

func readMatrices(f *os.File) ([][][]int, error) {
	var matrixCount int
	if _, err := fmt.Fscan(f, &matrixCount); err != nil {
		log.Println(err)
		return nil, err
	}
	var mm [][][]int
	for i := 0; i<matrixCount; i++{
		var dim int
		_, err := fmt.Fscan(f, &dim)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		m, err := scanMatrix(dim, f)
		if err != nil {
			log.Println(err)
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