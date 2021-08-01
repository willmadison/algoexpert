package algoexpert

import "sort"

func IsValidSubsequence(array, sequence []int) bool {
	indiciesByValue := map[int][]int{}

	for i, v := range array {
		indiciesByValue[v] = append(indiciesByValue[v], i)
	}

	var indicies []int

	for _, v := range sequence {
		if indexes, present := indiciesByValue[v]; !present || len(indexes) == 0 {
			return false
		} else {
			indicies = append(indicies, indexes[0])
			indexes = indexes[1:]
			indiciesByValue[v] = indexes
		}
	}

	return sort.IntsAreSorted(indicies)
}
