package main

import (
	"sort"
)

func largestPerimeter(nums []int) int {
    var result int
    N := len(nums)

    // sort in descending order
    sort.Sort(sort.Reverse(sort.IntSlice(nums)))

    for i := 0; i+2 < N; i++ {
        if (nums[i] >= (nums[i+1] + nums[i+2])) { continue }
        if (nums[i+1] >= (nums[i] + nums[i+2])) { continue }
        if (nums[i+2] >= (nums[i] + nums[i+1])) { continue }
        result = max(result, nums[i] + nums[i+1] + nums[i+2])
        break
    }
    return result
}

func main() {
	largestPerimeter([]int{1, 2, 1, 10})
}
