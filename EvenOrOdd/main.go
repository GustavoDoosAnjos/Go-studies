package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range numbers {
		if i%2 == 0 {
			fmt.Printf("The number %v is even \n", i)
		} else {
			fmt.Printf("The number %v is odd \n", i)
		}
	}
}
