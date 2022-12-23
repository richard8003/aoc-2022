package main

import (
	_ "embed"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

/*
30373
25512
65332
33549
35390
*/

func main() {
	//trees := readFile("test")
	var trees [][]int
	trees = append(trees, []int{1, 1, 1, 1, 1})
	//trees = append(trees, []int{2, 5, 7, 1, 2})
	trees = append(trees, []int{1, 2, 3, 2, 1})
	trees = append(trees, []int{1, 1, 1, 1, 1})

	//for i, row := range trees {
	//// skip row 0 and the last row
	//if i == 0 || i == len(trees)-1 {
	//continue
	//}
	//// slice everything between the first and last elemet in the row
	//slice := row[1 : len(row)-1]
	//
	//foobar(row, slice)
	//}
	row := trees[1]
	slice := row[1 : len(row)-1]

	foobar(row, slice)
}

func foobar(row []int, slice []int) {
	for i := 1; i < len(slice)+1; i++ {
		//fmt.Println(row[i])
		before := row[:i]
		after := row[i+1:]
		current := row[i]

		fmt.Println("Checking row:", row)
		if checkSides(before, current) {
			if checkSides(after, current) {
				fmt.Println("All clear")
			}
		}
		fmt.Println("---------------------")
	}
}

func checkSides(slice []int, pivot int) bool {
	fmt.Println("Checking if either elemet en", slice, "is greater than", pivot)
	for _, s := range slice {
		if s >= pivot {
			fmt.Println(s, "is taller than", pivot)
			return false
		}
	}
	fmt.Println("Pivot", pivot, "can bee seen")
	return true
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
