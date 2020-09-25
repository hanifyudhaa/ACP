package main

import (
	f "fmt"
)

func main() {
	// long declaration
	var salary = map[string]int{"aan": 25000000}
	salary["joshep"] = 15000000
	f.Println(salary)

	//short declaration
	salary_b := map[string]int{}
	salary_b["iqbal"] = 50000000
	f.Println(salary_b)

	var salary_c = make(map[string]int)
	salary_c["hanif"] = 75000000
	salary_c["yudha"] = 15000000
	f.Println(salary_c)

	delete(salary_c, "yudha")

	f.Println(salary_c)

}
