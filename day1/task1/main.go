package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {

	strData := strings.Split(input, "\n\n")

	for _, set := range strData {
		foobar := strings.Split(string(set), "\n")
		//fmt.Println(string(set))
		for _, row := range foobar {
			fmt.Println(row)
			fmt.Println("(((((((((())))))))")
		}
	}
}

func foobar() []int {

	var arrData []int

	return arrData
}

func taskOne(data []string) int {
	var result int

	for _, values := range data {
		val := strings.Split(string(values), "\n")
		sum := 0
		for _, v := range val {

			if v != "" {
				value, err := strconv.Atoi(v)
				if err != nil {
					fmt.Println(err.Error())
					os.Exit(1)
				}
				sum += value
			}
		}
		if sum > result {
			result = sum
		}

	}
	return result
}
