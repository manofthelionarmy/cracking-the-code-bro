package ch8

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRobotInAGrid(t *testing.T) {
	matrix := [][]int{
		{0, 0, 0, 1},
		{0, 1, 0, 0},
		{0, 1, 1, 0},
		{1, 0, 0, 0},
	}

	require.Equal(t, true, RobotInAGrid(matrix))

	matrix = [][]int{
		{0, 0, 0, 1},
		{0, 1, 0, 0},
		{0, 1, 1, 1},
		{1, 0, 0, 0},
	}
	require.Equal(t, false, RobotInAGrid(matrix))

	matrix = [][]int{
		{0, 1, 1, 1},
		{0, 0, 0, 1},
		{1, 1, 0, 0},
		{1, 0, 0, 0},
	}
	require.Equal(t, true, RobotInAGrid(matrix))

	matrix = [][]int{
		{0, 1, 1, 1},
		{0, 1, 0, 1},
		{1, 1, 1, 0},
		{0, 0, 0, 0},
	}
	require.Equal(t, false, RobotInAGrid(matrix))
}
