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

/* Main :
- create 2D slice
- launch parseSlice, which parse puzzle file into 2D slice : [X][password, policyCar, policyMin, policyMax]
- launch solve1, which access to 2D slice, count occurence of 'policyCar' in 'password',
and check if it is between 'policyMin' and 'policyMax'.
*/
func main() {
	file, err := os.Open("/home/lab0-dev/Projets/AdventOfCode/2020/2-puzzle")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	slicePuzzle := make([][]string, 0)
	slicePuzzle = parseFile(file, slicePuzzle)
	validPasswd := solve2(slicePuzzle)
	fmt.Println(validPasswd)
}

/*	Input :
	- file => puzzle file
	- sliceOut => 2D slice representation of puzzle content
	Output :
	- sliceOut
	Info :
	- sliceOut => [X][password, policyCar, policyMin, policyMax]
*/
func parseFile(file io.Reader, sliceOut [][]string) [][]string {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Create a new slice to not erased futur data
		var policySlice = make([]string, 4)
		// Get next line of puzzle
		line := scanner.Text()
		/* Parse puzzle, ex : 6-10 p: ctpppjmdpppppp
		policySlice[0] = " ctpppjmdpppppp"
		policySlice[1] = "p"
		policySlice[2] = "6"
		policySlice[3] = "10"
		*/
		policySlice[0] = strings.Split(line, ":")[1]
		policy := strings.Split(line, ":")[0]
		policySlice[1] = strings.Split(policy, " ")[1]
		policyDigit := strings.Split(policy, " ")[0]
		policySlice[2] = strings.Split(policyDigit, "-")[0]
		policySlice[3] = strings.Split(policyDigit, "-")[1]
		sliceOut = append(sliceOut, policySlice)
	}
	return sliceOut
}

/*	Input :
	- slicePuzzle == sliceOut
	Output :
	- validPassw => number of valid password according to the Instruction part 1
*/
func solve1(slicePuzzle [][]string) int {
	validPasswd := 0
	for i := 0; i < len(slicePuzzle); i++ {
		occurence := 0
		// Get occurence of 'policyCar' in 'password'
		// NOTE : skip first car of password because is a space
		for j := 1; j < len(slicePuzzle[i][0]); j++ {
			if string(slicePuzzle[i][0][j]) == string(slicePuzzle[i][1]) {
				occurence++
			}
		}

		// convert 3th and 4th elem of slice into integer to be compared with occurence
		min, err := strconv.Atoi(slicePuzzle[i][2])
		if err != nil {
			log.Fatal(err)
		}
		max, err := strconv.Atoi(slicePuzzle[i][3])

		if occurence >= min && occurence <= max {
			validPasswd++
		}

	}
	return validPasswd
}

// Same as solve1, with most beautiful code as 'for each' replace 3D access.
func solve2(slicePuzzle [][]string) int {
	validPasswd := 0

	for _, slice := range slicePuzzle {
		occurence := 0
		for i := 1; i < len(slice[0]); i++ {
			if string(slice[0][i]) == string(slice[1]) {
				occurence++
			}
		}

		// convert 3th and 4th elem of slice into integer to be compared with occurence
		min, err := strconv.Atoi(slice[2])
		if err != nil {
			log.Fatal(err)
		}
		max, err := strconv.Atoi(slice[3])

		if occurence >= min && occurence <= max {
			validPasswd++
		}
	}

	return validPasswd
}
