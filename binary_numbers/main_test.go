package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestToBinaryAndOneCount(t *testing.T){
	type testCase struct {
		in, o int
		b string
	}
	for _, tc := range []testCase{{13, 2,  "1101"}} {
		r, o, err := toBinaryAndOneCount(tc.in)
		as := assert.New(t)
		as.NoError(err)
		as.Equal(tc.b, r)
		as.Equal(tc.o, o)
	}
}