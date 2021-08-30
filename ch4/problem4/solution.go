package problem4

import (
	"math"
)

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func abs(val int) int {
	return int(math.Abs(float64(val)))
}

func isValid(root *Node) int {
	if root == nil {
		return -1
	}

	leftHeight := isValid(root.Left)
	if leftHeight == math.MinInt64 {
		return math.MinInt64
	}

	rightHeight := isValid(root.Right)
	if rightHeight == math.MinInt64 {
		return math.MinInt64
	}

	heightDiff := leftHeight - rightHeight

	if abs(heightDiff) > 1 {
		return math.MinInt64
	}
	return max(leftHeight, rightHeight) + 1
}

func CheckBalancedSolution(root *Node) bool {
	return isValid(root) != math.MinInt64
}
