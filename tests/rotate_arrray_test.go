package tests

import (
	"testing"

	"github.com/KellemNegasi/leetcode/solutions"
	"github.com/stretchr/testify/require"
)

func TestRotate(t *testing.T){
	testCases := []struct{
		input []int
		k int
		output []int
	}{
		{
			input: []int{1,2,3,4,5,6,7},
			k: 3,
			output: []int{5,6,7,1,2,3,4},
		},
		{
			input: []int{-1,-100,3,99},
			k:2,
			output: []int{3,99,-1,-100},
		},
		{
			input: []int{1,2,3},
			k: 0,
			output: []int{1,2,3},
		},
		{
			input: []int{1,2},
			k:2,
			output: []int{1,2},
		},
	}

	for _, tCase := range testCases{
		solutions.Rotate(tCase.input,tCase.k)
		require.Equal(t,tCase.output,tCase.input)
	}
}