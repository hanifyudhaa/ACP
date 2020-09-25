package main

import f "fmt"

func main(){

  var r, tinggi float32

  f.Print("Masukan R : ")
  f.Scan(&r)
  f.Print("Masukan Tinggi : ")
  f.Scan(&tinggi)

  f.Println("Luas Segitiga =  ", 2*3.14*r*(r+tinggi))

}