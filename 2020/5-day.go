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

	// part 1 - store high ID seat
	highSeatID := 0

	var numberSeatUsed = createNullSlice(128)
	var numberColumnUsed = createNullSlice(8)

	// Open and Read entire file
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		var sizeSeat = createSlice(128)
		var columnSeat = createSlice(8)

		// Get next line of puzzle
		line := scanner.Text()
		rowInfo := line[0:7]
		columnInfo := line[7:10]

		// select row and increment appropriate var
		rowSelected := binarySpacePartitioning(rowInfo, sizeSeat)
		numberSeatUsed[rowSelected]++

		// select column and increment appropriate var
		// NOTE : ALMOST WORKING, need to plug column into row table
		columnSelected := binarySpacePartitioning(columnInfo, columnSeat)
		numberColumnUsed[columnSelected]++

		// Get Seat ID from formula wrote in the instruction
		currentSeatID := rowSelected*8 + columnSelected

		// Part 1 - Get High Seat ID
		highSeatID = maxValue(currentSeatID, highSeatID)
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

func maxValue(valA int, valB int) int {
	if valA > valB {
		return valA
	}
	return valB
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

// Iteratre over a string representation of "binany space partitioning"
// and return the last element
func binarySpacePartitioning(str string, sizeSeat []int) int {
	var sizeSeatBuff []int
	for _, car := range str {
		s := string(car)
		if s == "B" || s == "R" {
			sizeSeatBuff = sliceArray(sizeSeat[:], s)
			sizeSeat = sizeSeatBuff
		} else if s == "F" || s == "L" {
			sizeSeatBuff = sliceArray(sizeSeat[:], s)
			sizeSeat = sizeSeatBuff
		}
	}
	return sizeSeatBuff[0]
}

// create a slice composed by n elem init to loop iteration
func createSlice(size int) []int {
	var slice = make([]int, size)
	for i := range slice {
		slice[i] = i
	}

	return slice
}

// create a slice composed by n elem init to 0
func createNullSlice(size int) []int {
	var slice = make([]int, size)
	for i := range slice {
		slice[i] = 0
	}

	return slice
}
