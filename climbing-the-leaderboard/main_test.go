package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRank(t *testing.T) {
	as := assert.New(t)
	act := getRank([]int32{100, 90, 90, 80, 75, 60}, 50)
	exp := int32(6)
	as.Equal(exp, act)
}

func TestRank2(t *testing.T) {
	as := assert.New(t)
	act := getRank([]int32{100, 100, 50, 40, 40, 20, 10}, 25)
	exp := int32(4)
	as.Equal(exp, act)
}

func TestClimbingLeaderboard(t *testing.T) {
	as := assert.New(t)
	act := climbingLeaderboard([]int32{100, 90, 90, 80, 75, 60}, []int32{50, 65, 77, 90, 102})
	exp := []int32{6, 5, 4, 2, 1}
	as.Equal(exp, act)
}

func TestClimbingLeaderboard2(t *testing.T) {
	as := assert.New(t)
	act := climbingLeaderboard([]int32{100, 100, 50, 40, 40, 20, 10}, []int32{5, 25, 50, 120})
	exp := []int32{6, 4, 2, 1}
	as.Equal(exp, act)
}

func TestMain(m *testing.M) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("TestMain")
	os.Exit(m.Run())
}
