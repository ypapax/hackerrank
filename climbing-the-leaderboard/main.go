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
	var ranks = make([]int32, len(alice))

	scoreRanks := toScoresRank(scores)

	/*for i := len(alice) - 1; i >= 0; i-- {
		aliceScore := alice[i]
		ranks[i], index = getRank(currentRank, scores, aliceScore)
		if aliceScore < scores[0] {
			for {
				if scores[0] == scores[1] {
					scores = scores[1:]
					//index++
					continue
				}
				break
			}
			scores = scores[index:]
			currentRank = ranks[i]
		}
	}*/

	for i, a := range alice {
		ranks[i], _ = getRank(0, len(scores)-1, scoreRanks, a)
	}
	return ranks
}

func toScoresRank(scores []int32) []scoreRank {
	var scoreRanks []scoreRank

	var rank int
	for i, sc := range scores {
		if i != 0 && scores[i-1] == sc {
			continue
		}
		rank++
		scoreRanks = append(scoreRanks, scoreRank{score: sc, rank: rank})
	}
	return scoreRanks
}

type scoreRank struct {
	score int32
	rank  int
}

func (sr *scoreRank) String() string {
	return fmt.Sprintf("%+v,%+v", sr.rank, sr.score)
}

func getRank(low, high int, scores []scoreRank, target int32) (rank int32, rankIndex int) {
	if low > high {
		return -1, -1
	}
	if low == high && low >= len(scores)-1 {
		return int32(scores[len(scores)-1].rank + 1), len(scores)
	}
	mid := (low + high) / 2
	if mid < 0 {
		panic("mid is negative")
	}
	if mid >= len(scores) {
		return int32(scores[len(scores)-1].rank + 1), len(scores)
	}
	if scores[mid].score == target {
		return int32(scores[mid].rank), mid
	}
	if scores[mid].score > target {
		return getRank(mid+1, high, scores, target)
	}

	prev := mid - 1
	if prev == 0 {
		return 1, 0
	}
	if prev < 0 {
		return 1, -1
	}
	if scores[prev].score > target {
		return int32(scores[mid].rank), mid
	}
	return getRank(low, mid-1, scores, target)
}

func getInputArrays(reader *bufio.Reader) ([]int32, []int32, error) {
	scores, err := getInputArr(reader)
	if err != nil {
		log.Println("error: ", err)
		return nil, nil, err
	}

	alice, err := getInputArr(reader)
	if err != nil {
		log.Println("error: ", err)
		return nil, nil, err
	}

	return scores, alice, nil
}

func getInputArr(reader *bufio.Reader) ([]int32, error) {
	count, err := strconv.ParseInt(readLine(reader), 10, 64)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}

	log.Println("count", count)
	line := readLine(reader)
	temp := strings.Split(line, " ")
	if len(temp) != int(count) {
		err := fmt.Errorf("amount of values %+v doesn't equal to count %+v", len(temp), count)
		log.Println("error: ", err)
		return nil, err
	}
	var arr []int32
	for i := 0; i < int(count); i++ {
		scoresItemTemp, err := strconv.ParseInt(temp[i], 10, 64)
		if err != nil {
			log.Println("error: ", err)
			return nil, err
		}
		scoresItem := int32(scoresItemTemp)
		arr = append(arr, scoresItem)
	}
	return arr, nil
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	scores, alice, err := getInputArrays(reader)
	if err != nil {
		log.Println("error: ", err)
		return
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
	var resultLine string
	for {
		str, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			return ""
		}
		resultLine += string(str)
		if !isPrefix {
			break
		}
	}

	return strings.TrimRight(resultLine, "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
