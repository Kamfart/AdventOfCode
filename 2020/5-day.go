package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// size := 0
	highSeatID := 0
	for scanner.Scan() {

		var sizeSeat = make([]int, 128)
		var columnSeat = make([]int, 8)

		for i := range sizeSeat {
			sizeSeat[i] = i
		}

		for i := range columnSeat {
			columnSeat[i] = i
		}

		// Get next line of puzzle
		line := scanner.Text()
		rowInfo := line[0:7]
		columnInfo := line[7:10]

		// println(rowLine, seatLine)

		var sizeSeatBuff []int
		var columnSeatBuff []int
		for _, car := range rowInfo {
			s := string(car)
			if s == "B" {
				sizeSeatBuff = sliceArray(sizeSeat[:], "B")
				sizeSeat = sizeSeatBuff
			} else if s == "F" {
				sizeSeatBuff = sliceArray(sizeSeat[:], "F")
				sizeSeat = sizeSeatBuff
			}

		}

		for _, car := range columnInfo {
			s := string(car)
			if s == "L" {
				columnSeatBuff = sliceArray(columnSeat[:], "L")
				columnSeat = columnSeatBuff
			} else if s == "R" {
				columnSeatBuff = sliceArray(columnSeat[:], "R")
				columnSeat = columnSeatBuff
			}

		}

		currentSeatID := sizeSeatBuff[0]*8 + columnSeatBuff[0]
		if currentSeatID > highSeatID {
			highSeatID = currentSeatID
		}
	}
	println(highSeatID)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
}

// Split array and return the "back" or "front" part
func sliceArray(array []int, part string) []int {
	if part == "B" || part == "R" {
		return array[len(array)/2 : len(array)]
	} else if part == "F" || part == "L" {
		return array[0 : len(array)/2]
	} else {
		return nil
	}
}
