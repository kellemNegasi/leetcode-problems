package solutions

import "sort"

func HIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i]>citations[j]
	})
	currentHIndex := 0

	for _, citationNum := range citations{
		if citationNum>currentHIndex{
			currentHIndex++
		}else{
			break
		}
	}
	
    return currentHIndex
}