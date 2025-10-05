package main

func maxArea(height []int) int {
    N := len(height)
    var result int
    left, right := 0, N - 1

    for right >= 0 && left < N && right > left {
        area := min(height[right], height[left]) * (right - left)
        result = max(result, area)

        if height[left] > height[right] {
            right--
        } else if height[right] > height[left] {
            left++
        } else {
            right--
            left++
        }
    }
    return result
}

func main() {
	maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
}
