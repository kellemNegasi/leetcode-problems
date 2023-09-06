package tests

import (
	"testing"

	"github.com/KellemNegasi/leetcode/solutions"
	"github.com/stretchr/testify/require"
)

func TestLengthOfLastWord(t *testing.T) {
	testCases := []struct {
		input  string 
		output int
	}{
		{
			input:  "a",
			output: 1,
		},
	}

	for _, tCase := range testCases {
		output :=  solutions.LengthOfLastWord(tCase.input)
		require.Equal(t, tCase.output, output)
	}
}