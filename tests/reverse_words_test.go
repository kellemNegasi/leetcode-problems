package tests

import (
	"testing"

	"github.com/KellemNegasi/leetcode/solutions"
	"github.com/stretchr/testify/require"
)

func TestReverseWords(t *testing.T) {
	type testCase struct {
		input    string
		output   string
	}
	cases := []testCase{
		{
			input:    "the sky is blue",
			output:   "blue is sky the",
		},
		{
			input:    "  hello world  ",
			output:   "world hello",
		},
		{
			input:    "a good   example",
			output:  "example good a",
		},
	}

	for _, tCase := range cases {
		actual :=  solutions.ReverseWords(tCase.input)
		require.Equal(t, tCase.output, actual)
	}

}