package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {

	fmt.Printf("Task 1: %v\n", foobar()[0])
	fmt.Printf("Task 2: %v\n", foobar()[0]+foobar()[1]+foobar()[2])

}

func foobar() []int {
	var data []int

	for _, group := range strings.Split(input, "\n\n") {

		test := strings.Split(string(group), "\n")
		var groupSum int
		for _, t := range test {
			num, _ := strconv.Atoi(t)
			groupSum += num
		}
		data = append(data, groupSum)
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i] > data[j]
	})

	return data

}
