package main

import (
	"slices"
)

func maxFrequencyElements(nums []int) int {
    var result int
    hashmap := make(map[int]int)

    for _, num := range nums {
        hashmap[num]++
    }

    var values []int
    for _, val := range hashmap {
        values = append(values, val)
    }

    maximum := slices.Max(values)
    for _, val := range values {
        if val == maximum {
            result += val
        }
    }

    return result
}

func main() {
	maxFrequencyElements([]int{1, 2, 2, 3, 1, 4})
}
