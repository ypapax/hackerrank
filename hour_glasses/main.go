package main

import (
	"fmt"
	"log"
	"os"
)

const matrix_size = 6

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	s, err := largestHourGlassSum(os.Stdin)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(s)
}

func largestHourGlassSum(f *os.File) (int, error) {
	//var i int

	m, err := scanMatrix(matrix_size, f)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	var sum int
	for i := 1; i < matrix_size-1; i++ {
		log.Println("line-----------", i)
		for j := 1; j < matrix_size-1; j++ {
			log.Println("-------j", j)
			center := m[i][j]
			down := m[i+1][j]
			up := m[i-1][j]
			downleft := m[i+1][j-1]
			upleft := m[i-1][j-1]
			downright := m[i+1][j+1]

			upright := m[i-1][j+1]
			log.Println(upleft, down, upright)
			log.Println(" ", center, " ")
			log.Println(downleft, down, downright)
			cur_sum := center + down + up + downleft + upleft + downright + upright

			printMatrix(m, []int{i, j}, []int{i + 1, j}, []int{i - 1, j},
				[]int{i + 1, j - 1}, []int{i - 1, j - 1}, []int{i + 1, j + 1},
				[]int{i - 1, j + 1})
			log.Println("cur_sum", cur_sum)
			if cur_sum > sum {
				sum = cur_sum
			}
		}
	}

	return sum, nil
}

func scanMatrix(size int, f *os.File) ([][]int, error) {
	var m [][]int
	for i := 0; i < matrix_size; i++ {
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
