package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// Solution : BFBFFFFLRL
func main() {
	start := time.Now()

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// part 1 - store high ID seat
	highSeatID := 0

	var numberSeatUsed = createNullSlice(128)
	var numberColumnUsed = createNullSlice(8)

	for scanner.Scan() {

		var sizeSeat = createSlice(128)
		var columnSeat = createSlice(8)

		// Get next line of puzzle
		line := scanner.Text()
		rowInfo := line[0:7]
		columnInfo := line[7:10]

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

		numberSeatUsed[sizeSeatBuff[0]]++
		numberColumnUsed[columnSeatBuff[0]]++
	}
	println(highSeatID)
	println(numberSeatUsed)

	min := 8
	index := 0
	for i := range numberSeatUsed {
		val := numberSeatUsed[i]
		if val != 0 && val < min {
			min = val
			index = i
		}
	}

	println("i : " + strconv.Itoa(index) + " ; value : " + strconv.Itoa(numberSeatUsed[index]))

	min = 1
	index = 0
	for i := range numberColumnUsed {
		val := numberColumnUsed[i]
		if val != 0 && val < min {
			min = val
			index = i
		}
	}

	println("i : " + strconv.Itoa(index) + " ; value : " + strconv.Itoa(numberColumnUsed[index]))

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

func fillSizeSeat(rowInfo, sizeSeat []int) int {
	var sizeSeatBuff []int
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
	return 0
}

func createSlice(size int) []int {
	var slice = make([]int, size)
	for i := range slice {
		slice[i] = i
	}

	return slice
}

func createNullSlice(size int) []int {
	var slice = make([]int, size)
	for i := range slice {
		slice[i] = 0
	}

	return slice
}
