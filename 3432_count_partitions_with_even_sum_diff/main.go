package main

func countPartitions(nums []int) int {
    var result, rightSum, leftSum int

    // count total sum for right part
    for _, value := range nums {
        rightSum += value
    }

    for i := 0; i < len(nums) - 1; i++ {
        leftSum += nums[i]
        rightSum -= nums[i]
        // check if difference between left and right is even
        if ((leftSum - rightSum) % 2) == 0 {
            result++
        }
    }
    return result
}

func main() {
	countPartitions([]int{10, 10, 3, 7, 6})
}
