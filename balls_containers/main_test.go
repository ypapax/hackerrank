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
