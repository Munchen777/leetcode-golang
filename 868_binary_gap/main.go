package main

import (
	"strconv"
)

func binaryGap(n int) int {
	var result int
	binaryString := strconv.FormatInt(int64(n), 2)

	for i := 0; i < len(binaryString); i++ {
		if string(binaryString[i]) == "1" {
			for j := i + 1; j < len(binaryString); j++ {
				if string(binaryString[j]) == "1" {
					result = max(result, j-i)
					break
				}
			}
		}
	}
	return result
}

func main() {
	binaryGap(22)
}
