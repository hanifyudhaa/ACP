package main

import f "fmt"

func main() {
	maxMin([]int{1, 4, 5, 2, 4, 8, 4, 2, 5, -100, 537, -3})
}

func maxMin(input []int) {

	max, min := input[0], input[0]

	for i := 0; i < len(input); i++ {

		if input[i] > max {
			max = input[i]
		}

		if input[i] < min {
			min = input[i]
		}
	}

	f.Print("max : ", max, " ")
	f.Println("min : ", min)
}
