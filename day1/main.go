package main

import (
	_ "embed"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	data := creataArryOfValues()

	fmt.Printf("Task 1: %v\n", data[0])
	fmt.Printf("Task 2: %v\n", data[0]+data[1]+data[2])

}

func creataArryOfValues() []int {
	var data []int

	strData := strings.Split(input, "\n\n")

	for _, set := range strData {
		sumOfSet := getSumOfSet(set)
		data = append(data, sumOfSet)
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i] > data[j]
	})

	return data
}

// used by creataArryOfValues
func getSumOfSet(set string) int {
	var sum int

	slicedSet := strings.Split(string(set), "\n")
	for _, el := range slicedSet {

		if el != "" {

			elInt, err := strconv.Atoi(el)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			sum += elInt
		}
	}

	return sum
}
