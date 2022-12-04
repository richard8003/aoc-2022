package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	stringData := loadData("data.txt")
	result := 0

	for _, v := range stringData {
		eachV := strings.Split(string(v), "\n")
		sum := 0
		for _, v := range eachV {

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
	fmt.Println(result)
}

func loadData(fileName string) []string {
	// read the file
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		os.Exit(1)
	}

	// slice of lines from the content of data.txt
	return strings.Split(string(file), "\n\n")
}
