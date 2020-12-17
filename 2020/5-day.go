package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

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

	println("High Seat ID : " + strconv.Itoa(highSeatID))

	index := solveGetRowSeat(numberSeatUsed, 8)

	var sizeSeat = createSlice(128)
	column := solveGetColumnSeat(index, sizeSeat)

	solutionPart2 := index*8 + column
	println("Solution : " + strconv.Itoa(solutionPart2) + " ; row : " + strconv.Itoa(index) + " ; column : " + strconv.Itoa(column))

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
}

// Part 2 - Find Row's Seat
// NOTE : find column should be the same function and be used
// Maybe change loop to iterate over the 2nd and last -1 possibilities as mentioned in the instruction
func solveGetRowSeat(numberSeatUsed []int, min int) int {
	index := 0
	for i := range numberSeatUsed {
		val := numberSeatUsed[i]
		if val != 0 && val < min {
			min = val
			index = i
		}
	}
	return index
}

func solveGetColumnSeat(rowID int, sizeSeat []int) int {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rowString := ""

	// Find row which split the seats up to the value in args
	for scanner.Scan() {
		line := scanner.Text()
		rowInfo := line[0:7]
		rowSelected := binarySpacePartitioning(rowInfo, sizeSeat)
		if rowSelected == rowID {
			rowString = rowInfo
			break
		}
	}

	// Extract all row matching the pattern found in the previous loop
	ArrayRowsGrep := tinyGrep(rowString, file)
	column := solveFindLastSeat(ArrayRowsGrep)

	return column
}

func solveFindLastSeat(rows []string) int {
	columnUsed := createNullSlice(8)
	for i := 0; i < len(rows); i++ {
		columnInfo := rows[i][7:10]
		// Create a slice which will store all the value already used
		columnSeat := createSlice(8)
		columnUsed[binarySpacePartitioning(columnInfo, columnSeat)]++
	}

	min := 1
	for i := 0; i < len(columnUsed); i++ {
		min = minValue(columnUsed[i], min)
		if min == 0 {
			return i
		}
	}

	return -1
}

func tinyGrep(pattern string, fileArg io.Reader) []string {
	var patterns []string

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, pattern) {
			patterns = append(patterns, line)
		}
	}
	return patterns
}

func maxValue(valA int, valB int) int {
	if valA > valB {
		return valA
	}
	return valB
}

func minValue(valA int, valB int) int {
	if valA < valB {
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
