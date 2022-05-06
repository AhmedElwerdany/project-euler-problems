package main

import "fmt"

func main() {

big:
	for i := 2520; ; i++ {

		for j := 1; j <= 20; j++ {
			if i%j != 0 {
				continue big
			}
		}

		fmt.Println(i)
		break
	}

}
