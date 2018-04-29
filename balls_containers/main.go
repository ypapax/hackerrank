package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"flag"
	"io"
)

var debug bool

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	bind := flag.String("bind", "", "API server listenning for algorithm debug")
	flag.Parse()

	if len(*bind) > 0 {
		arrangeAPI(*bind)
	} else {
		log.Println("bind option is not specified")
	}

	mm, err := readMatrices(os.Stdin)
	if err != nil {
		log.Println(err)
	}
	for i, m := range mm {
		log.Println("matrix", i)
		printMatrix(m)
		_, possible := arrangeMatrix(m, false)
		if possible {
			fmt.Println("Possible")
			continue
		}
		fmt.Println("Impossible")
	}
}

func arrangeMatrix(m [][]int, debug bool) ([][]int, bool) {
	if isArranged(m) {
		return m, true
	}

	for boxNumber, box := range m {
		for ballTypeNumber := range box {
			if boxNumber == ballTypeNumber {
				continue
			}
			targetBox := ballTypeNumber
			ballTypeNumber2 := boxNumber
			swapByBoxFromToAndBallNumber(boxNumber, ballTypeNumber, targetBox, ballTypeNumber2, m)
		}
		if isArranged(m) {
			return m, true
		}
	}
	arranged := isArranged(m)
	if !arranged && debug {
		m = debugArrange(m)
	}
	return m, isArranged(m)
}

func swapByBoxFromToAndBallNumber(boxNumber, ballTypeNumber, targetBox, ballTypeNumber2 int, m [][]int) error {
	sw := newSwapping(boxNumber, ballTypeNumber, targetBox, ballTypeNumber2)
	a1 := m[sw.ballMove1.from.row][sw.ballMove1.from.column]
	a2 := m[sw.ballMove2.from.row][sw.ballMove2.from.column]
	log.Println("box", boxNumber, "contains", a1, "balls of type", ballTypeNumber)
	log.Println("box", targetBox, "contains", a2, "balls of type", ballTypeNumber2)
	sw.amount = int(math.Min(float64(a1), float64(a2)))
	pr := func(msg string) {
		log.Println(msg)
		printMatrix(m, sw.ballMove1.from, sw.ballMove1.to, sw.ballMove2.from, sw.ballMove2.to)
	}

	if sw.amount == 0 {
		pr("swap amount is 0")
		return fmt.Errorf("swap amount is 0")
	}
	pr(fmt.Sprintf("before swap %+v", sw))
	m = swap(m, sw)
	pr("after swap")
	return nil
}

func isArranged(m [][]int) bool {
	for containerNumber, container := range m {
		typeDetected := false
		for _, amount := range container {
			if amount == 0 {
				continue
			}
			if typeDetected {
				log.Printf("container %+v has different types of balls\n", containerNumber)
				return false
			}
			typeDetected = true
		}
	}
	return true
}

func biggestAmount(m [][]int) int {
	max := m[0][0]
	for _, container := range m {
		for _, amount := range container {
			if amount > max {
				max = amount
			}
		}
	}
	return max
}

func swap(m [][]int, sw swapping) [][]int {
	for _, bm := range []ballMove{sw.ballMove1, sw.ballMove2} {
		m[bm.from.row][bm.from.column] -= sw.amount
		m[bm.to.row][bm.to.column] += sw.amount
		log.Printf("%+v balls of type %+v had moved from box %+v to box %+v\n", sw.amount, bm.from.column, bm.from.row, bm.to.row)
	}
	return m
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

type point struct {
	row, column int
	color       color
}

type color int

const (
	red     color = 31
	green   color = 32
	redBG   color = 101
	greenBG color = 102
)

func getColor(c color) string {
	if c == 0 {
		return "\033[0m" // turn off the color
	}
	return fmt.Sprintf("\033[0;%+vm", c) // https://misc.flogisoft.com/bash/tip_colors_and_formatting
}

func printMatrix(m [][]int, highlight ...point) {
	maxLength := func() int {
		var max int
		for _, r := range m {
			for _, v := range r {
				if l := len(fmt.Sprintf("%+v", v)); l > max {
					max = l
				}
			}
		}
		return max
	}()

	line := func() {
		l := "-"
		size := maxLength * len(m)
		for i := 0; i <= size; i++ {
			l += "-"
		}
		fmt.Fprint(os.Stderr, l, "\n")
	}
	line()
	for i, r := range m {
		for j, v := range r {
			_ = i
			_ = j
			var clr, noClr string
			for _, h := range highlight {
				if i == h.row && j == h.column {
					clr = getColor(h.color)
					noClr = getColor(0)
					break
				}

			}
			tab := " "
			delta := maxLength + 1 - len(fmt.Sprintf("%+v", v))
			for i := 0; i < delta; i++ {
				tab += " "
			}
			fmt.Fprint(os.Stderr, clr, v, noClr, tab)
		}
		fmt.Fprintln(os.Stderr)
	}
	line()
}

type ballMove struct {
	from, to point
}

type swapping struct {
	ballMove1, ballMove2 ballMove
	amount               int
}

func newSwapping(box1, ballType1, box2, ballType2 int) swapping {
	return swapping{
		ballMove1: ballMove{
			point{box1, ballType1, red}, point{box2, ballType1, green},
		},
		ballMove2: ballMove{
			point{box2, ballType2, redBG}, point{box1, ballType2, greenBG},
		},
		amount: 1,
	}
}
