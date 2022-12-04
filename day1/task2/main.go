package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	stringData := loadData("example.txt")
	rows := []int{}

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
		rows = append(rows, sum)
	}

	sort.Slice(rows, func(i, j int) bool {
		return rows[i] > rows[j]
	})

	result := rows[0] + rows[1] + rows[2]
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
