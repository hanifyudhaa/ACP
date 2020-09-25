package main

import (
	f "fmt"
)

func main() {
	f.Println(maxSum([]int{2, 3, 4, 1, 5}, 2))
}

func maxSum(arr []int, k int) int {

	a, b := 0, 0

	for i := 0; i <= len(arr)-3; i++ {
		b = 0

		for j := i; j < k+i; j++ {
			b += arr[j]
		}

		if b > a {
			a = b
		}
	}

	return a
}
