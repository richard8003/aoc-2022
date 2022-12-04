package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

/*
A  | rock     1p | Y  paper  2p      win 6p
B  | paper    2p | X  rock  1p       loss 0p
C  | scissors 3p | Z  scissors 3p    draw 3p
*/
type tableOfScores [][]string

func main() {
	data := makeInputSlice()
	sum := calcSumOfAllScores(data)
	fmt.Println(sum)
}

func calcSumOfAllScores(t tableOfScores) int {
	var sum int
	for _, row := range t {
		foo := row[1]
		bar := row[0]

		switch {
		case bar == "A":
			switch {
			case foo == "Y":
				sum += 8
			case foo == "X":
				sum += 4
			case foo == "Z":
				sum += 3
			}
		case bar == "B":
			switch {
			case foo == "Y":
				sum += 5
			case foo == "X":
				sum += 1
			case foo == "Z":
				sum += 9
			}

		case bar == "C":
			switch {
			case foo == "Y":
				sum += 2
			case foo == "X":
				sum += 7
			case foo == "Z":
				sum += 6
			}
		}
	}

	return sum
}

func makeInputSlice() tableOfScores {
	var data tableOfScores
	sliceData := strings.Split(string(input), "\n")

	for _, v := range sliceData {
		if v != "" {
			row := strings.Split(v, " ")
			data = append(data, row)
		}
	}

	return data
}
