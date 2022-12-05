package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed ex.txt
var ex string

/*
A  | rock     1p | Y  paper  2p      win 6p
B  | paper    2p | X  rock  1p       loss 0p
C  | scissors 3p | Z  scissors 3p    draw 3p

Y draw
X loose
Z win
*/

func main() {

	fmt.Println(ex)

	test2 := createList(ex)
	fmt.Println(test2)
	//fmt.Println(calcPointsOfList(test2))

}

//================================================================

func calcPointsOfList(s string) int {
	var total int

	for _, i := range strings.Split(string(ex), "\n") {

		if i != "" {

			you := string(i[0])
			me := string(i[2])

			switch {
			case you == "A":
				switch {
				case me == "Y":
					total += 8
				case me == "X":
					total += 4
				case me == "Z":
					total += 3
				}
			case you == "B":
				switch {
				case me == "Y":
					total += 5
				case me == "X":
					total += 1
				case me == "Z":
					total += 9
				}

			case you == "C":
				switch {
				case me == "Y":
					total += 2
				case me == "X":
					total += 7
				case me == "Z":
					total += 6
				}
			}
		}
	}

	return total

}

func createList(list string) string {
	var newList string

	for _, v := range strings.Split(ex, "\n") {

		if v != "" {
			one := string(v[0])
			two := string(v[2])

			var add string

			if two == "Y" {
				switch {
				case one == "A":
					add = "A X"
				case one == "B":
					add = "B Y"
				case one == "C":
					add = "C Z"
				}
			}
			if two == "X" {
				switch {
				case one == "A":
					add = "A Z"
				case one == "B":
					add = "B X"
				case one == "C":
					add = "C Y"
				}
			}
			if two == "Z" {
				switch {
				case one == "A":
					add = "A Y"
				case one == "B":
					add = "B Z"
				case one == "C":
					add = "C X"
				}
			}

			fmt.Println(add)
			newList += add
		}
	}

	return newList

}
