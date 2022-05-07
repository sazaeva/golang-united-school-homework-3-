package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	key := make([]int, 0, len(input))
	for val := range input {
		key = append(key, val)
	}
	sort.Ints(key)

	return result
}
