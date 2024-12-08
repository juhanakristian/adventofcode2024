package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input1 string

var wordLen = 4

func countWords(lines []string, x int, y int) int {
	subLines := []string{}
	if y >= wordLen-1 {
		if len(lines[0])-x >= wordLen {
			subLine := string(lines[y][x]) + string(lines[y-1][x+1]) + string(lines[y-2][x+2]) + string(lines[y-3][x+3])
			subLines = append(subLines, subLine)
		}
		if x >= wordLen-1 {
			subLine := string(lines[y][x]) + string(lines[y-1][x-1]) + string(lines[y-2][x-2]) + string(lines[y-3][x-3])
			subLines = append(subLines, subLine)
		}

		subLine := ""
		for i := 0; i < wordLen; i++ {
			subLine += string(lines[y-i][x])
		}
		subLines = append(subLines, subLine)
	}

	if len(lines)-wordLen >= y {
		if len(lines[0])-x >= wordLen {
			subLine := string(lines[y][x]) + string(lines[y+1][x+1]) + string(lines[y+2][x+2]) + string(lines[y+3][x+3])
			subLines = append(subLines, subLine)
		}
		if x >= wordLen-1 {
			subLine := string(lines[y][x]) + string(lines[y+1][x-1]) + string(lines[y+2][x-2]) + string(lines[y+3][x-3])
			subLines = append(subLines, subLine)
		}

		subLine := ""
		for i := 0; i < wordLen; i++ {
			subLine += string(lines[y+i][x])
		}
		subLines = append(subLines, subLine)
	}

	line := lines[y]
	if len(line)-wordLen >= x {
		subLine := ""
		for i := 0; i < wordLen; i++ {
			subLine += string(line[x+i])
		}
		subLines = append(subLines, subLine)
	}

	if x >= wordLen-1 {
		subLine := ""
		for i := 0; i < wordLen; i++ {
			subLine += string(line[x-i])
		}
		subLines = append(subLines, subLine)
	}

	count := 0
	for _, subLine := range subLines {
		if subLine == "XMAS" {
			count++
		}
	}

	return count
}

func part1() {
	lines := strings.Split(input1, "\n")

	count := 0
	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			character := string(line[x])
			if character == "X" {
				count += countWords(lines, x, y)
			}
		}
	}

	fmt.Println(count)
}

func part2() {
	lines := strings.Split(input1, "\n")
	count := 0
	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			character := string(line[x])
			if character == "A" {
				if y >= 1 && y < len(lines)-1 {
					if x >= 1 && x < len(line)-1 {
						tlbr := string(lines[y-1][x-1]) + string(lines[y][x]) + string(lines[y+1][x+1])
						bltr := string(lines[y+1][x-1]) + string(lines[y][x]) + string(lines[y-1][x+1])

						if tlbr == "MAS" || tlbr == "SAM" {
							if bltr == "MAS" || bltr == "SAM" {
								count++
							}
						}
					}

				}

			}
		}
	}

	fmt.Println(count)
}

func main() {
	part1()
	part2()
}
