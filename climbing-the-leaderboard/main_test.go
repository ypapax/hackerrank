package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
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

func Test2000(t *testing.T) {
	as := assert.New(t)
	f, err := os.Open("./test2000.txt")
	if !as.NoError(err) {
		return
	}
	scores, alice, err := getInputArrays(bufio.NewReader(f))
	if !as.NoError(err) {
		return
	}
	act := climbingLeaderboard(scores, alice)
	fExp, err := os.Open("./test2000.expected.txt")
	if !as.NoError(err) {
		return
	}
	exp, err := getArray(bufio.NewReader(fExp))
	if !as.NoError(err) {
		return
	}
	if !as.Equal(exp, act) {
		return
	}

}

func getArray(reader *bufio.Reader) ([]int32, error) {
	var arr []int32
	for {
		str := readLine(reader)
		if len(str) == 0 {
			break
		}
		i, err := strconv.ParseInt(str, 10, 32)
		if err != nil {
			return nil, err
		}
		arr = append(arr, int32(i))
	}

	return arr, nil
}

func TestGetInputArr(t *testing.T) {
	as := assert.New(t)
	f, err := os.Open("./test2000_1.txt")
	if !as.NoError(err) {
		return
	}
	scores, err := getInputArr(bufio.NewReader(f))
	if !as.NoError(err) {
		return
	}
	if !as.Equal(200000, len(scores)) {
		return
	}
}

func TestMain(m *testing.M) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("TestMain")
	os.Exit(m.Run())
}
