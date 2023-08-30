package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProductExceptSelf(t *testing.T){
	testCases := []struct{
		input []int
		output []int
	}{
		{
			input: []int{-1,1,0,-3,3},
			output: []int{0,0,9,0,0},
		},
	}

	for _, tCase := range testCases{
		output := productExceptSelf(tCase.input)
		require.Equal(t, tCase.output, output)
	}
}