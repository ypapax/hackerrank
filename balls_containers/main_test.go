package main

import (
	"fmt"
	"testing"

	"os"

	"github.com/stretchr/testify/assert"
)

func TestArrangeMatrix(t *testing.T) {
	type testCase struct {
		in string
		o  bool
	}
	for _, tc := range []testCase{
		{"possible.txt", true},
		{"impossible.txt", false},
		{"2_big.txt", true},
		{"possible4.txt", true},
		{"possible100.txt", true},
	} {
		t.Run(fmt.Sprintf("test%+v", tc.in), func(t *testing.T) {
			as := assert.New(t)
			f, err := os.Open(tc.in)
			if !as.NoError(err) {
				t.Log(err)
				return
			}
			defer f.Close()
			mm, err := readMatrices(f)
			if !as.NoError(err) {
				t.Log(err)
				return
			}
			o := arrangeMatrix(mm[0])
			if !as.Equal(tc.o, o) {
				t.Error("should be equal")
			}
		})
	}
}

func TestSwap(t *testing.T) {
	type testCase struct {
		in [][]int
		sw swapping
		o  [][]int
	}
	cases := []testCase{
		{[][]int{
			[]int{1, 1},
			[]int{1, 1},
		}, newSwapping(0, 1, 1, 0, 1), [][]int{
			[]int{2, 0},
			[]int{0, 2}},
		},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("test%+v", tc.in), func(t *testing.T) {
			as := assert.New(t)
			t.Logf("sw %+v", tc.sw)
			o := swap(tc.in, tc.sw)
			if !as.Equal(tc.o, o) {
				t.Error("should be equal")
			}
		})
	}
}
