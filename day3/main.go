package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var input string

func main() {
	fmt.Println(taskOne())
}

func taskOne() string {
	rucksacks := strings.Split(input, "\n")
	valeuTable := makeValueTable()
	var total int

	for _, rucksack := range rucksacks {
		if rucksack != "" {
			total += valeuTable[checkHash(rucksack)]
		}
	}
	return fmt.Sprintf("Task 1: %v\n", total)
}

func checkHash(rucksack string) string {
	m := make(map[string]bool)

	for _, v := range rucksack[:len(rucksack)/2] {
		m[string(v)] = true
	}

	for _, r := range rucksack[len(rucksack)/2:] {
		if m[string(r)] {
			return string(r)
		}
	}
	return ""
}

func makeValueTable() map[string]int {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	m := make(map[string]int)

	for i, v := range chars {
		m[string(v)] = i + 1
	}

	return m
}
