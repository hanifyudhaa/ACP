package main

import f "fmt"

func main() {
	coinChange(100, []int{25, 10, 5, 1})
}

func coinChange(jumlahUang int, uangTersedia []int) {
	var a []int
	for i := 0; i < len(uangTersedia); i++ {
		for jumlahUang-uangTersedia[i] >= 0 {
			jumlahUang -= uangTersedia[i]
			a = append(a, uangTersedia[i])
		}
	}
	f.Println(a)
}
