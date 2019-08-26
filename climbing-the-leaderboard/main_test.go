package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRank(t *testing.T) {
	type testCase struct {
		scores     []int32
		inpFile    string
		aliceScore int32
		expRank    int32
	}
	cases := []testCase{
		{scores: []int32{100, 90, 90, 80, 75, 60}, aliceScore: 50, expRank: 6},
		{scores: []int32{100, 100, 50, 40, 40, 20, 10}, aliceScore: 25, expRank: 4},
		{inpFile: "./test2000.txt", aliceScore: 5090, expRank: 199975},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%d-%d", c.aliceScore, c.expRank), func(t *testing.T) {
			as := assert.New(t)
			if len(c.inpFile) > 0 {
				f, err := os.Open("./test2000.txt")
				if !as.NoError(err) {
					return
				}
				scores, _, err := getInputArrays(bufio.NewReader(f))
				if !as.NoError(err) {
					return
				}
				c.scores = scores
			}
			_, a := getRank(c.scores, getRanks(c.scores), c.aliceScore)
			if !as.Equal(c.expRank, a) {
				return
			}
		})
	}
}

func TestBinarySearch(t *testing.T) {
	type testCase struct {
		inp           []int32
		target        int32
		expectedIndex int
	}
	cases := []testCase{
		{inp: []int32{100, 90, 90, 80, 75, 60}, target: 50, expectedIndex: 5},
		{inp: []int32{100, 90, 90, 80, 75, 60}, target: 100, expectedIndex: 0},
		{inp: []int32{100, 90, 80, 75, 60}, target: 90, expectedIndex: 1},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%d-%d", c.target, c.expectedIndex), func(t *testing.T) {
			as := assert.New(t)
			a := binarySearch(c.inp, len(c.inp)-1, 0, c.target)
			if !as.Equal(c.expectedIndex, a) {
				return
			}
		})
	}
}

func TestClimbingLeaderboard(t *testing.T) {
	type testCase struct {
		scores        []int32
		alice         []int32
		expAliceRanks []int32
	}
	var testCases = []testCase{
		{scores: []int32{100, 90, 90, 80, 75, 60}, alice: []int32{50, 65, 77, 90, 102}, expAliceRanks: []int32{6, 5, 4, 2, 1}},
	}
	for _, c := range testCases {
		t.Run(fmt.Sprintf("%+v %+v", c.scores, c.alice), func(t *testing.T) {
			as := assert.New(t)
			act := climbingLeaderboard(c.scores, c.alice)
			if !as.Equal(c.expAliceRanks, act) {
				return
			}
		})
	}
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
	//t.Log("exp", exp)
	//t.Log("act", act)
	t.Log("len exp", len(exp))
	t.Log("len act", len(act))
	if !as.Equal(exp, act) {
		for i, a := range act {
			e := exp[i]
			if a != e {
				t.Errorf("actual element number %+v is %+v which is not equal to expected %+v for alice score %+v", i, a, e, alice[i])
				return
			}
		}
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
	ret := m.Run()
	//log.Printf("binarySearchCalls: %+v", binarySearchCalls)
	os.Exit(ret)
}
