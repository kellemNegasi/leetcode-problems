package main

func majorityElement(nums []int) int {
    count := 0
	majElement := nums[0]
	for _, num := range nums{
		if count==0{
			majElement = num
		}

		if num==majElement{
			count++
		}else{
			count--
		}
	}
	return majElement
}


