package main

import f "fmt"

func draw(input int) {
	var counter int = 1

	for i := 0; i < input; i++ {
		for j := 0; j < input; j++ {

			var kelipatan3, genap int

			kelipatan3 = counter%3
			genap = counter%2

			switch {
			case kelipatan3 == 0 :
				f.Print("*\t")

			case kelipatan3 != 0 && genap == 0 :
				f.Print("$\t")

			case kelipatan3 != 0 && genap != 0 :
				f.Print("@\t")
			
			}

			counter ++

		}

		f.Println("")
	}

}

func main() {
	draw(5)
}
