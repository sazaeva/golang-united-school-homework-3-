package homework

import (
	"fmt"
	"sort"
)

func SortMapValues(input map[int]string) (result []string) {
	keys := make([]int, 0, len(input))
	for k := range input {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Printf("%s ", input[k])

	}
	return
}
