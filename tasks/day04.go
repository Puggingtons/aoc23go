package tasks

import (
	"aoc23/util"
	"log"
	"strconv"
	"strings"
)

func Day04Part1() int {
	lines, err := util.ReadFile("input/04.txt")
	// lines, err := util.ReadFile("input/04_example.txt")

	if err != nil {
		log.Fatal(err)
	}

	return calcPoints(lines)
}

func Day04Part2() int {
	lines, err := util.ReadFile("input/04.txt")
	// lines, err := util.ReadFile("input/04_example.txt")

	if err != nil {
		log.Fatal(err)
	}

	return calcNumberOfTickets(lines)
}

func calcPoints(lines []string) int {
	sum := 0

	for _, line := range lines {
		res, _ := scoreScratchcard(line)
		// res, num := scoreScratchcard(line)
		// fmt.Println(line, "<-->", res, "|", num)
		sum += res
	}

	return sum
}

func calcNumberOfTickets(lines []string) int {
	amounts := make([]int, len(lines))

	for i := 0; i < len(amounts); i++ {
		amounts[i] = 1
	}

	for i, line := range lines {
		_, num := scoreScratchcard(line)

		for j := 1; j <= num; j++ {
			amounts[i+j] += amounts[i]
		}
	}

	sum := 0

	for i := 0; i < len(amounts); i++ {
		sum += amounts[i]
	}

	return sum
}

func scoreScratchcard(line string) (int, int) {
	numbersString := strings.Split(line, ":")[1]

	sections := strings.Split(numbersString, " | ")

	winners := sections[0]
	pulls := sections[1]

	winningNumbers := generateWinningNumbers(winners)

	return calculateScore(winningNumbers, pulls)
}

func generateWinningNumbers(winners string) []int {
	winningNumbers := make([]int, 0)

	for _, split := range strings.Split(winners, " ") {
		if split == "" {
			continue
		}

		num, err := strconv.Atoi(split)

		if err != nil {
			log.Fatal(err)
		}

		winningNumbers = append(winningNumbers, num)
	}

	return winningNumbers
}

func calculateScore(winningNumbers []int, pulls string) (int, int) {
	score := 0
	amount := 0

	for _, split := range strings.Split(pulls, " ") {
		if split == "" {
			continue
		}

		num, err := strconv.Atoi(split)

		if err != nil {
			log.Fatal(err)
		}

		if util.InArray(winningNumbers, num) {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
			amount++
		}
	}

	return score, amount
}
