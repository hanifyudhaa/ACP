package main

import f "fmt"

var dynamicProgram = make([]int, 30)
var result = 0
var putar = 0

func main() {
	dynamicProgram[1] = 1
	f.Println(fib2(5))
	f.Println(dynamicProgram)
	f.Println(putar)
}

func fib2(number int) int {
	putar++
	if number < 2 {
		return number
	}

	// if dynamicProgram[number] != 0 {
	// 	return dynamicProgram[number]
	// }

	result = fib2(number-1) + fib2(number-2)
	dynamicProgram[number] = result

	f.Println(dynamicProgram)
	return result
}
