package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input1.txt
var input1 string

func strArrToInt(str []string) []int {
	result := make([]int, len(str))
	for i, v := range str {
		j, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		result[i] = j
	}

	return result
}

func isValidStep(step int, increasing bool) bool {
	if increasing {
		return step > 0 && step < 4
	}

	return step < 0 && step > -4
}

func part1() {
	fmt.Println("PART 1")
	lines := strings.Split(input1, "\n")

	count := 0
	for _, l := range lines {
		levels := strArrToInt(strings.Split(l, " "))
		increasing := levels[0] < levels[len(levels)-1]
		validSteps := 0

		for i := 1; i < len(levels); i++ {
			step := levels[i] - levels[i-1]
			if !isValidStep(step, increasing) {
				break
			}

			validSteps++
		}

		if validSteps == len(levels)-1 {
			count += 1
		}
	}

	fmt.Println(count)
}

func part2() {
	fmt.Println("PART 2")
	lines := strings.Split(input1, "\n")

	check := [][]int{}
	for _, l := range lines {
		values := strArrToInt(strings.Split(l, " "))
		check = append(check, values)
	}

	count := 0

	for _, values := range check {
		for a, _ := range values {
			levels := append([]int{}, values[:a]...)
			levels = append(levels, values[a+1:]...)

			if levels[0] == levels[1] {
				continue
			}

			increasing := levels[0] < levels[len(levels)-1]
			validSteps := 0
			for i := 1; i < len(levels); i++ {
				step := levels[i] - levels[i-1]
				if !isValidStep(step, increasing) {
					break
				}

				validSteps++
			}

			if validSteps == len(levels)-1 {
				count += 1
				break
			}
		}
	}

	fmt.Println(count)
}

func main() {
	part1()
	part2()
}
