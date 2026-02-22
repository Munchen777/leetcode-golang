package binarywatch

import (
	"fmt"
	"math/bits"
	"strconv"
)

func readBinaryWatch(turnedOn int) []string {
	var result []string

	for hour := 0; hour < 12; hour++ {
		for minute := 0; minute < 60; minute++ {
			hBits := bits.OnesCount(uint(hour))
			mBits := bits.OnesCount(uint(minute))
			if hBits+mBits == turnedOn {
				hStr := strconv.Itoa(hour)
				mStr := strconv.Itoa(minute)
				if minute < 10 {
					mStr = "0" + mStr
				}
				result = append(result, fmt.Sprintf("%s:%s", hStr, mStr))
			}
		}
	}
	return result
}

func main() {
	readBinaryWatch(1)
}
