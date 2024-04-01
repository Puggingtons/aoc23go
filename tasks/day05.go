package tasks

import (
	"aoc23/tasks/day05"
	"aoc23/util"
	"log"
)

func Day05Part1() int64 {
	lines, err := util.ReadFile("input/05.txt")
	// lines, err := util.ReadFile("input/05_example.txt")

	if err != nil {
		log.Fatal(err)
	}

	solver := day05.MakePart1Solver(lines)

	return solver.Solve()
}
