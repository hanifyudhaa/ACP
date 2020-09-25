package main

import f "fmt"

func arrayMerge (arrayA, arrayB []string) {

  a := append(arrayA,arrayB...)

  f.Println(remomveDuplicate(a))
}

func remomveDuplicate(s []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range s {
		f.Println("Entry : ", entry)

    values := keys[entry]
    
			f.Println("Keys : ", keys)
			f.Println("Values : ", values)

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

func main () {
  arrayMerge([]string{"halo", "nama"}, []string{"halo", "hanif"})
}