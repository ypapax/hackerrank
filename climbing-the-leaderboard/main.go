package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const workers = 1

type task struct {
	index      int
	aliceScore int32
}

// Complete the climbingLeaderboard function below.
func climbingLeaderboard(scores []int32, alice []int32) []int32 {
	t1 := time.Now()
	defer func() {
		log.Printf("climbingLeaderboard for scores len: %+v: %s", len(scores), time.Since(t1))
	}()
	ranks := getRanks(scores)
	var resultMtx sync.Mutex
	var result = make([]int32, len(alice))
	var tasks = make(chan task)
	wg := sync.WaitGroup{}
	wg.Add(len(alice))
	go func() {
		for i, a := range alice {
			tasks <- task{index: i, aliceScore: a}
		}
	}()
	for i := 0; i < workers; i++ {
		go func() {
			for {
				t, ok := <-tasks
				if !ok {
					return
				}
				r := getRank(scores, ranks, t.aliceScore)
				resultMtx.Lock()
				result[t.index] = r
				resultMtx.Unlock()
				wg.Done()
			}
		}()
	}
	wg.Wait()
	close(tasks)
	return result
}

func getRanks(scores []int32) []int32 {
	t1 := time.Now()
	defer func() {
		log.Printf("getRanks for scores len: %+v: %s", len(scores), time.Since(t1))
	}()
	if len(scores) == 0 {
		return nil
	}
	var rank = make([]int32, len(scores))
	rank[0] = 1
	for i, sc := range scores {
		if i == 0 {
			continue
		}
		if sc == scores[i-1] {
			rank[i] = rank[i-1]
			continue
		}
		rank[i] = rank[i-1] + 1
	}
	return rank
}

func getRank(scores, ranks []int32, aliceScore int32) int32 {
	t1 := time.Now()
	defer func() {
		log.Printf("getRank for scores len: %+v and aliceScore %+v: %s", len(scores), aliceScore, time.Since(t1))
	}()
	if aliceScore >= scores[0] {
		return 1
	}
	lastScore := scores[len(scores)-1]
	if lastScore > aliceScore {
		lastRank := ranks[len(ranks)-1]
		return lastRank + 1
	}
	leftIndex := binarySearch(scores, len(scores)-1, 0, aliceScore)
	if scores[leftIndex] == aliceScore {
		return ranks[leftIndex]
	}
	return ranks[leftIndex] + 1
}

func binarySearch(scores []int32, low, high int, target int32) int {
	if high > low {
		return -1
	}
	if scores[0] == target {
		return 0
	}
	last := scores[len(scores)-1]
	if last >= target {
		return len(scores) - 1
	}
	if scores[len(scores)-1] == target {
		return len(scores) - 1
	}
	mid := low + high/2
	if scores[mid] == target {
		return mid
	}
	if scores[mid-1] > target && target > scores[mid] {
		return mid - 1
	}
	if scores[mid] > target {
		high = mid + 1
		return binarySearch(scores, low, high, target)
	}
	low = mid - 1
	return binarySearch(scores, low, high, target)
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
