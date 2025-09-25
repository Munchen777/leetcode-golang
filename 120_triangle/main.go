package main

import (
	"math"
	"slices"
)

func minimumTotal(triangle [][]int) int {
    dp := make([]int, len(triangle) + 1)

    slices.Reverse(triangle)

    for _, row := range triangle {
        for i, val := range row {
            dp[i] = val + int(math.Min(float64(dp[i]), float64(dp[i + 1])))
        }
    }

    return dp[0]
}

func main() {
	minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}})
}
