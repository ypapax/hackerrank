package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"

	"github.com/mattn/go-tty"
)

func debugArrange(m [][]int) [][]int {
	// debugging

	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
		return m
	}
	defer tty.Close()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			log.Println("closing", sig)
			tty.Close()
			os.Exit(1)
		}
	}()

	scanInt := func(msg string) (int, error) {
		fmt.Fprintln(os.Stderr, msg)
		s, err := tty.ReadString()
		if err != nil {
			log.Println(err)
			return 0, err
		}
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			log.Println(err)
			return 0, err
		}
		fmt.Fprintln(os.Stderr, "you choosed", i)
		return int(i), nil
	}

	for {
		boxFrom, err := scanInt("what box")
		if err != nil {
			log.Println(err)
			continue
		}
		ballNumber, err := scanInt("ball type number")
		if err != nil {
			log.Println(err)
			continue
		}
		boxTo, err := scanInt("where to move")
		if err != nil {
			log.Println(err)
			continue
		}
		ballNumberInReturn, err := scanInt("ball type number in return")
		if err != nil {
			log.Println(err)
			continue
		}

		swapByBoxFromToAndBallNumber(boxFrom, ballNumber, boxTo, ballNumberInReturn, m)
	}
	return m
}
