package solutions

func ProductExceptSelf(nums []int) []int {
    ans := make([]int, len(nums))
    product := getProduct(nums,-1)
    for idx, val := range nums{
        if val!=0{
            ans[idx] = product/val
        }else{
            p := getProduct(nums,idx)
            ans[idx] = p
        }
    }
    return ans
}

func getProduct(nums []int, ignore int) int{
    p := 1
    for idx, num := range nums{
        if ignore !=-1 && ignore==idx{
            continue
        }
        p*=num
    }
    return p
}