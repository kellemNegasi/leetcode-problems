package solutions

func Merge(nums1 []int, m int, nums2 []int, n int){
	if len(nums2)==0{
		return
	}

	if len(nums1)==0{
		nums1=nums2
		return
	}
	cur := 0
	l := len(nums1)
	j := 0
	for cur<l && j<n{
		v1 := nums1[cur]
		v2 := nums2[j]
		if v1>v2 && cur <m+j{
			nums1[cur]=nums2[j]
			j++
			for i:=cur+1;i<l;i++{
				next := nums1[i]
				nums1[i]=v1
				v1 = next
			}
		}else if cur>= m+j{
			nums1[cur] = v2
			j++
		}
		cur++
	}

}