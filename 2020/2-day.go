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

/* 	Instruction part 1 : Each line gives the password policy and then the password.
The password policy indicates the lowest and highest number of times a given letter must appear for the password to be valid.
For example, 1-3 a means that the password must contain a at least 1 time and at most 3 times.

	Insutrction part 2 : Each policy actually describes two positions in the password, where 1 means the first character, 2 means the second character, and so on.
(Be careful; Toboggan Corporate Policies have no concept of "index zero"!) Exactly one of these positions must contain the given letter. */

/* Main :
- create 2D slice
- launch parseSlice, which parse puzzle file into 2D slice : [X][password, policyCar, policyMin, policyMax]
- launch solve1, which access to 2D slice, count occurence of 'policyCar' in 'password',
and check if it is between 'policyMin' and 'policyMax'.
- launch solve2, which access to 2D slice and check if only one of posX and posY are equals to the policy.
*/
func main() {
	start := time.Now()

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	slicePuzzle := make([][]string, 0)
	slicePuzzle = parseFile(file, slicePuzzle)
	validPasswd := solve1v2(slicePuzzle)
	fmt.Println(validPasswd)
	validPasswd = solve2v1(slicePuzzle)
	fmt.Println(validPasswd)

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
}

/*	Input :
	- file => puzzle file
	- sliceOut => 2D slice representation of puzzle content
	Output :
	- sliceOut
	Info :
	sliceOut struct does not change between instructions of part 1 or 2.
	Only key change to be more accurate :
	- sliceOut (part 1) => [X][password, policyCar, policyMin, policyMax]
	- sliceOut (part 2) => [X][password, policyCar, posX, posY]
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
func solve1v1(slicePuzzle [][]string) int {
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
		if err != nil {
			log.Fatal(err)
		}

		if occurence >= min && occurence <= max {
			validPasswd++
		}

	}
	return validPasswd
}

// Same as solve1, with most beautiful code as 'for each' replace 3D access.
func solve1v2(slicePuzzle [][]string) int {
	validPasswd := 0

	for _, slice := range slicePuzzle {
		occurence := 0
		for i := 1; i < len(slice[0]); i++ {
			if string(slice[0][i]) == string(slice[1]) {
				occurence++
			}
		}

		// convert 3th and 4th elem of slice into integer to be compared with occurence
		min, _ := strconv.Atoi(slice[2])
		max, _ := strconv.Atoi(slice[3])

		if occurence >= min && occurence <= max {
			validPasswd++
		}
	}

	return validPasswd
}

/*	Input :
	- slicePuzzle == sliceOut
	Output :
	- validPassw => number of valid password according to the Instruction part 2
*/
func solve2v1(slicePuzzle [][]string) int {
	validPasswd := 0

	for _, slice := range slicePuzzle {
		// convert 3th and 4th elem of slice into integer to get char in string
		posX, _ := strconv.Atoi(slice[2])
		posY, _ := strconv.Atoi(slice[3])

		// check if only one of the both position contains the policy
		// NOTE : because of the blank at the begenning of the string, the position of the first car is 1.
		// Perfect according to the warning from the instruction part 2 :)
		if string(slice[0][posX]) == string(slice[1]) && string(slice[0][posY]) != string(slice[1]) {
			validPasswd++
		} else if string(slice[0][posX]) != string(slice[1]) && string(slice[0][posY]) == string(slice[1]) {
			validPasswd++
		} else {
			continue
		}
	}

	return validPasswd
}
