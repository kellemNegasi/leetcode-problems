package solutions

func RemoveElement(nums []int, val int) int {
	idx := 0
	curPos := 0
	removedElements := 0
	l := len(nums)
	for idx < l {
		if nums[idx] == val {
			idx++
			removedElements++
			continue
		}
		nums[curPos] = nums[idx]
		curPos++
		idx++
	}

	if idx != curPos {
		for curPos < l {
			nums[curPos] = 0
			curPos++
		}
	}
	return l - removedElements
}

func RemoveDuplicates(nums []int) int {
	idx := 1
	curPos := 1
	removedElements := 0
	l := len(nums)
	for idx < l {
		prev := nums[idx-1]
		cur := nums[idx]
		if cur == prev {
			idx++
			removedElements++
			continue
		}
		nums[curPos] = nums[idx]
		curPos++
		idx++
	}

	if idx != curPos {
		for curPos < l {
			nums[curPos] = 0
			curPos++
		}
	}
	return l - removedElements
}

func RemoveDuplicatesII(nums []int) int {
	idx := 1
	curPos := 1
	removedElements := 0
	l := len(nums)
	count := map[int]int{}
	for idx < l {
		prev := nums[idx-1]
		cur := nums[idx]
		if cur == prev {
			if count[cur] == 0 {
				count[cur]++
			} else {
				idx++
				removedElements++
				continue
			}

		}
		nums[curPos] = nums[idx]
		curPos++
		idx++
	}

	if idx != curPos {
		for curPos < l {
			nums[curPos] = 0
			curPos++
		}
	}
	return l - removedElements
}
