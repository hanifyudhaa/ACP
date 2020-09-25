package main

import (
	f "fmt"
)

func main() {

	var offset int32
	var s string

	f.Print("Offset = ")
	f.Scanln(&offset)
	f.Print("string = ")
	f.Scan(&s)

	caesar(offset, s)
	
}

func caesar(offset int32, s string) {

	var temp string

	for _, i := range s {

		if offset < 0 {
			if i+offset < 97 {
				i = i + offset + 26
				temp += string(i)
			} else {
				temp += string(i + offset)
			}
		} else {
			if i+offset > 122 {
				i = i + offset - 26
				temp += string(i)
			} else {
				temp += string(i + offset)
			}
		}

	}

	f.Println(temp)
}
