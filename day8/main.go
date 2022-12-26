package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	//trees := readFile("onerow")
	trees := readFile("input")
	var total int

	for rowIndx, row := range trees {

		rowMax := -1
		sum := 0

		for treeIndx, tree := range row {
			// if the current row is the first or last row, automatically
			// sum ++
			if rowIndx == 0 {
				sum++
				fmt.Println(tree, "is visible from the top")
				continue
			}

			// chech it the current tree is visible from the left
			if visibleFromLeft(tree, rowMax) {
				sum++
				rowMax = setMax(tree, rowMax)
				fmt.Println(tree, "is visible from the left")
				continue
			}

			// chech it the current tree is visible from the right
			if treeIndx == len(row)-1 || visibleFromRight(tree, row[treeIndx+1:]) {
				sum++
				fmt.Println(tree, "is visible from the right")
				continue
			}

			if rowIndx == len(row)-1 {
				sum++
				fmt.Println(tree, "is visible from the bottom")
				continue
			} else {

				slice := trees[:rowIndx]

				for _, s := range slice {
					if s[treeIndx] > tree {
						fmt.Println(tree, "is blocked by", s[treeIndx])
						break
					}
				}

			}

		}
		fmt.Println("---------------------------")
		total += sum
	}
	fmt.Println(total, "visuble trees in total")

}

func visibleFromTop() bool {

	return true
}

func visibleFromRight(tree int, slice []int) bool {
	//fmt.Println(slice)
	for _, i := range slice {
		if tree <= i {
			return false
		}
	}
	return true
}

func visibleFromLeft(tree, rowMax int) bool {
	if tree > rowMax {
		return true
	}
	return false
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
