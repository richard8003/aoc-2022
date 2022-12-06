package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed ex.txt
var ex string

/*
A X | rock     1p
B Y | paper    2p
C Z | scissors 3p

Y draw
X loose
Z win
*/

func main() {
	file, _ := os.Open("./ex.txt")

	defer file.Close()

	matrix := map[string]map[string]int{
		"A": {
			"X": 4,
			"Y": 8,
			"Z": 3,
		},
		"B": {
			"X": 1,
			"Y": 5,
			"Z": 9,
		},
		"C": {
			"X": 7,
			"Y": 2,
			"Z": 6,
		},
	}
	scanner := bufio.NewScanner(file)

	sumPoints := 0

	for scanner.Scan() {
		text := scanner.Text()
		splitText := strings.Split(text, " ")

		choice := getPartTwoChoice(splitText[0], splitText[1])
		sumPoints += matrix[splitText[0]][choice]
	}

	fmt.Println(sumPoints)
}
