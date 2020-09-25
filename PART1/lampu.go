package main

import f "fmt"

func main(){

  var input int

  f.Print("Masukan Bilangan Jumlah Tombol : ")
  f.Scan(&input)

  if (input%2==0){
    f.Println("Lampu Menyala")
  } else {
    f.Println("Lampu Mati")
  }

}