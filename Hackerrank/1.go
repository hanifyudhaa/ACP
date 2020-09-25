package main

import "fmt"

func main() {
	fmt.Println(countMax([]string{"1 4", "2 3", "4 1"}))
}

func countMax(upRight []string) int64{
	// Write your code here

	var maxa, maxb int = 0, 0

	sliceA := []int{}
	sliceB := []int{}
	gabung := []int{}

	for i := 0; i < len(upRight); i++ {
		var a, b int

		fmt.Sscanf(upRight[i], "%d %d", &a, &b)
		fmt.Println(a, b)

		sliceA = append(sliceA, a)
		sliceB = append(sliceB, b)

		if a > maxa {
			maxa = a
		}

		if b > maxb {
			maxb = b
		}

	}

	fmt.Println(maxa, maxb)
	fmt.Println(sliceA, sliceB)

	var x = make([][]int, maxa)

	for i := 0; i < maxa; i++ {
		x[i] = make([]int, maxb)
	}

	for i := 0; i < len(sliceA); i++ {
		for j := 0; j < sliceA[i]; j++ {
			for k := 0; k < sliceB[i]; k++ {
				x[j][k]++
			}
		}
	}

	fmt.Println(x)

	for i := 0; i < len(x); i++ {
		gabung = append(gabung, x[i]...)
	}

	fmt.Println(gabung)

	maximum := 0

	list:= make(map[int]int)
    for _, num :=  range gabung {
		list[num] = list[num]+1
		
		if list[num] > maximum && num == len(sliceA) {
			maximum = list[num]
		}
	}
	
	fmt.Println(list)

	return int64(maximum)
}
