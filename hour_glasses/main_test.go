package main

import (
	"fmt"
	"testing"

	"os"

	"github.com/stretchr/testify/assert"
)

func TestLargestHourGlassSum(t *testing.T) {
	type testCase struct {
		in string
		o  int
	}
	for _, tc := range []testCase{
		{`19.txt`, 19},
	} {
		t.Run(fmt.Sprintf("test%+v", tc.in), func(t *testing.T) { // https://youtu.be/hVFEV-ieeew?t=1037
			as := assert.New(t)
			file, err := os.Open(tc.in) // For read access.
			if !as.NoError(err) {
				return
			}
			defer file.Close()
			o, err := largestHourGlassSum(file)
			as.NoError(err)
			as.Equal(tc.o, o)
		})
	}
}
