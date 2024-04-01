package day05

import "fmt"

type mapRange struct {
	destinationRangeStart int64
	sourceRangeStart      int64
	rangeLength           int64
}

func MakeMapRange(destinationRangeStart int64, sourceRangeStart int64, rangeLength int64) mapRange {
	ret := mapRange{destinationRangeStart, sourceRangeStart, rangeLength}
	return ret
}

func (r *mapRange) InSourceRange(val int64) bool {
	fmt.Println("testing in range: val:", val, "start:", r.sourceRangeStart, "end:", r.sourceRangeStart+r.rangeLength)
	isInRange := val >= r.sourceRangeStart && val < r.sourceRangeStart+r.rangeLength
	fmt.Println("testing in range:", isInRange)
	return isInRange
}

func (r *mapRange) MapToDestination(val int64) int64 {
	fmt.Println("mapping", val, "to", (val-r.sourceRangeStart)+r.destinationRangeStart)
	return (val - r.sourceRangeStart) + r.destinationRangeStart
}
