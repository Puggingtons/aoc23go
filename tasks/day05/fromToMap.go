package day05

type fromToMap struct {
	from      string
	to        string
	mapRanges []mapRange
}

func MakeFromToMap(from string, to string) *fromToMap {
	return &fromToMap{from: from, to: to}
}

func (m *fromToMap) AddRange(mapRange mapRange) {
	m.mapRanges = append(m.mapRanges, mapRange)
}

func (m *fromToMap) MapValue(val int64) int64 {
	for _, mapRange := range m.mapRanges {
		if mapRange.InSourceRange(val) {
			return mapRange.MapToDestination(val)
		}
	}

	return val
}
