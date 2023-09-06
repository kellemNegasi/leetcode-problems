package tests

import (
	"testing"

	"github.com/KellemNegasi/leetcode/solutions"
	"github.com/stretchr/testify/require"
)

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		input  []int
		output int
	}{
		{
			input: []int{7,6,4,3,1},
			output: 0,
		},
		{
			input: []int{7,1,5,3,6,4},
			output: 5,
		},
		{
			input:[]int {1,2},
			output: 1,
		},
		{
			input: []int{2,4,1},
			output: 2,
		},
		{
			input: []int{2, 1, 2, 0, 1},
			output: 1,
		},
		{
			input: []int{7,1,5,3,6,4},
			output: 5,
		},
	}

	for _, tCase := range testCases {
		output := solutions.MaxProfit(tCase.input)
		require.Equal(t, tCase.output, output)
	}
}
