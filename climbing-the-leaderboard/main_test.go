package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbingLeaderboard(t *testing.T) {
	log.Println("TestClimbingLeaderboard")
	as := assert.New(t)
	act := climbingLeaderboard([]int32{100, 90, 90, 80, 75, 60}, []int32{50, 65, 77, 90, 102})
	exp := []int32{6, 5, 4, 2, 1}
	as.Equal(exp, act)
}

func TestMain(m *testing.M) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("TestMain")
	os.Exit(m.Run())
}
