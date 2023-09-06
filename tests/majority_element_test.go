package tests

import (
	"testing"

	"github.com/KellemNegasi/leetcode/solutions"
	"github.com/stretchr/testify/require"
)

func TestMajorityElement(t *testing.T) {
	testCases := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{3, 2, 3},
			output: 3,
		},
		{
			input:  []int{2,2,1,1,1,2,2},
			output: 2,
		},
		{
			input: []int{2,2,1,1,1,3,4},
			output: 1,
		},
	}
	for _, tCase := range testCases{
		answer :=  solutions.MajorityElement(tCase.input)
		require.Equal(t, tCase.output,answer)
	}
}
