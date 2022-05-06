package main

import (
	"fmt"
)

func pow(number int) int {

	return number * number
}

func main() {
	sumOfSquares := 0
	squareOfSum := 0

	for index := 0; index <= 100; index++ {
		sumOfSquares = sumOfSquares + pow(index)
		squareOfSum = squareOfSum + index
	}

	squareOfSum = pow(squareOfSum)

	fmt.Printf("sum of the squares : %v \n", sumOfSquares)
	fmt.Printf("square of the sum : %v \n", squareOfSum)

	fmt.Printf("the difference : %v \n", squareOfSum-sumOfSquares)
}
