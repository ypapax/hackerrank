package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestHourGlassSum(t *testing.T) {
	type testCase struct {
		in int
		o  bool
	}
	for _, tc := range []testCase{
		{3, true},
		{12, false},
		{5, true},
		{7, true},
		{104729, true},
		{104730, false},
		{1, false},
	} {
		t.Run(fmt.Sprintf("test%+v", tc.in), func(t *testing.T) {
			as := assert.New(t)
			o := isPrime(tc.in)
			as.Equal(tc.o, o)
		})
	}
}
