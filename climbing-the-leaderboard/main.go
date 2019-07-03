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
	var index int
	var currentRank int32 = 1
	for i := len(alice) - 1; i >= 0; i-- {
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
	}
	return ranks
}

func getRank(currentRank int32, scores []int32, aliceScore int32) (rank int32, rankIndex int) {

	for i, sc := range scores {
		//log.Printf("currentRank %+v", currentRank)
		//log.Printf("score %+v for %+v", sc, i)
		if i == 0 && aliceScore > sc {
			return 1, i
		}
		if aliceScore == sc {
			return currentRank, i
		}
		if i >= len(scores)-1 {
			break
		}
		if scores[i] > aliceScore && aliceScore > scores[i+1] {
			return currentRank + 1, i
		}
		if scores[i] != scores[i+1] {
			currentRank++
		}
	}
	return currentRank + 1, len(scores) - 1
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
