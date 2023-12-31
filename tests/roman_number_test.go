package tests

import (
	"testing"

	"github.com/KellemNegasi/leetcode/solutions"
	"github.com/stretchr/testify/require"
)

func TestRomanToInt(t *testing.T) {
	testCases := []struct {
		input  string
		output int
	}{
		{
			input:  "III",
			output: 3,
		},
		{
			input:  "II",
			output: 2,
		},
		{
			input:  "VI",
			output: 6,
		},
		{
			input:  "XV",
			output: 15,
		},
		{
			input:  "IX",
			output: 9,
		},
	}

	for _, tCase := range testCases {
		output :=  solutions.RomanToInt(tCase.input)
		require.Equal(t, tCase.output, output)
	}
}

func TestIntToRoman(t *testing.T) {
	testCases := []struct {
		input  int
		output string 
	}{
		{
			input: 3,
			output:  "III",
		},
		{
			input: 2,
			output:  "II",
		},
		{
			input: 6,
			output:  "VI",
		},
		{
			input: 15,
			output:  "XV",
		},
		{
			input: 9,
			output:  "IX",
		},
		{
			input: 1994,
			output: "MCMXCIV",
		},
		{
			input: 10,
			output: "X",
		},
		{
			input: 20,
			output: "XX",
		},
		{
			input: 2032,
			output: "MMXXXII",
		},
	}

	for _, tCase := range testCases {
		output := solutions.IntToRoman(tCase.input)
		require.Equal(t, tCase.output, output)
	}
}
