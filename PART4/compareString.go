package main

import f "fmt"

func main() {
	compare("aka", "akashi")
}

func compare(string1, string2 string) {
	if len(string1) > len(string2) {
		f.Println(string1, string2)

		string1, string2 = string2, string1

		f.Println(string1, string2)

	}

	for index, i := range string1 {

		if string1[index] == string2[index] {
			f.Print(string(string1[index]))
		}

	}
}
