package main

import (
	"math"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
    if numerator == 0 {
        return "0"
    }
    if denominator == 0 {
        return "0"
    }

    result := ""

    // define a sign
    if (numerator < 0) != (denominator < 0) {
        result += "-"
    }

    numerator = int(math.Abs(float64(numerator)))
    denominator = int(math.Abs(float64(denominator)))

    integerPart := numerator / denominator
    result += strconv.Itoa(integerPart)
    remainder := numerator % denominator

    if remainder == 0 {
        return result
    }

    result += "."
    remainders := make(map[int]int)

    for remainder != 0 {
        // if we've already seen this remainder
        if pos, ok := remainders[remainder]; ok {
            return result[:pos] + "(" + result[pos:] + ")"
        }
        remainders[remainder] = len(result)

        remainder *= 10
        digit := remainder / denominator
        result += strconv.Itoa(digit)
        remainder %= denominator
    }

    return result
}

func main() {
	fractionToDecimal(4, 333)
}
