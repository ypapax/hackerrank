package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// Complete the climbingLeaderboard function below.
func climbingLeaderboard(scores []int32, alice []int32) []int32 {
	log.Println("alice", alice)
	var ranks = make([]int32, len(alice))
	var currentRank int32 = 1
	for i, sc := range scores {
		log.Printf("score %+v for %+v", sc, i)
		for j, a := range alice {
			if ranks[j] > 0 {
				log.Printf("rank %+v for %+v is already set", ranks[j], j)
				continue
			}
			if i == 0 && a > sc {
				ranks[j] = 1
			}
			if a == sc {
				log.Printf("setting rank %+v, for alice %+v\n", currentRank, alice[j])
				ranks[j] = currentRank + 1
			}
			if i >= len(scores)-1 {
				continue
			}
			if scores[i] > a && a > scores[i+1] {
				ranks[j] = currentRank + 2
				log.Printf("setting rank %+v, for alice %+v\n", currentRank+1, alice[j])
			}
		}
		if i == 0 {
			continue
		}
		if scores[i-1] != scores[i] {
			currentRank++
		}
	}
	for k, a := range ranks {
		if a != 0 {
			continue
		}
		ranks[k] = currentRank + 1
	}
	return ranks
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	scoresCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	scoresTemp := strings.Split(readLine(reader), " ")

	var scores []int32

	for i := 0; i < int(scoresCount); i++ {
		scoresItemTemp, err := strconv.ParseInt(scoresTemp[i], 10, 64)
		checkError(err)
		scoresItem := int32(scoresItemTemp)
		scores = append(scores, scoresItem)
	}

	aliceCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	aliceTemp := strings.Split(readLine(reader), " ")

	var alice []int32

	for i := 0; i < int(aliceCount); i++ {
		aliceItemTemp, err := strconv.ParseInt(aliceTemp[i], 10, 64)
		checkError(err)
		aliceItem := int32(aliceItemTemp)
		alice = append(alice, aliceItem)
	}

	result := climbingLeaderboard(scores, alice)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
