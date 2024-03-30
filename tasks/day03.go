package tasks

import (
	"aoc23/util"
	"log"
	"strconv"
)

type coord struct {
	row int
	col int
}

func Day03Part1() int {
	lines, err := util.ReadFile("input/03.txt")
	// lines, err := util.ReadFile("input/03_example.txt")

	if err != nil {
		log.Fatal(err)
	}

	chars := buildMatrix(lines)

	return calculateAdjacents(chars)
}

func Day03Part2() int {
	lines, err := util.ReadFile("input/03.txt")
	// lines, err := util.ReadFile("input/03_example.txt")

	if err != nil {
		log.Fatal(err)
	}

	chars := buildMatrix(lines)

	return calcGearRatios(chars)
}

func buildMatrix(lines []string) [][]rune {
	chars := make([][]rune, 0)

	for _, line := range lines {
		chars = append(chars, []rune(line))
	}

	return chars
}

func calculateAdjacents(chars [][]rune) int {

	var sum int = 0

	var inNumber bool = false
	var startOfNumber int = -1
	var endOfNumber int = -1

	for i := 0; i < len(chars); i++ {
		for j := 0; j < len(chars[i]); j++ {
			// we're detecting the first digit of a number
			if !inNumber && runeIsDigit(chars[i][j]) {
				inNumber = true
				startOfNumber = j
				continue
			}

			// we're detecting that this rune is not part of the number anymore
			// so the rune before the current one is the last digit
			if inNumber && !runeIsDigit(chars[i][j]) {
				endOfNumber = j - 1

				if checkSurroundingOfNumber(chars, i, startOfNumber, endOfNumber) {
					x := getNumberFromChars(chars, i, startOfNumber, endOfNumber)
					// fmt.Println("adding", x)
					sum += x
				}

				inNumber = false
				startOfNumber = -1
				endOfNumber = -1
			}
		}

		// the number ends at the end of the line, so we need to check it
		if inNumber {
			endOfNumber = len(chars[i]) - 1

			if checkSurroundingOfNumber(chars, i, startOfNumber, endOfNumber) {
				x := getNumberFromChars(chars, i, startOfNumber, endOfNumber)
				// fmt.Println("adding", x)
				sum += x
			}

			inNumber = false
			startOfNumber = -1
			endOfNumber = -1
		}
	}

	return sum
}

func runeIsDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

/**
 * Check if the surrounding of a number contains a character other than a '.'
 */
func checkSurroundingOfNumber(chars [][]rune, line int, start int, end int) bool {
	// fmt.Println("checking surrounding of line:", line, "s: ", start, "e:", end)
	var topLeft coord
	var bottomRight coord

	topLeft = coord{line - 1, start - 1}
	bottomRight = coord{line + 1, end + 1}

	if topLeft.row < 0 {
		topLeft.row = 0
	}

	if topLeft.col < 0 {
		topLeft.col = 0
	}

	if bottomRight.row >= len(chars) {
		bottomRight.row = len(chars) - 1
	}

	if bottomRight.col >= len(chars[line]) {
		bottomRight.col = len(chars[line]) - 1
	}

	for i := topLeft.row; i <= bottomRight.row; i++ {
		for j := topLeft.col; j <= bottomRight.col; j++ {
			if isSpecialChar(chars[i][j]) {
				return true
			}
		}
	}

	return false
}

func getNumberFromChars(chars [][]rune, line int, start int, end int) int {
	str := ""

	for i := start; i <= end; i++ {
		str += string(chars[line][i])
	}

	ret, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}

	return ret
}

func isSpecialChar(r rune) bool {
	return !runeIsDigit(r) && r != '.'
}

func isGear(r rune) bool {
	return r == '*'
}

func calcGearRatios(chars [][]rune) int {
	var sum int = 0

	gearToNumberMap := make(map[string][]int)

	var inNumber bool = false
	var startOfNumber int = -1
	var endOfNumber int = -1

	for i := 0; i < len(chars); i++ {
		for j := 0; j < len(chars[i]); j++ {
			// we're detecting the first digit of a number
			if !inNumber && runeIsDigit(chars[i][j]) {
				inNumber = true
				startOfNumber = j
				continue
			}

			// we're detecting that this rune is not part of the number anymore
			// so the rune before the current one is the last digit
			if inNumber && !runeIsDigit(chars[i][j]) {
				endOfNumber = j - 1

				// todo
				number := getNumberFromChars(chars, i, startOfNumber, endOfNumber)
				gearKey := getGearKey(chars, i, startOfNumber, endOfNumber)

				if gearKey != "" {
					if len(gearToNumberMap[gearKey]) == 0 {
						gearToNumberMap[gearKey] = make([]int, 0)
					}
					gearToNumberMap[gearKey] = append(gearToNumberMap[gearKey], number)
				}

				inNumber = false
				startOfNumber = -1
				endOfNumber = -1
			}
		}

		// the number ends at the end of the line, so we need to check it
		if inNumber {
			endOfNumber = len(chars[i]) - 1

			// todo
			number := getNumberFromChars(chars, i, startOfNumber, endOfNumber)
			gearKey := getGearKey(chars, i, startOfNumber, endOfNumber)

			if gearKey != "" {
				if len(gearToNumberMap[gearKey]) == 0 {
					gearToNumberMap[gearKey] = make([]int, 0)
				}
				gearToNumberMap[gearKey] = append(gearToNumberMap[gearKey], number)
			}

			inNumber = false
			startOfNumber = -1
			endOfNumber = -1
		}
	}

	for _, entry := range gearToNumberMap {
		prod := 1

		if len(entry) < 2 {
			continue
		}

		for _, num := range entry {
			prod *= num
		}

		sum += prod
	}

	return sum
}

func getGearKey(chars [][]rune, line int, start int, end int) string {
	var topLeft coord
	var bottomRight coord

	topLeft = coord{line - 1, start - 1}
	bottomRight = coord{line + 1, end + 1}

	if topLeft.row < 0 {
		topLeft.row = 0
	}

	if topLeft.col < 0 {
		topLeft.col = 0
	}

	if bottomRight.row >= len(chars) {
		bottomRight.row = len(chars) - 1
	}

	if bottomRight.col >= len(chars[line]) {
		bottomRight.col = len(chars[line]) - 1
	}

	for i := topLeft.row; i <= bottomRight.row; i++ {
		for j := topLeft.col; j <= bottomRight.col; j++ {
			if isGear(chars[i][j]) {
				return string(i) + " - " + string(j)
			}
		}
	}

	return ""
}
