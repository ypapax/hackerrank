package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"log"
)



func main() {
	reader := bufio.NewReader(os.Stdin)
	line1, err := reader.ReadString('\n')
	if err != nil {
		errPrint(err)
	}
	line2, err := reader.ReadString('\n')
	if err != nil {
		errPrint(err)
	}
	arr1 := strings.Split(strings.TrimSpace(line1), " ")
	if len(arr1) != 2 {
		errPrint("line 1 is too short")
	}


}

func errPrint(err error) {
	log.Println(err)
	os.Exit(2)
}
