package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type inputData [][]string

/*

A  | rock     1p | Y  paper  2p      win 6p
B  | paper    2p | X  rock  1p       loss 0p
C  | scissors 3p | Z  scissors 3p    draw 3p

*/

func main() {
	data := readData("input.txt")
	//data := readData("input.txt")
	printData(data)

	sum := 0

	for _, row := range data {
		me := row[1]
		opp := row[0]

		switch {
		case opp == "A":
			if me == "Y" {
				sum += 8
			}
			if me == "X" {
				sum += 4
			}
			if me == "Z" {
				sum += 3
			}
		case opp == "B":
			if me == "Y" {
				sum += 5
			}
			if me == "X" {
				sum += 1
			}
			if me == "Z" {
				sum += 9
			}
		case opp == "C":
			if me == "Y" {
				sum += 2
			}
			if me == "X" {
				sum += 7
			}
			if me == "Z" {
				sum += 6
			}

		}

	}

	fmt.Println(sum)
}

func printData(o [][]string) {
	for _, v := range o {
		fmt.Println(v)
	}
}

func readData(filename string) inputData {
	var data inputData

	stringData, _ := ioutil.ReadFile(filename)
	sliceData := strings.Split(string(stringData), "\n")

	for _, v := range sliceData {
		if v != "" {
			row := strings.Split(v, " ")
			data = append(data, row)
		}
	}
	return data
}
