package main

import f "fmt"

func main(){

  var alas, tinggi float32

  f.Print("Masukan Alas : ")
  f.Scan(&alas)
  f.Print("Masukan Tinggi : ")
  f.Scan(&tinggi)

  f.Println("Luas Segitiga =  ", 0.5*alas*tinggi)

}