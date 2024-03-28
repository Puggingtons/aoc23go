package tasks

import (
	"aoc23/util"
	"errors"
	"log"
	"strings"
)

type fn func(string) (int, error)

func Day01Part1() int {
	lines, err := util.ReadFile("input/01.txt")

	if err != nil {
		log.Fatal(err)
	}

	return solve(lines, firstDigit, lastDigit)
}

func Day01Part2() int {
	lines, err := util.ReadFile("input/01.txt")

	if err != nil {
		log.Fatal(err)
	}

	return solve(lines, firstDigitOrNumberString, lastDigitOrNumberString)
}

func firstDigit(line string) (int, error) {
	for _, char := range line {
		if char >= '0' && char <= '9' {
			return int(char - '0'), nil
		}
	}

	return 0, errors.New("no digit found in line")
}

func lastDigit(line string) (int, error) {
	return firstDigit(util.Reverse(line))
}

func firstDigitByStringMatch(line string, substrs []string, vals []int) (int, error) {
	var ret int = 0
	var index int = len(line)

	for i, substr := range substrs {
		currentIndex := len(strings.Split(line, substr)[0])

		if currentIndex < index {
			index = currentIndex
			ret = vals[i]
		}
	}

	if ret == 0 {
		return 0, errors.New("no digit found in line")
	}

	return ret, nil
}

func firstDigitOrNumberString(line string) (int, error) {
	var substrs = []string{"1", "one", "2", "two", "3", "three", "4", "four", "5", "five", "6", "six", "7", "seven", "8", "eight", "9", "nine"}
	var vals = []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}

	return firstDigitByStringMatch(line, substrs, vals)
}

func lastDigitOrNumberString(line string) (int, error) {
	var substrs = []string{"1", "eno", "2", "owt", "3", "eerht", "4", "ruof", "5", "evif", "6", "xis", "7", "neves", "8", "thgie", "9", "enin"}
	var vals = []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}

	return firstDigitByStringMatch(util.Reverse(line), substrs, vals)
}

func solve(lines []string, firstNum fn, lastNum fn) int {
	firsts := make([]int, 0)
	lasts := make([]int, 0)

	for _, line := range lines {
		first, err := firstNum(line)
		if err != nil {
			log.Fatal(err)
		}

		firsts = append(firsts, first)

		last, err := lastNum(line)
		if err != nil {
			log.Fatal(err)
		}

		lasts = append(lasts, last)
	}

	sum := 0

	for index := range firsts {
		sum += firsts[index]*10 + lasts[index]
	}

	return sum
}
