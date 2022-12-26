package main

import (
	_ "embed"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	trees := readFile("live")
	var total int

	for x, row := range trees {

		rowMax := -1
		sum := 0

		for i, tree := range row {
			fmt.Println(row)
			if tree > rowMax {
				sum++
				rowMax = setMax(tree, rowMax)
				fmt.Println(tree, "is visibles from the left")
				continue
			} else if checkFromRight(tree, row[i+1:]) {
				sum++
				fmt.Println(tree, "is visibles from the right")
				continue
			} else if checkTop(trees, tree, row, i, x) {
				sum++
				fmt.Println(tree, "is visible from the top")
				continue
			} else if checkBottom(trees, tree, row, i, x) {
				sum++
				fmt.Println(tree, "is visible from the bottom")
				continue
			} else {
				fmt.Println("tree is not visible")
			}
		}

		total += sum
	}
	fmt.Println(total)

}

func checkTop(trees [][]int, tree int, row []int, i int, x int) bool {
	slice := trees[:x]
	for _, t := range slice {

		if tree > t[i] {
			continue
		} else {
			return false
		}

	}

	return true
}

func checkBottom(trees [][]int, tree int, row []int, i int, x int) bool {
	slice := trees[x:]

	for _, t := range slice {

		if tree > t[i] {
			continue
		} else {
			return false
		}

	}

	return true
}

func checkFromRight(pivot int, slice []int) bool {
	for _, x := range slice {
		if slice == nil {
			return true

		}
		if x > pivot {
			return false
		}
	}

	return true
}

func setMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func readFile(f string) [][]int {
	file, _ := ioutil.ReadFile(f)
	data := strings.TrimRight(string(file), "\n")

	var foobar [][]int

	for _, i := range strings.Split(data, "\n") {
		var row []int

		for _, j := range i {
			num, _ := strconv.Atoi(string(j))
			row = append(row, num)
		}
		foobar = append(foobar, row)
	}

	return foobar

}
