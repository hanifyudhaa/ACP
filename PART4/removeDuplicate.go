package main

import (
	f "fmt"
)

func main() {
	uniqueSlice := remomveDuplicate([]int{2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 5, 6, 7})
	f.Println(uniqueSlice)
}

func remomveDuplicate(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}

	for _, entry := range intSlice {
		f.Println("Entry : ", entry)

		values := keys[entry]

		if !values {
			f.Println("Keys : ", keys)
			f.Println("Values : ", values)

			keys[entry] = true
			list = append(list, entry)
			f.Println("List : ", list)
		}

	}
	return list
}
