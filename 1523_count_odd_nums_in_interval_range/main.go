package main

func countOdds(low int, high int) int {
    var result int

    N := high - low + 1

    if N % 2 == 0 { // if number of elements is even
        result = N / 2
    } else if N % 2 == 1 { // if number of elements is odd
        if low % 2 == 1 || high % 2 == 1 {
            result = (N / 2) + 1
        } else {
            result = N / 2
        }
    }
    return result
}

func main() {
	countOdds(8, 10)
}
