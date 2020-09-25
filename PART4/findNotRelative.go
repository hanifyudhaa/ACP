package main

import f "fmt"

func main() {
	findNotRelative([]int{3, 6, 10, 12, 15}, []int{1, 3, 5, 10, 16})
}

func findNotRelative(arr1, arr2 []int) {
	for i := 0; i < len(arr1); i++ {
		temp := 0
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				temp++
			}
		}
		if temp == 0 {
			f.Println(arr1[i])
		}
	}
}
