package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToBinaryAndOneCount(t *testing.T) {
	type testCase struct {
		in, o int
		b     string
	}
	for _, tc := range []testCase{
		{13, 2, "1101"},
		{5, 1, "101"},
		{15, 4, "1111"},
	} {
		t.Run(fmt.Sprintf("test%+v", tc.in), func(t *testing.T) { // https://youtu.be/hVFEV-ieeew?t=1037
			r, o, err := toBinaryAndOneCount(tc.in)
			as := assert.New(t)
			as.NoError(err)
			as.Equal(tc.b, r)
			as.Equal(tc.o, o)
		})
	}
}
