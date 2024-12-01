package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input1.txt
var input1 string

func part1() {
	left := []int{}
	right := []int{}

	var rows = strings.Split(input1, "\n")

	for i := 0; i < len(rows); i++ {
		var values = strings.Split(rows[i], "   ")
		if len(values) == 2 {
			first, err := strconv.Atoi(values[0])
			if err != nil {
				panic(err)
			}

			left = append(left, first)

			second, err := strconv.Atoi(values[1])
			if err != nil {
				panic(err)
			}

			right = append(right, second)
		}
	}

	slices.Sort(left)
	slices.Sort(right)

	sum := 0
	for i := 0; i < len(left); i++ {
		sum = sum + int(math.Abs(float64(left[i]-right[i])))
	}

	fmt.Println(sum)
}

func part2() {
	left := []int{}
	right := []int{}

	var rows = strings.Split(input1, "\n")

	for i := 0; i < len(rows); i++ {
		var values = strings.Split(rows[i], "   ")
		if len(values) == 2 {
			first, err := strconv.Atoi(values[0])
			if err != nil {
				panic(err)
			}

			left = append(left, first)

			second, err := strconv.Atoi(values[1])
			if err != nil {
				panic(err)
			}

			right = append(right, second)
		}
	}

	slices.Sort(left)
	slices.Sort(right)

	/*
	  Because the lists are sorted, we only need to check the right list for values which we havent seen before.
	  We can also break the inner loop if the value on the right array is greater then the value on the left.
	*/
	result := make([]int, len(left))
	rightStart := 0
	for i := 0; i < len(left); i++ {
		for a := rightStart; a < len(right); a++ {
			rightStart = a
			if right[a] > left[i] {
				break
			}

			if right[a] == left[i] {
				result[i] = result[i] + left[i]
			}
		}
	}

	sum := 0
	for _, num := range result {
		sum = sum + num
	}
	fmt.Println(sum)
}

func main() {
	fmt.Println("Part 1")
	part1()
	fmt.Println("Part 2")
	part2()
}
