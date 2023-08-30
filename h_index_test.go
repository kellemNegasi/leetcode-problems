package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHIndex(t *testing.T){
	testCases := []struct{
		input []int
		output int
	}{
		{
			input: []int{3,0,6,1,5},
			output: 3,
		},
		{
			input: []int{1,3,1},
			output: 1,
		},
		{
			input: []int{1,2,2},
			output: 2,
		},
	}

	for _, tCase := range testCases{
		output := hIndex(tCase.input)
		require.Equal(t, tCase.output, output)
	}
}