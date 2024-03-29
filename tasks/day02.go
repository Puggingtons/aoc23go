package tasks

import (
	"aoc23/util"
	"log"
	"strconv"
	"strings"
)

const MAX_RED int = 12
const MAX_GREEN int = 13
const MAX_BLUE int = 14

func Day02Part1() int {
	lines, err := util.ReadFile("input/02.txt")

	var ret int = 0

	if err != nil {
		log.Fatal(err)
	}

	for i, line := range lines {
		if checkGame(line) {
			ret += i + 1
		}
	}

	return ret
}

func Day02Part2() int {
	lines, err := util.ReadFile("input/02.txt")

	var ret int = 0

	if err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		ret += power(line)
	}

	return ret
}

func checkGame(line string) bool {
	pulls := strings.Split(strings.Split(line, ":")[1], ";")

	for _, pull := range pulls {
		colorPulls := strings.Split(pull, ",")

		for _, colorPull := range colorPulls {
			amount, err := strconv.Atoi(strings.Split(colorPull, " ")[1])

			if err != nil {
				for _, colorPull := range strings.Split(colorPull, " ") {
					log.Println("#", colorPull, "#")
				}
				log.Fatal(err)
			}

			if strings.Contains(colorPull, "red") {
				if amount > MAX_RED {
					return false
				}
			} else if strings.Contains(colorPull, "green") {
				if amount > MAX_GREEN {
					return false
				}
			} else if strings.Contains(colorPull, "blue") {
				if amount > MAX_BLUE {
					return false
				}
			}
		}
	}

	return true
}

func power(line string) int {
	var minRed int = 0
	var minGreen int = 0
	var minBlue int = 0

	pulls := strings.Split(strings.Split(line, ":")[1], ";")

	for _, pull := range pulls {
		colorPulls := strings.Split(pull, ",")

		for _, colorPull := range colorPulls {
			amount, err := strconv.Atoi(strings.Split(colorPull, " ")[1])

			if err != nil {
				for _, colorPull := range strings.Split(colorPull, " ") {
					log.Println("#", colorPull, "#")
				}
				log.Fatal(err)
			}

			if strings.Contains(colorPull, "red") {
				if amount > minRed {
					minRed = amount
				}
			} else if strings.Contains(colorPull, "green") {
				if amount > minGreen {
					minGreen = amount
				}
			} else if strings.Contains(colorPull, "blue") {
				if amount > minBlue {
					minBlue = amount
				}
			}
		}
	}

	return minRed * minGreen * minBlue
}
