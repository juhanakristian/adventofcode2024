package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func intOrPanic(value string) int {
	result, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	return result
}

//go:embed input.txt
var input1 string

func part1() {
	re := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")
	matches := re.FindAllStringSubmatch(input1, -1)

	sum := 0
	for _, match := range matches {
		fmt.Println(match[0])
		first := intOrPanic(match[1])
		second := intOrPanic(match[2])
		sum += first * second
	}

	fmt.Println(sum)
}

func biggestSmallerThen(values []int, valueToCompare int) int {
	previousValue := values[0]

	for i := 1; i < len(values); i++ {
		value := values[i]
		if value > valueToCompare {
			break
		}

		previousValue = value
	}

	if previousValue > valueToCompare {
		return -1
	}

	return previousValue
}

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

func part2() {
	mul := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")
	do := regexp.MustCompile("do\\(\\)")
	dont := regexp.MustCompile("don't\\(\\)")

	allDos := do.FindAllStringIndex(input1, -1)
	allDonts := dont.FindAllStringIndex(input1, -1)
	allMuls := mul.FindAllStringIndex(input1, -1)
	allDos = append([][]int{{0, 4}}, allDos...)

	fmt.Println(len(allDos), len(allDonts), len(allMuls))

	dos := []int{}
	for _, d := range allDos {
		dos = append(dos, d[0])
	}

	donts := []int{}
	for _, d := range allDonts {
		donts = append(donts, d[0])
	}

	sum := 0
	for _, mulIndex := range allMuls {
		previousDo := biggestSmallerThen(dos, mulIndex[0])
		previousDont := biggestSmallerThen(donts, mulIndex[0])

		if previousDo > previousDont {
			mulString := input1[mulIndex[0]:mulIndex[1]]
			fmt.Println(mulString, mulIndex[0], previousDo, previousDont)
			mulString = strings.Replace(mulString, "mul(", "", 1)
			mulString = strings.Replace(mulString, ")", "", 1)
			values := strings.Split(mulString, ",")
			intValues := strArrToInt(values)
			sum += intValues[0] * intValues[1]
		}
	}

	fmt.Println(sum)
}

func main() {
	part2()
}
