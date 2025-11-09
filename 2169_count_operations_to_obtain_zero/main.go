package main

func countOperations(num1 int, num2 int) int {
    var result int
    for num1 != 0 && num2 != 0 {
        if num1 >= num2 {
            num1 -= num2
        } else {
            num2 -= num1
        }
        result++
    }
    return result
}

func main() {
	countOperations(2, 3)
}
