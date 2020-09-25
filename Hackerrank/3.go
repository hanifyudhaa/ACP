package main

import (
	f "fmt"
	"sort"
)

func main() {
	uniqueSlice := remomveDuplicate([]int{1,1,2,3,4,5,6}, 2)
	f.Println(uniqueSlice)
}

func remomveDuplicate(intSlice []int, k int) int {

	sum := 0

	sort.Ints(intSlice)

	keys := make(map[int]bool)
	list := []int{}

	for _, entry := range intSlice {
		values := keys[entry]

		if !values {
			keys[entry] = true
			list = append(list, entry)
		}

	}

	f.Println(list)

	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list); j++ {
			if i != j {
				if list[i]-list[j] == -k {
					f.Println(list[i], list[j])
					sum++
				}
			}
		}
	}

	return sum
}
