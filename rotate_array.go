package main

func rotate(nums []int, k int)  {
	k = k%len(nums)
	reverse(nums,0,len(nums)-1)
	reverse(nums,0,k-1)
	reverse(nums,k,len(nums)-1)
}

func reverse(arr []int, start, end int){
	for (start<end){
		val := arr[end]
		arr[end] = arr[start]
		arr[start] = val
		start++
		end--
	}
}