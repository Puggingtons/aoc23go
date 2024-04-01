package util

func InArray(haystack []int, needle int) bool {
	for _, hay := range haystack {
		if hay == needle {
			return true
		}
	}

	return false
}

func GetLast[T any](arr []T) T {
	return arr[len(arr)-1]
}
