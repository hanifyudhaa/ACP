package main

import "fmt"

func main() {
	reductionCost([]int32{1, 2, 3, 4})
}

func reductionCost(num []int32) {
	// Write your code here

	var sum, sum2 int32 = 0, 0
	sliceA := []int32{}

	for i := 0; i < len(num); i++ {

		sum += num[i]

		sliceA = append(sliceA, sum)

	}

	fmt.Println(sliceA)

	for i := 1; i < len(sliceA); i++ {
		sum2 += sliceA[i]
	}

	fmt.Println(sum2)

}
