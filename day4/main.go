package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

/*
2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
*/

//go:embed input
var input string

func main() {
	sum := 0
	foo := strings.Split(input, "\n")

	for _, row := range foo {
		if row != "" {

			if compare(row) {
				sum += 1
			}
		}
	}

	fmt.Println(sum)
}

func compare(foobar string) bool {

	splitFoobar := strings.Split(foobar, ",")
	foo := strings.Split(splitFoobar[0], "-")
	bar := strings.Split(splitFoobar[1], "-")

	foo0, _ := strconv.Atoi(foo[0])
	foo1, _ := strconv.Atoi(foo[1])
	bar0, _ := strconv.Atoi(bar[0])
	bar1, _ := strconv.Atoi(bar[1])

	//if foo0 <= bar0 && bar1 <= foo1 {
	if foo0 <= bar0 && foo1 >= bar1 || bar0 <= foo0 && bar1 >= foo1 {
		return true
	} else {
		return false
	}

}
