package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveElement(t *testing.T) {
	type testCase struct {
		input    []int
		val      int
		output   int
		expected []int
	}
	cases := []testCase{
		{
			input:    []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			output:   5,
			expected: []int{0, 1, 3, 0, 4, 0, 0, 0},
		},
		{
			input:    []int{3, 2, 2, 3},
			val:      3,
			output:   2,
			expected: []int{2, 2, 0, 0},
		},
	}

	for _, tCase := range cases {
		actual := removeElement(tCase.input, tCase.val)
		require.Equal(t, tCase.output, actual)
		require.Equal(t, tCase.expected, tCase.input)
	}

}

func TestRemoveDuplicates(t *testing.T) {
	type testCase struct {
		input    []int
		output   int
		expected []int
	}
	cases := []testCase{
		{
			input:    []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			output:   5,
			expected: []int{0, 1, 2, 3, 4, 0, 0, 0, 0, 0},
		},
		{
			input:    []int{3, 2, 2, 3},
			output:   2,
			expected: []int{2, 2, 0, 0},
		},
	}

	for _, tCase := range cases {
		actual := removeDuplicates(tCase.input)
		require.Equal(t, tCase.output, actual)
		require.Equal(t, tCase.expected, tCase.input)
	}
}

func TestRemoveDuplicatesII(t *testing.T){
	type testCase struct {
		input    []int
		output   int
		expected []int
	}
	cases := []testCase{
		{
			input:    []int{1,1,1,2,2,3},
			output:   5,
			expected: []int{1,1,2,2,3,0},
		},
		{
			input:    []int{0,0,1,1,1,1,2,3,3},
			output:   7,
			expected: []int{0,0,1,1,2,3,3,0,0},
		},
	}

	for _, tCase := range cases {
		actual := removeDuplicatesII(tCase.input)
		require.Equal(t, tCase.output, actual)
		require.Equal(t, tCase.expected, tCase.input)
	}
}
