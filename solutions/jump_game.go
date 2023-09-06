package solutions

import (
	"fmt"
	"math"
)

func CanJump(nums []int) bool {
	i := 0
	count := 0
	found := false
	for dest := 0; i < len(nums) && i <= dest; {
		prev := dest
		dest = max(dest, i+nums[i])
		if prev != dest {
			count++
		}
		if dest >= len(nums)-1 {
			found = true
			break
		}
		i++

	}
	if !found{
		count=0
	}
	fmt.Println(count)
	return i == len(nums)
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func CanJumpII(nums []int) int {
	jumpTable := make([]int, len(nums)-1)
	return computeJumpsDynamic(len(nums)-2, nums, jumpTable)
}

func computeJumps(startIndex int, nums []int) int {
	l := len(nums)
	if l == 1 {
		return 0
	}
	min := math.MaxInt
	for i := 1; i <= nums[startIndex]; i++ {
		if i+startIndex == l-1 {
			return 1
		} else if i+startIndex > l-1 {
			return -1
		} else {
			jumps := 1 + computeJumps(startIndex+i, nums)
			if jumps < min && jumps > 0 {
				min = jumps
			}
		}
	}
	return min
}

func computeJumpsDynamic(startIndex int, nums []int, table []int) int {
	l := len(nums)
	if l == 1 {
		return 0
	}

	minJump := math.MaxInt
	for i := 1; i <= nums[startIndex]; i++ {
		if i+startIndex == l-1 {
			minJump = 1
		} else if i+startIndex > l-1 {
			continue
		} else {
			if table[i+startIndex] == 0 {
				continue
			}
			jumps := 1 + table[i+startIndex]
			if jumps < minJump && jumps > 0 {
				minJump = jumps
			}
		}
	}

	table[startIndex] = minJump
	if startIndex == 0 {
		return minJump
	}
	computeJumpsDynamic(startIndex-1, nums, table)
	return table[0]
}
