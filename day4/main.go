package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {
	data := parseData()

	fmt.Println(partOne(data))
	fmt.Println(two(data))

}

func two(data [][][]int) int {
	sum := 0

	for _, l := range data {
		left := l[0]
		right := l[1]

		if left[0] <= right[1] && left[1] >= right[0] {
			sum += 1
		}

	}
	return sum
}

func partOne(data [][][]int) int {
	sum := 0

	for _, l := range data {
		left := l[0]
		right := l[1]

		if left[0] >= right[0] && left[1] <= right[1] || right[0] >= left[0] && right[1] <= left[1] {
			sum += 1
		}

	}
	return sum
}

func parseData() [][][]int {
	raw := strings.TrimRight(input, "\n")
	//raw = strings.TrimRight(input, "\n")
	rows := strings.Split(raw, "\n")
	var data [][][]int

	for _, el := range rows {
		temp := splitSliceOfRows(el)
		data = append(data, temp)
	}
	return data
}

// used by parseData
func splitSliceOfRows(s string) [][]int {
	var row [][]int

	cellStr := strings.Split(s, ",")
	for _, c := range cellStr {
		cell := createNumberSlice(c)
		row = append(row, cell)
	}

	return row
}

// used by splitSliceOfRows > parseData
func createNumberSlice(s string) []int {
	var cell []int

	for _, v := range strings.Split(s, "-") {
		digit, _ := strconv.Atoi(string(v))
		cell = append(cell, digit)
	}

	return cell
}
