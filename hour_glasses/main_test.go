package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestHourGlassSum(t *testing.T) {
	type testCase struct {
		in string
		o int
	}
	for _, tc := range []testCase{
		{`1 1 1 0 0 0
0 1 0 0 0 0
1 1 1 0 0 0
0 0 2 4 4 0
0 0 0 2 0 0
0 0 1 2 4 0`, 19},
	} {
		t.Run(fmt.Sprintf("test%+v", tc.in), func(t *testing.T) { // https://youtu.be/hVFEV-ieeew?t=1037
			r, o, err := largestHourGlassSum(tc.in)
			as := assert.New(t)
			as.NoError(err)
			as.Equal(tc.b, r)
			as.Equal(tc.o, o)
		})
	}
}
