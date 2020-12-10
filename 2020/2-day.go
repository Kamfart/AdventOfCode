package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("/home/lab0-dev/Projets/AdventOfCode/2020/2-puzzle")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// [password, policyCar, policyMin, policyMax]
		var policySlice = make([]string, 4)
		policySlice[0] = strings.Split(line, ":")[1]
		policy := strings.Split(line, ":")[0]
		policySlice[1] = strings.Split(policy, " ")[1]
		policyDigit := strings.Split(policy, " ")[0]
		policySlice[2] = strings.Split(policyDigit, "-")[0]
		policySlice[3] = strings.Split(policyDigit, "-")[1]

		fmt.Println(policy)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
