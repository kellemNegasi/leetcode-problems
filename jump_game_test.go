package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanJump(t *testing.T){
	testCases := []struct{
		input []int
		output bool
	}{
		{
			input: []int{2,3,1,1,4},
			output: true,
		},
		{
			input: []int{3,2,1,1,4},
			output: false,
		},
		{
			input: []int{2,0},
			output: true,
		},
		{
			input: []int{1,2},
			output: true,
		},
		{
			input: []int{0},
			output: true,
		},
		{
			input: []int{5},
			output: true,
		},
		{
			input: []int{0,1},
			output: false,
		},
	}

	for _, tCase := range testCases{
		output := canJump(tCase.input)
		_ = output
		require.Equal(t, tCase.output, output)
	}
}

func TestCanJumpII(t *testing.T){
	testCases := []struct{
		input []int
		output int
	}{
		{
			input: []int{2,3,1,1,4},
			output: 2,
		},
		{
			input: []int{2,0},
			output: 1,
		},
		{
			input: []int{1,2},
			output: 1,
		},
		{
			input: []int{0},
			output: 0,
		},
		{
			input: []int{5},
			output: 0,
		},
		{
			input: []int{1, 3, 5, 8, 9, 2, 6, 7, 6, 8, 9},
			output: 3,
		},
		{
			input: []int{1,1,1,1},
			output: 3,
		},
	}

	for _, tCase := range testCases{
		output := canJumpII(tCase.input)
		require.Equal(t, tCase.output, output)
	}
}