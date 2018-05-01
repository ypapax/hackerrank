package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"sort"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	//bind := flag.String("bind", "", "API server listenning for algorithm debug")
	//flag.Parse()
	//
	//if len(*bind) > 0 {
	//	arrangeAPI(*bind)
	//} else {
	//	log.Println("bind option is not specified")
	//}

	mm, err := readMatrices(os.Stdin)
	if err != nil {
		log.Println(err)
	}
	for i, m := range mm {
		log.Println("matrix", i)
		possible := isPossibleToArrange(m)
		if possible {
			fmt.Println("Possible")
			continue
		}
		fmt.Println("Impossible")
	}
}

func isPossibleToArrange(m [][]int) bool {
	var rowSum, colSum = make([]int, len(m)), make([]int, len(m))

	for ri, r := range m {
		for ci := range r {
			rowSum[ri] += m[ri][ci]
			colSum[ci] += m[ri][ci]
		}
	}

	return isSame(rowSum, colSum)
}

func isSame(a, b []int) bool {
	sort.Ints(a)
	sort.Ints(b)
	return reflect.DeepEqual(a, b)
}

func readMatrices(f io.Reader) ([][][]int, error) {
	var matrixCount int
	if _, err := fmt.Fscan(f, &matrixCount); err != nil {
		log.Println(err)
		return nil, err
	}
	var mm [][][]int
	for i := 0; i < matrixCount; i++ {
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

func scanMatrix(size int, f io.Reader) ([][]int, error) {
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

func scanSlice(size int, f io.Reader) ([]int, error) {
	in := make([]int, size)
	for i := range in {
		_, err := fmt.Fscan(f, &in[i])
		if err != nil {
			return in[:i], err
		}
	}
	return in, nil
}
