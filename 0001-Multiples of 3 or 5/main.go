package main

import "fmt"

func main() {
	result := 0

	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			result = result + i
		}

	}

	fmt.Println(result)

}
