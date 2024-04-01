package day05

import (
	"aoc23/util"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type part1Solver struct {
	seeds      []int64
	fromToMaps []*fromToMap
}

func (s *part1Solver) Solve() int64 {
	lowestSolution := int64(math.MaxInt64)

	for _, seed := range s.seeds {
		fmt.Println("seed:", seed)
		location := s.findLocationOfSeed(seed)
		fmt.Println("location", location)

		if location < lowestSolution {
			lowestSolution = location
		}
	}

	return lowestSolution
}

func (s *part1Solver) findLocationOfSeed(seed int64) int64 {
	currentStage := "seed"
	currentValue := seed

	for currentStage != "location" {
		fromToMap := s.getMapBySource(currentStage)
		fmt.Println(fromToMap.from, ">", fromToMap.to)

		currentValue = fromToMap.MapValue(currentValue)
		currentStage = fromToMap.to
	}

	return currentValue
}

func (s *part1Solver) getMapBySource(source string) *fromToMap {
	for _, fromToMap := range s.fromToMaps {
		if fromToMap.from == source {
			return fromToMap
		}
	}

	panic("no map with name '" + source + "' found")
}

func MakePart1Solver(lines []string) part1Solver {
	ret := part1Solver{}
	ret.fromToMaps = make([]*fromToMap, 0)

	for _, line := range lines {
		if strings.Contains(line, "seeds:") {
			ret.seeds = parseSeeds(line)

		} else if strings.Contains(line, "map:") {
			ret.fromToMaps = append(ret.fromToMaps, parseNewMap(line))

		} else if len(strings.Split(line, " ")) == 3 {
			latestMap := util.GetLast(ret.fromToMaps)
			latestMap.AddRange(parseRange(line))
		}
	}

	return ret
}

func parseSeeds(line string) []int64 {
	seedsSeciton := strings.Replace(line, "seeds: ", "", 1)
	seedstrings := strings.Split(seedsSeciton, " ")

	ret := make([]int64, len(seedstrings))

	for i, seedstring := range seedstrings {
		ret[i] = parseNum(seedstring)
	}

	return ret
}

func parseNewMap(line string) *fromToMap {
	nameSection := strings.Split(line, " ")[0]

	names := strings.Split(nameSection, "-")

	return MakeFromToMap(names[0], names[2])
}

func parseRange(line string) mapRange {
	nums := strings.Split(line, " ")

	return MakeMapRange(parseNum(nums[0]), parseNum(nums[1]), parseNum(nums[2]))
}

func parseNum(str string) int64 {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	return num
}
