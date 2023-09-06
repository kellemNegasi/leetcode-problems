package tests

import (
	"testing"

	"github.com/KellemNegasi/leetcode/solutions"
	"github.com/stretchr/testify/require"
)

func TestMerge(t *testing.T){

	type testCase struct {
		nums1 []int
		m int
		nums2 []int
		n int
		expected []int
	}
	
	cases := []testCase{
		{
			nums1: []int{1,2,3,0,0,0},
			m: 3,
			nums2: []int{2,5,6},
			n:3,
			expected: []int{1,2,2,3,5,6},
		},{
			nums1: []int{1},
			m: 1,
			nums2: []int{},
			n:0,
			expected: []int{1},
		},{
			nums1: []int{0},
			m: 0,
			nums2: []int{1},
			n:1,
			expected: []int{1},
		},{
			nums1: []int{2,0},
			m: 1,
			nums2: []int{1},
			n: 1,
			expected: []int{1,2},
		},{
			nums1: []int{4,5,6,0,0,0},
			m:3,
			nums2: []int{1,2,3},
			n:3,
			expected: []int{1,2,3,4,5,6},
		},
	}

	for _, tCase := range cases{
		 solutions.Merge(tCase.nums1,tCase.m,tCase.nums2,tCase.n)
		require.Equal(t,tCase.expected,tCase.nums1)
	}
}